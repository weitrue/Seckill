package jwt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestFromMD(t *testing.T) {
	token := &Token{
		TokenType:    "access",
		RandomId:     "abcdefgh",
		LoginType:    LoginTypePhone,
		IsAdmin:      true,
		UserId:       1000,
		DepartmentId: 2000,
		CompanyId:    3000,
		RoleIds:      []int64{4000, 5000, 6000},
	}

	var pairs []string
	token.Visit(func(key, val string) bool {
		pairs = append(pairs, key, val)
		return true
	})

	md := metadata.Pairs(pairs...)
	t.Log(md)

	parseToken, ok := FromMD(md)
	assert.True(t, ok)
	if assert.Equal(t, token, parseToken) {
		t.Logf("%+v", parseToken)
	}
}

func TestWrapAndUnwrapContext(t *testing.T) {
	token := &Token{
		TokenType:    "access",
		RandomId:     "abcdefgh",
		LoginType:    LoginTypePhone,
		IsAdmin:      true,
		UserId:       1000,
		DepartmentId: 2000,
		CompanyId:    3000,
		RoleIds:      []int64{4000, 5000, 6000},
	}

	var pairs []string
	token.Visit(func(key, val string) bool {
		pairs = append(pairs, key, val)
		return true
	})

	ctx := metadata.AppendToOutgoingContext(context.Background(), pairs...)
	md, ok := metadata.FromOutgoingContext(ctx)
	assert.True(t, ok)
	t.Log(md)

	mdToken, ok := FromMD(md)
	assert.True(t, ok)
	t.Log(mdToken)

	ctx = context.WithValue(ctx, TokenKey, mdToken)
	ctxToken, ok := FromContext(ctx)
	assert.True(t, ok)
	t.Log(ctxToken)
}
