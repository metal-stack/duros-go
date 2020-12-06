package duros

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	v2 "github.com/metal-stack/duros-go/api/duros/v2"
)

type lbJWTClaims struct {
	// "standard" (aka "registered") claims:
	//   mandatory: sub, exp
	//   optional (validated and logged if present): nbf, iss, jti
	jwt.StandardClaims

	Roles []string `json:"roles,omitempty"` // mandatory custom claim
}

// NewJWTTokenForCredential create a new JWTToken where subject and kid is taken from the credential
func NewJWTTokenForCredential(credential *v2.Credential, roles []string, expires time.Duration, keyPair *rsa.PrivateKey) (string, error) {
	return NewJWTToken(credential.ProjectName, credential.ID, roles, expires, keyPair)
}

// NewJWTToken create a JWT Token to use to authenticate against a duros API endpoint
//
// subject: 'sub' claim, who will be using this JWT, example a persons or tenants name
// kid: this is the "key ID", the name of the credential (pub key) as uploaded
// to LightOS. it is of the form "<scope>:<name>", e.g.:
//     system:root - your root pub key, installed during system deployment
//     tenant-foo:first-cred - pub key of tenant tenant-foo uploaded as credential named first-cred.
//   the JWTs are validated using the specific pub keys, so a corresponding
//   credential must already exist in LightOS.
// roles: list of roles this token should contain, must be in the form of
//   foo:admin which gives this user  (subject) admin rights to the foo resource
// expires: Duration after which this token will expire.
// keyPair: RSA public and private key which should be used to sign this token
func NewJWTToken(subject string, kid string, roles []string, expires time.Duration, keyPair *rsa.PrivateKey) (string, error) {
	now := time.Now().UTC()
	claims := &lbJWTClaims{
		// see overview of "standard" JWT claims as used by jwt-go here:
		//   https://godoc.org/github.com/dgrijalva/jwt-go#StandardClaims
		// see the semantics of the registered claims here:
		//   https://en.wikipedia.org/wiki/JSON_Web_Token#Standard_fields
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(expires).UTC().Unix(),
			IssuedAt:  now.Unix(),

			// ID is for your traceability, doesn't have to be UUID:
			Id: uuid.New().String(),

			// put name/title/ID of whoever will be using this JWT here:
			Subject: subject,
		},
		Roles: roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = kid
	res, err := token.SignedString(keyPair)
	if err != nil {
		return "", fmt.Errorf("unable to sign RS256 JWT: %v", err)
	}
	return res, nil
}
