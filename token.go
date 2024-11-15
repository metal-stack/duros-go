package duros

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	v2 "github.com/metal-stack/duros-go/api/duros/v2"
)

type lbJWTClaims struct {
	//   mandatory: sub, exp
	//   optional (validated and logged if present): nbf, iss, jti
	jwt.RegisteredClaims

	Roles []string `json:"roles,omitempty"` // mandatory custom claim
}

// NewJWTTokenForCredential create a new JWTToken where subject and kid is taken from the credential
func NewJWTTokenForCredential(subject, issuer string, credential *v2.Credential, roles []string, expires time.Duration, keyPair *rsa.PrivateKey) (string, error) {
	return NewJWTToken(subject, issuer, credential.ProjectName+":"+credential.ID, roles, expires, keyPair)
}

// NewJWTToken create a JWT Token to use to authenticate against a duros API endpoint
//
// subject: 'sub' claim, who will be using this JWT, example a persons or tenants name
// kid: this is the "key ID", the name of the credential (pub key) as uploaded
// to LightOS. it is of the form "<scope>:<name>", e.g.:
//
//	  system:root - your root pub key, installed during system deployment
//	  tenant-foo:first-cred - pub key of tenant tenant-foo uploaded as credential named first-cred.
//	the JWTs are validated using the specific pub keys, so a corresponding
//	credential must already exist in LightOS.
//
// roles: list of roles this token should contain, must be in the form of
//
//	foo:admin which gives this user  (subject) admin rights to the foo resource
//
// expires: Duration after which this token will expire.
// keyPair: RSA public and private key which should be used to sign this token
func NewJWTToken(subject, issuer string, kid string, roles []string, expires time.Duration, keyPair *rsa.PrivateKey) (string, error) {
	now := time.Now().UTC()
	claims := &lbJWTClaims{
		// see overview of "registered" JWT claims as used by jwt-go here:
		//   https://pkg.go.dev/github.com/golang-jwt/jwt/v5?utm_source=godoc#RegisteredClaims
		// see the semantics of the registered claims here:
		//   https://en.wikipedia.org/wiki/JSON_Web_Token#Standard_fields
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expires)),
			IssuedAt:  jwt.NewNumericDate(now),

			// ID is for your traceability, doesn't have to be UUID:
			ID: uuid.New().String(),

			// put name/title/ID of whoever will be using this JWT here:
			Subject: subject,
			Issuer:  issuer,
		},
		Roles: roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = kid
	res, err := token.SignedString(keyPair)
	if err != nil {
		return "", fmt.Errorf("unable to sign RS256 JWT: %w", err)
	}
	return res, nil
}
