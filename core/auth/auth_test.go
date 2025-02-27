package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yomorun/yomo/core/frame"
)

// mockAuth implement `Authentication` interface,
// Authenticate returns true if authed is true, false to false.
type mockAuth struct{ authed bool }

func (auth mockAuth) Init(args ...string)              {}
func (auth mockAuth) Authenticate(payload string) bool { return auth.authed }
func (auth mockAuth) Name() string                     { return "mock" }

func init() { Register(mockAuth{}) }

func TestAuthenticate(t *testing.T) {
	type args struct {
		auths map[string]Authentication
		obj   Object
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "auths is nil",
			args: args{
				auths: nil,
				obj:   frame.NewHandshakeFrame("", "", byte(1), []frame.Tag{}, "mock", "mock_payload"),
			},
			want: true,
		},
		{
			name: "auth obj is nil",
			args: args{
				auths: map[string]Authentication{"mock": mockAuth{authed: true}},
				obj:   nil,
			},
			want: false,
		},
		{
			name: "auth obj not found",
			args: args{
				auths: map[string]Authentication{"mock": mockAuth{authed: true}},
				obj:   frame.NewHandshakeFrame("", "", byte(1), []frame.Tag{}, "mock_not_match", "mock_payload"),
			},
			want: false,
		},
		{
			name: "auth success",
			args: args{
				auths: map[string]Authentication{"mock": mockAuth{authed: true}},
				obj:   frame.NewHandshakeFrame("", "", byte(1), []frame.Tag{}, "mock", "mock_payload"),
			},
			want: true,
		},
		{
			name: "auth failed",
			args: args{
				auths: map[string]Authentication{"mock": mockAuth{authed: false}},
				obj:   frame.NewHandshakeFrame("", "", byte(1), []frame.Tag{}, "mock", "mock_payload"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Authenticate(tt.args.auths, tt.args.obj)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewCredential(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name string
		args args
		want *Credential
	}{
		{
			name: "key value pair",
			args: args{
				payload: "token:the-token",
			},
			want: &Credential{
				name:    "token",
				payload: "the-token",
			},
		},
		{
			name: "not key value pair",
			args: args{
				payload: "abcdefg",
			},
			want: &Credential{
				name:    "none",
				payload: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCredential(tt.args.payload)

			assert.Equal(t, tt.want.Name(), got.Name())
			assert.Equal(t, tt.want.Payload(), got.Payload())
		})
	}
}
