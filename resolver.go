package duros

import (
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

// LightOS cluster resolver: -------------------------------------------------

// lbResolver is similar to gRPC manual.Resolver in that it's primed by a number
// of LightOS cluster member addresses. the difference is that lbResolver tries
// to rotate this list of addresses on failures. on the one hand, it keeps the
// client talking mostly to the same mgmt API server avoiding consistency
// issues, but on the other it avoids some of the pathological cases of
// "pick_first" balancer (e.g. the fact that it never even TRIES anything other
// than the "first" if that first was inaccessible on first dial - it just
// burns the entire deadline budget on fruitless attempts to get through to the
// first address and then fails DialContext() on 'i/o timeout'. i guess it's
// just an unfortunate interplay between "pick_first" and default ClientConn
// behaviour - grep for "We can potentially spend all the time trying the
// first address" in addrConn.resetTransport()).
//
// later on this can also be extended quite trivially to update the resolver
// with the latest LightOS cluster member list at runtime based on info fetched
// straight from the horses mouth, to accommodate for added/removed nodes.
//
// oh, yeah, and lbResolver "is also a resolver builder". it's traditional, you
// know...
type lbResolver struct {
	// scheme is not really a URL scheme, but rather a unique per-resolver
	// (or, equivalently, per-LightOS cluster) thing, an artefact of how
	// the dialling gRPC machinery be looking up the resolver later.
	scheme string
	log    *zap.SugaredLogger

	cc resolver.ClientConn

	// mu protects all things EPs related. a bit heavy handed, but trivial
	// and no contention is expected.
	mu   sync.Mutex
	eps  EPs    // LightOS node EPs in the order to be tried.
	tgts string // cached string repr of `eps`

}

func newLbResolver(log *zap.SugaredLogger, scheme string, targets EPs) *lbResolver {
	r := &lbResolver{
		scheme: scheme,
		eps:    targets.clone(),
		log:    log,
		tgts:   targets.String(),
	}
	r.log.Infow("initializing...", "targets", r.tgts)
	return r
}

// Build() implements (together with Scheme()) the resolver.Builder interface
// and is the way gRPC requests to create a resolver instance when its
// pseudo-scheme is mentioned while dialling. in case of lbResolver, each
// resolver is unique to a LightOS cluster. lbResolver "builds" itself.
func (r *lbResolver) Build(
	target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions,
) (resolver.Resolver, error) {
	r.cc = cc
	r.log.Infow("building...", "targets", r.tgts)
	r.updateCCState()
	return r, nil
}

func (r *lbResolver) Scheme() string {
	return r.scheme
}

// ResolveNow() rotates the current list of LightOS cluster node addresses left
// by one and returns it. since this method is typically called when gRPC has
// connectivity problems to the first node in the slice, this achieves the
// effect of making dialer to try to connect to the next node, while putting
// the offending node at the end of the list of nodes to retry.
//
// in the future it could, potentially, be made to trigger querying of the
// LightOS cluster it represents for the latest list of cluster members...
func (r *lbResolver) ResolveNow(o resolver.ResolveNowOptions) {
	r.mu.Lock()
	// not particularly efficient mem-wise, but this should be rare enough
	// event for it not to matter...
	if len(r.eps) > 1 {
		r.eps = append(r.eps[1:], r.eps[0])
	}
	r.tgts = r.eps.String()
	r.mu.Unlock()
	r.log.Infow("resolving...", "targets", r.tgts)

	r.updateCCState()
}

func (r *lbResolver) Close() {
	r.log.Info("closing...")
}

// updateCCState() updates the underlying ClientConn with the currently
// known list of LightOS cluster nodes.
func (r *lbResolver) updateCCState() {
	r.mu.Lock()
	addrs := make([]resolver.Address, len(r.eps))
	for i, ep := range r.eps {
		addrs[i].Addr = ep.String()
	}
	r.mu.Unlock()
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
