package shift

import (
	"github.com/spaolacci/murmur3"
)

const (
	// INDEX 61 用于截取索引数
	INDEX = 0x0000003D
)

// ShortUrlBase 短地址生成对象
type ShortUrlBase struct {
	StrLen   int    // 生成短地址标识长度
	Shift    int64  // 右移移位数
	Alphabet []byte // 对应进制数字符表
}

// Generator 生成短地址
func (s ShortUrlBase) Generator(longURL string) string {
	hexVal := getMurmur3Str(longURL)

	var index int64
	var tempUri []byte
	for i := 0; i < s.StrLen; i++ {
		index = INDEX & hexVal
		tempUri = append(tempUri, s.Alphabet[index])
		hexVal >>= s.Shift // 右移位操作
	}

	return string(tempUri)
}

// NewShortUrlBase 新建短地址创建对象
func NewShortUrlBase(shift, strLen int64, alphabet string) ShortUrlBase {
	return ShortUrlBase{
		Shift:    shift,
		StrLen:   int(strLen),
		Alphabet: []byte(alphabet),
	}
}

// getMurmur3Str 生成URL的20位校验和
func getMurmur3Str(str string) int64 {
	return int64(murmur3.Sum64([]byte(str)))
}
