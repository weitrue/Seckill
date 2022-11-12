package shorturl

import (
	"github.com/pkg/errors"
	"github.com/weitrue/Seckill/infrastructure/services/local/shorturl/shift"
)

// ShortUrl 短地址接口
type ShortUrl interface {
	Generator(longURL string) string
}

// Config 短地址配置信息
type Config struct {
	StrLen   int64  // 生成短地址标识长度
	Shift    int64  // 右移移位数
	Alphabet string // 对应进制数字符表
}

// newShortUrl 新建短地址生成对象
func newShortUrl(c *Config) ShortUrl {
	if c == nil {
		return nil
	}

	return shift.NewShortUrlBase(c.Shift, c.StrLen, c.Alphabet)
}

// MustNewShortUrl 新建短地址生成对象
func MustNewShortUrl(c *Config) ShortUrl {
	s := newShortUrl(c)
	if s == nil {
		panic(errors.New("ShortUrlBase: init err"))
	}

	return s
}
