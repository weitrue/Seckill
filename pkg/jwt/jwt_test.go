package jwt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWT_CreateToken(t *testing.T) {
	c := &Config{Issuer: "backend-micro", SecretKey: "", ExpirationTime: 72 * time.Hour}
	_, err := NewJWT(c)
	assert.EqualError(t, err, "jwt: illegal jwt configure")

	c = &Config{Issuer: "backend-micro", SecretKey: "ABCDEFGH", ExpirationTime: -72 * time.Hour}
	_, err = NewJWT(c)
	assert.EqualError(t, err, "jwt: illegal jwt configure")

	c = &Config{Issuer: "backend-micro", SecretKey: "ABCDEFGH", ExpirationTime: 72 * time.Hour}
	j, err := NewJWT(c)
	require.NoError(t, err)

	type testToken struct {
		UserId int64 `json:"user_id"`
		RoleId int64 `json:"role_id"`
	}
	token := &testToken{UserId: 10000000, RoleId: 1}
	tokenStr, err := j.CreateToken(token)
	if assert.NoError(t, err) {
		t.Log(tokenStr)
	}

	parseToken := &testToken{}
	err = j.ParseToken(tokenStr, parseToken)
	assert.NoError(t, err)
	if assert.Equal(t, token, parseToken) {
		t.Logf("%+v", parseToken)
	}
}

func TestJWT_ParseToken(t *testing.T) {
	c := &Config{Issuer: "backend-micro", SecretKey: "", ExpirationTime: 72 * time.Hour}
	_, err := NewJWT(c)
	assert.EqualError(t, err, "jwt: illegal jwt configure")

	c = &Config{Issuer: "backend-micro", SecretKey: "ABCDEFGH", ExpirationTime: -72 * time.Hour}
	_, err = NewJWT(c)
	assert.EqualError(t, err, "jwt: illegal jwt configure")

	c = &Config{Issuer: "backend-micro", SecretKey: "ABCDEFGH", ExpirationTime: 72 * time.Hour}
	j, err := NewJWT(c)
	require.NoError(t, err)

	type testToken struct {
		UserId int64 `json:"user_id"`
		RoleId int64 `json:"role_id"`
	}
	token := &testToken{UserId: 10000000, RoleId: 1}
	tokenStr, err := j.CreateToken(token)
	if assert.NoError(t, err) {
		t.Log(tokenStr)
	}

	parseToken := &testToken{}
	err = j.ParseToken(tokenStr, parseToken)
	assert.NoError(t, err)
	if assert.Equal(t, token, parseToken) {
		t.Logf("%+v", parseToken)
	}
}

func TestName(t *testing.T) {
	dd := Md5Buf([]byte(TrimBearer("IY55z9Boy2kBjatl41NH7Bent1yXRESgKjWbqeHehMo=")))
	fmt.Println(dd)
}

func Md5Buf(buf []byte) string {
	hashMd5 := md5.New()
	hashMd5.Write(buf)
	//return fmt.Sprintf("%x", hashMd5.Sum(nil))
	return hex.EncodeToString(hashMd5.Sum(nil))
}

func TrimBearer(token string) string {
	token = strings.TrimSpace(token)
	if len(token) >= 7 && strings.EqualFold(token[:7], "Bearer ") {
		token = strings.TrimSpace(token[7:])
	}
	return token
}
