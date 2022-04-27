package duros

import (
	"fmt"
	"net"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var hostRegex *regexp.Regexp

func init() {
	hostRegex = regexp.MustCompile(`^([a-zA-Z0-9.\[\]:%-]+)$`)
}

// EP is a duros API Endpoint
type EP struct {
	Host string // either a hostname or an IP address
	Port uint16 // port number
}

// EPs is a slice of Endpoints
type EPs []EP

// ParseEndpoint is like parseStricter(), but it disregards spaces before and after
// the endpoint string.
func ParseEndpoint(endpoint string) (EP, error) {
	return parseStricter(strings.TrimSpace(endpoint))
}

// MustParseEndpoint is similar to Parse(), but it panics if endpoint is invalid and
// can't be parsed. useful primarily for tests and global "consts" inits.
func MustParseEndpoint(endpoint string) EP {
	ep, err := ParseEndpoint(endpoint)
	if err != nil {
		panic(fmt.Sprintf("endpoint: Parse(%s): %s", endpoint, err.Error()))
	}
	return ep
}

func parseStricter(endpoint string) (EP, error) {
	host, port, err := net.SplitHostPort(endpoint)
	if err != nil {
		//nolint dunno howto convert this to errors.As
		if addrErr, ok := err.(*net.AddrError); ok {
			return EP{}, fmt.Errorf("%s", addrErr.Err)
		}
		// shouldn't happen, but...
		return EP{}, fmt.Errorf("%w", err)
	}
	if host == "" {
		return EP{}, fmt.Errorf("invalid empty host")
	}
	if !hostRegex.MatchString(host) {
		return EP{}, fmt.Errorf("invalid host %q", host)
	}
	portNum, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return EP{}, fmt.Errorf("invalid port number %q", port)
	}
	return EP{Host: host, Port: uint16(portNum)}, nil
}

// isValid returns true for a valid EP
func (ep EP) isValid() bool {
	return ep != EP{}
}

func (ep EP) String() string {
	return net.JoinHostPort(ep.Host, strconv.FormatUint(uint64(ep.Port), 10))
}

func (eps EPs) String() string {
	tgts := make([]string, len(eps))
	for i, ep := range eps {
		tgts[i] = ep.String()
	}
	return strings.Join(tgts, ",")
}

// isValid returns true for a valid EPs
func (eps EPs) isValid() bool {
	for _, ep := range eps {
		if !ep.isValid() {
			return false
		}
	}
	return len(eps) >= 1
}

// clone create a copy of the given EPs
func (eps EPs) clone() EPs {
	res := make([]EP, len(eps))
	copy(res, eps)
	return res
}

// for sort.Interface:
// Len returns the length of the EPs
func (eps EPs) Len() int { return len(eps) }

// Swap the content of the EPs
func (eps EPs) Swap(i, j int) { eps[i], eps[j] = eps[j], eps[i] }

// Less compares two EPs entries for equality
func (eps EPs) Less(i, j int) bool { return strings.Compare(eps[i].String(), eps[j].String()) == -1 }

func canonicalize(targets []string, parser func(string) (EP, error)) (EPs, error) {
	if len(targets) == 0 {
		return nil, nil
	}

	uniq := make(map[EP]bool)
	for _, tgt := range targets {
		ep, err := parser(tgt)
		if err != nil {
			return nil, err
		}
		uniq[ep] = true
	}

	res := make([]EP, 0, len(uniq))
	for k := range uniq {
		res = append(res, k)
	}
	sort.Sort(EPs(res))
	return res, nil
}

// ParseEndpoints parse a slice of strings in the form of <host>:<port> into a slice of EP.
func ParseEndpoints(endpoints []string) (EPs, error) {
	return canonicalize(endpoints, ParseEndpoint)
}

// MustParseEndpoints is equal to ParseEndpoints but panics on error
func MustParseEndpoints(endpoints []string) EPs {
	eps, err := ParseEndpoints(endpoints)
	if err != nil {
		panic(fmt.Sprintf("endpoint: ParseEndpoints(%s): %s", endpoints, err.Error()))
	}
	return eps
}

// ParseCSV parses a string containing comma-separated list of target
// endpoints, validates them syntactically, and returns a slice of EP structs.
// it does NOT attempt to resolve the names present in the endpoints nor does
// it try to connect to any of the targets. `targets` must be in a format:
//     <host>:<port>[,<host>:<port>...]
func ParseCSV(endpoints string) (EPs, error) {
	return canonicalize(strings.Split(endpoints, ","), parseStricter)
}

// MustParseCSV is equal to ParseCSV but panics on error
func MustParseCSV(endpoints string) EPs {
	eps, err := ParseCSV(endpoints)
	if err != nil {
		panic(fmt.Sprintf("endpoint: ParseCSV(%s): %s", endpoints, err.Error()))
	}
	return eps
}
