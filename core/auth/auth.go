package auth

import (
	"strings"
)

var (
	auths = make(map[string]Authentication)
)

// Authentication for server
type Authentication interface {
	// Init authentication initialize arguments
	Init(args ...string)
	// Authenticate authentication client's credential
	Authenticate(payload string) bool
	// Name authentication name
	Name() string
}

// Register register authentication
func Register(authentication Authentication) {
	auths[authentication.Name()] = authentication
}

// GetAuth get authentication by name
func GetAuth(name string) (Authentication, bool) {
	auth, ok := auths[name]
	return auth, ok
}

// Credential client credential
type Credential struct {
	name    string
	payload string
}

// NewCredential create client credential
func NewCredential(payload string) *Credential {
	idx := strings.Index(payload, ":")
	if idx != -1 {
		authName := payload[:idx]
		idx++
		authPayload := payload[idx:]
		return &Credential{
			name:    authName,
			payload: authPayload,
		}
	}
	return &Credential{name: "none"}
}

// Payload client credential payload
func (c *Credential) Payload() string {
	return c.payload
}

// Name client credential name
func (c *Credential) Name() string {
	return c.name
}

// Object is the object to be authenticated,
// The Object usually be pass to `Authenticate` function to be authed.
type Object interface {
	// AuthName returns the auth name, the name will be used to find the auth way.
	AuthName() string

	// AuthPayload returns the auth payload be passed to `auth.Authenticate`.
	AuthPayload() string
}

// Authenticate finds an authentication way in `auths` and authenticates the Object.
//
// If `auths` is nil or empty, It returns true, It think that authentication is not required.
func Authenticate(auths map[string]Authentication, obj Object) bool {
	if auths == nil || len(auths) <= 0 {
		return true
	}

	if obj == nil {
		return false
	}

	auth, ok := auths[obj.AuthName()]
	if !ok {
		return false
	}

	return auth.Authenticate(obj.AuthPayload())
}
