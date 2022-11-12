package qrcode

import (
	"github.com/pkg/errors"

	"github.com/weitrue/Seckill/infrastructure/services/remote/qrcode/vrcode"
)

// QRCode 二维码接口
type QRCode interface {
	// GetQRCode 获取二维码
	GetQRCode(imgCode, codeInfo string) (string, error)
}

// Config 二维码配置
type Config struct {
	Cloud    string // 服务商（当前支持VRCode）
	Domain   string // 域名
	UserName string // 用户
	Password string // 密码
	Mode     int64  // 模式
}

// NewQRCode 新建二维码生成对象
func NewQRCode(c *Config) (QRCode, error) {
	if c == nil {
		return nil, errors.New("express: illegal qrcode configure")
	}

	var client QRCode
	var err error

	switch c.Cloud {
	case vrcode.Cloud:
		client, err = vrcode.NewVRCode(c.Domain, c.UserName, c.Password, c.Mode)
	default:
		err = errors.New("illegal qrcode cloud")
	}
	if err != nil {
		return nil, errors.WithMessage(err, "qrcode: new qrcode service err")
	}

	return &defaultQRCode{
		client: client,
	}, nil
}

// MustNewQRCode 新建二维码生成对象
func MustNewQRCode(c *Config) QRCode {
	e, err := NewQRCode(c)
	if err != nil {
		panic(err)
	}

	return e
}

// defaultExpress 默认快递查询服务结构详情
type defaultQRCode struct {
	client QRCode
}

func (d *defaultQRCode) GetQRCode(imgCode, codeInfo string) (string, error) {
	return d.client.GetQRCode(imgCode, codeInfo)
}
