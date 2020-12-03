package duros

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	keys, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Errorf("failed to generate RSA key pair: %s", err)
	}
	token, err := NewJWTToken(
		"tenant-foo",          // 'sub' claim, who'll be using this JWT
		"foo:cred-1",          // matches LightOS cred 'cred-1' in proj 'foo'
		[]string{"foo:admin"}, // user of JWT can act as 'admin' of 'foo'
		30*24*time.Hour,       // Expire after 30Days
		keys,                  // RSA private key that corresponds to 'cred-1' pub key
	)
	if err != nil {
		fmt.Printf("unable to create a JWT token:%v\n", err)
		return
	}
	t.Logf("%s\n", token)

	// the above should result in a JWT being printed on stdout. if you
	// decode that JWT, e.g. using one of the online parsers like the one
	// on jwt.io homepage, the result should look similar to:
	//   {
	//    alg: "RS256",
	//    kid: "foo:cred-1",
	//    typ: "JWT"
	//   }.
	//   {
	//    exp: 1605690043,
	//    jti: "5c855db0-79cf-4d1d-b8a2-64a1222379b1",
	//    iat: 1605603643,
	//    sub: "tenant-foo",
	//    roles: [
	//     "foo:admin"
	//    ]
	//   }.
	//   [signature]
}
