package vrcode

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/weitrue/Seckill/pkg/errcode"
	"github.com/weitrue/Seckill/pkg/xhttp"

	"github.com/weitrue/Seckill/infrastructure/services/remote/qrcode/cache"
)

const (
	// Cloud 运营商
	Cloud = "VRCode"
	// CodeSucc 响应成功，仅为 0 时才响应成功
	CodeSucc = 0
	// CodeFailedAuth 请求失败，用户名或密码错误
	CodeFailedAuth = 1
	// CodeFailedToken 请求失败，token 无效
	CodeFailedToken = 2
	// CodeFailedLimit 制码失败，制码次数超出限制
	CodeFailedLimit = 3
	// CodeFailedErr 制码失败，请重试或联系管理人员
	CodeFailedErr = 5
	// CodeFailedParam 参数错误，请核对后重新请求
	CodeFailedParam = 6
)

var (
	tokenCache = cache.NewObjCache()
	errMap     = map[int64]string{
		CodeFailedErr: "制码失败，请重试或联系管理人员",
	}
)

// VRCode 快递100快递查询服务结构详情
type VRCode struct {
	domain   string // 域名
	userName string // 用户名
	password string // 密码
	mode     int64  // 模式
	client   *xhttp.Client
}

// NewVRCode 新建VRCode客户端
func NewVRCode(domain, userName, password string, mode int64) (*VRCode, error) {
	if domain == "" || userName == "" || password == "" || mode <= 0 {
		return nil, errors.New("qrcode: illegal vrcode config")
	}

	// 绕过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}

	return &VRCode{
		domain:   domain,
		userName: userName,
		password: password,
		mode:     mode,
		client:   xhttp.NewClientWithHTTPClient(client),
	}, nil
}

// GetQRCode 获取三维码
// imgCode 制码底图的base64编码，不包含空格和头部
// codeInfo 编码内容
// returns string 三维码图片的 base64 位值 (图片格式为 png,1000*1000 像素)
func (c *VRCode) GetQRCode(imgCode, codeInfo string) (string, error) {
	// 获取token
	token, err := c.getToken()
	if err != nil {
		return "", err
	}

	// 请求地址
	url := c.getUrl("/vrcode/create")
	// 构建请求头
	header := make(map[string]string)
	header["Content-Type"] = "application/json;utf-8"
	header["Authorization"] = token

	// 构建签名参数
	paramBytes, err := json.Marshal(&createVRCodeReq{
		Mode:           c.mode,
		CodeInfoPublic: codeInfo,
		Img:            imgCode,
	})
	if err != nil {
		return "", errors.WithMessage(err, "qrcode: createVRCodeReq Marshal to json err")
	}

	// 远程调用创建三维码
	resp, err := c.createVRCode(url, string(paramBytes), header)
	if err != nil {
		return "", err
	}

	// token失效 重新登录再获取三维码
	if resp.Code == CodeFailedToken || resp.Status == http.StatusUnauthorized {
		lResp, err := c.login()
		if err != nil {
			return "", err
		}

		header["Authorization"] = lResp.Token
		resp, err = c.createVRCode(url, string(paramBytes), header)
		if err != nil {
			return "", err
		}
	}

	// Status>0时, code=0状态不能代表正常返回结果, 应以status状态码为准
	code := resp.Code
	message := resp.Message
	if resp.Status > 0 && resp.Code == 0 {
		code = resp.Status
		if resp.Error != "" {
			message = resp.Error
		}
	}

	return getResult(resp.VRCode, message, code)
}

// getToken 获取token
func (c *VRCode) getToken() (string, error) {
	// 先从缓存中获取token
	if token, ok := tokenCache.Get(c.userName); ok {
		return token.(string), nil
	}

	// 缓存不存在，远程登录获取token
	resp, err := c.login()
	if err != nil {
		return "", err
	}

	return getResult(resp.Token, resp.Message, resp.Code)
}

// login 登录
// 登录成功之后返回的token值，并存入缓存
func (c *VRCode) login() (*loginResp, error) {
	url := c.getUrl("/auth/login")
	// 构建请求头
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	// 构建签名参数
	paramBytes, err := json.Marshal(&loginReq{
		UserName: c.userName,
		Password: c.password,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "qrcode: loginReq Marshal to json err")
	}

	var resp loginResp
	err = c.client.Call(http.MethodPost, url, header, strings.NewReader(string(paramBytes)), &resp)
	if err != nil {
		return nil, errors.WithMessage(err, "qrcode: login response err")
	}

	tokenCache.Set(c.userName, resp.Token)

	return &resp, nil
}

// loginReq 登录请求
type loginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// loginResp 登录响应
type loginResp struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

// 创建VRCode
func (c *VRCode) createVRCode(url, body string, header map[string]string) (*createVRCodeResp, error) {
	var resp createVRCodeResp
	err := c.client.Call(http.MethodPost, url, header, strings.NewReader(body), &resp)
	if err != nil {
		return nil, errors.WithMessage(err, "qrcode: createVRCode response err")
	}

	return &resp, nil
}

// createVRCodeReq 创建vrcode请求
type createVRCodeReq struct {
	Mode           int64  `json:"mode"`
	CodeInfoPublic string `json:"codeInfoPublic"`
	Img            string `json:"img"`
}

// createVRCodeResp 创建vrcode响应
type createVRCodeResp struct {
	Code           int64  `json:"code"`    // 返回状态码
	Message        string `json:"message"` // 返回消息
	SelfCheck      int64  `json:"selfCheck"`
	InaptRect      int64  `json:"inaptRect"`
	VRCode         string `json:"vrcode"` // 生成图片base64编码
	PointNum       int64  `json:"pointNum"`
	Cost           string `json:"cost"` // 耗费时间
	IfFaceDetected int64  `json:"ifFaceDetected"`
	OBSURL         string `json:"OBSURL"`
	VersionDetail  string `json:"versionDetail"`
	Status         int64  `json:"status"` // 服务内部错误，会有对应status，正常返回不会出现该字段
	Error          string `json:"error"`  // 服务内部错误，会有对应error，正常返回不会出现该字段
}

func (c *VRCode) getUrl(api string) string {
	if c.domain[len(c.domain)-1:] == "/" && strings.Index(api, "/") == 0 {
		return c.domain + api[1:]
	} else if c.domain[len(c.domain)-1:] == "/" && strings.Index(api, "/") > 0 {
		return c.domain + api
	} else if c.domain[len(c.domain)-1:] != "/" && strings.Index(api, "/") == 0 {
		return c.domain + api
	} else {
		return c.domain + "/" + api
	}
}

// getResult 获取结果
func getResult(res, message string, code int64) (string, error) {
	switch code {
	case CodeSucc, http.StatusOK:
		if res != "" {
			return res, nil
		}

		return "", errors.New("qrcode: get result empty string err")
	default:
		if msg, ok := errMap[code]; ok {
			return "", errcode.NewCustomErr(msg)
		}

		return "", errors.WithMessage(fmt.Errorf("qrcode: remote call api err, code: %d ", code), message)
	}
}
