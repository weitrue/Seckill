package errcode

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weitrue/Seckill/pkg/utils/convert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCodeToErr(t *testing.T) {
	for c, e := range codeToErr {
		assert.Equal(t, c, e.Code())
	}
}

func TestIsErr(t *testing.T) {
	cases := []struct {
		e      error
		expect bool
	}{
		{e: ErrCustom, expect: true},
		{e: ErrUnexpected, expect: true},
		{e: errors.New("测试错误"), expect: false},
		{e: ErrTokenExpire, expect: true},
		{e: NewErr(7000, "test"), expect: true},
		{e: nil, expect: true},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, IsErr(c.e))
	}
}

func TestParseErr(t *testing.T) {
	cases := []struct {
		e              error
		expectCode     uint32
		expectHTTPCode int
		expectErr      string
	}{
		{e: ErrCustom, expectCode: 7000, expectHTTPCode: 200, expectErr: "自定义错误"},
		{e: ErrUnexpected, expectCode: 7777, expectHTTPCode: 200, expectErr: "服务器繁忙，请稍后重试"},
		{e: errors.New("测试错误"), expectCode: 7777, expectHTTPCode: 200, expectErr: "服务器繁忙，请稍后重试"},
		{e: ErrTokenExpire, expectCode: 10004, expectHTTPCode: 401, expectErr: "token过期"},
		{e: nil, expectCode: 200, expectHTTPCode: 200, expectErr: "OK"},
		{e: WrapErr(ErrCustom), expectCode: 7000, expectHTTPCode: 200, expectErr: "自定义错误"},
		{e: WrapErr(NewErr(7000, "test 7000 code")), expectCode: 7000, expectHTTPCode: 200, expectErr: "test 7000 code"},
		{e: status.Error(codes.Code(7000), "测试错误"), expectCode: 7000, expectHTTPCode: 200, expectErr: "测试错误"},
	}

	for _, c := range cases {
		e := ParseErr(c.e)
		assert.Equal(t, c.expectCode, e.Code())
		assert.Equal(t, c.expectHTTPCode, e.HTTPCode())
		assert.Equal(t, c.expectErr, e.Error())
	}
}

func TestParseCode(t *testing.T) {
	cases := []struct {
		code       uint32
		expectCode uint32
		expectErr  string
	}{
		{code: 7000, expectCode: 7000, expectErr: "自定义错误"},
		{code: 7777, expectCode: 7777, expectErr: "服务器繁忙，请稍后重试"},
		{code: 200, expectCode: 200, expectErr: "OK"},
		{code: 123, expectCode: 7777, expectErr: "服务器繁忙，请稍后重试"},
	}

	for _, c := range cases {
		e := ParseCode(c.code)
		assert.NotNil(t, e)
		assert.Equal(t, c.expectCode, e.Code())
		assert.Equal(t, c.expectErr, e.Error())
	}
}

func TestCheckErrCode(t *testing.T) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, "./errcode.go", nil, parser.AllErrors)
	assert.NoError(t, err)

	var count int
	for _, decl := range f.Decls {
		if d, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range d.Specs {
				if s, ok := spec.(*ast.ValueSpec); ok {
					for i := 0; i < len(s.Names); i++ {
						v, ok := s.Values[i].(*ast.CallExpr)
						if ok {
							for _, arg := range v.Args {
								var code uint32
								switch a := arg.(type) {
								case *ast.BasicLit:
									if a.Kind == token.INT {
										count++
										code = convert.ToUint32(a.Value)
									}
								case *ast.Ident:
									aa, _ := a.Obj.Decl.(*ast.ValueSpec).Values[0].(*ast.BasicLit)
									if aa.Kind == token.INT {
										count++
										code = convert.ToUint32(aa.Value)
									}
								}
								if code > 0 {
									if _, ok := codeToErr[code]; !ok {
										t.Logf("this code: (%d: %s) is not in the map", code, s.Names[i].Name)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	assert.Equal(t, len(codeToErr), count)
}

func TestParseVars(t *testing.T) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, "./errcode.go", nil, parser.AllErrors)
	assert.NoError(t, err)

	b := &strings.Builder{}
	b.WriteString("# 存证、溯源后台管理模块接口错误码\n\n" +
		"> [!ATTENTION|label:注意]\n" +
		"> 错误码与 HTTP 状态码都为 `200` 时，才代表接口请求成功\n\n" +
		"| **错误** | **错误码** | **释义** | **HTTP 状态码** |\n" +
		"|:---------|:-----------|:---------|:----------------|\n")

	for _, decl := range f.Decls {
		if d, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range d.Specs {
				if s, ok := spec.(*ast.ValueSpec); ok {
					for i := 0; i < len(s.Names); i++ {
						id := s.Names[i]
						v, ok := s.Values[i].(*ast.CallExpr)
						if ok {
							b.WriteString(fmt.Sprintf("| %s | ", id.Name))
							for _, arg := range v.Args {
								switch a := arg.(type) {
								case *ast.BasicLit:
									b.WriteString(fmt.Sprintf("%s | ", strings.ReplaceAll(a.Value, `"`, "")))
								case *ast.Ident:
									vv := a.Obj.Decl.(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value
									b.WriteString(fmt.Sprintf("%s | ", strings.ReplaceAll(vv, `"`, "")))
								case *ast.SelectorExpr:
									b.WriteString(fmt.Sprintf("<font color='red'>%s.%s</font> | ", a.X, a.Sel.Name))
								}
							}
							if len(v.Args) < 3 {
								b.WriteString("<font color='green'>200</font> | ")
							}
							b.WriteString("\n")
						}
					}
				}
			}
		}
	}

	out := b.String()
	out = strings.ReplaceAll(out, "http.StatusBadRequest", "400")
	out = strings.ReplaceAll(out, "http.StatusUnauthorized", "401")

	fmt.Println(out)
}
