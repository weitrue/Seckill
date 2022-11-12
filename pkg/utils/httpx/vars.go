package httpx

import (
	"net/http"
)

const (
	// ApplicationJson means application/json.
	ApplicationJson = "application/json"
	// ContentEncoding means Content-Encoding.
	ContentEncoding = "Content-Encoding"
	// ContentSecurity means X-Content-Security.
	ContentSecurity = "X-Content-Security"
	// ContentType means Content-Type.
	ContentType = "Content-Type"
	// KeyField means key.
	KeyField = "key"
	// SecretField means secret.
	SecretField = "secret"
	// TypeField means type.
	TypeField = "type"
	// CryptionType means cryption.
	CryptionType = 1
)

const (
	// CodeSignaturePass means signature verification passed.
	CodeSignaturePass = iota
	// CodeSignatureInvalidHeader means invalid header in signature.
	CodeSignatureInvalidHeader
	// CodeSignatureWrongTime means wrong timestamp in signature.
	CodeSignatureWrongTime
	// CodeSignatureInvalidToken means invalid token in signature.
	CodeSignatureInvalidToken
)

var pathVars = contextKey("pathVars")

// Vars parses path variables and returns a map.
func Vars(r *http.Request) map[string]string {
	vars, ok := r.Context().Value(pathVars).(map[string]string)
	if ok {
		return vars
	}

	return nil
}

type contextKey string

func (c contextKey) String() string {
	return "rest/internal/context key: " + string(c)
}
