/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/12 上午11:50
 * Description: token 鉴权实现
 *   - AES 对 user.Info 做对称加解密 + base64 包装
 *   - Verify(ctx, token) 解包 token 得到 *user.Info,过期返回错误
 *   - Login(ctx, uid, passwd) 构造 Info 并签成 token(密码当前未校验,留给 user.Service.Login)
 *
 *   注意: 当前 AES 密钥硬编码,仅为搬迁最小改动;生产应从配置读取,后续迁移到 JWT 或可轮换的密钥对
 **/

package auth

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/internal/domain/user"
)

// authKey AES 密钥(当前硬编码,TODO: 从配置读取或迁移 JWT)
var authKey = []byte("seckill2021")

// 默认 token 有效期(秒)
const defaultTokenTTL int64 = 24 * 3600

// 常见错误
var (
	ErrEmptyToken   = errors.New("auth: empty token")
	ErrInvalidToken = errors.New("auth: invalid token")
	ErrExpiredToken = errors.New("auth: token expired")
)

/*Verify
 *@Description: 验证 token 并解出用户信息
 *@param ctx 请求上下文(当前未使用,预留链路追踪)
 *@param token base64 编码的 AES 密文
 *@return *user.Info 用户信息(验证失败返回 nil)
 *@return error
 */
func Verify(ctx context.Context, token string) (*user.Info, error) {
	_ = ctx // 预留,后续接入 trace 时使用
	if token == "" {
		return nil, ErrEmptyToken
	}
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("logType:AuthVerify, panic:%v, token:%s", r, token)
		}
	}()
	cipher, err := aes.NewCipher(padding(authKey, 16))
	if err != nil {
		return nil, err
	}
	src, err := base64.StdEncoding.DecodeString(token)
	if err != nil || len(src) == 0 {
		return nil, ErrInvalidToken
	}
	src = padding(src, cipher.BlockSize())
	output := make([]byte, len(src))
	cipher.Decrypt(output, src)
	var info *user.Info
	if err := json.Unmarshal(output, &info); err != nil || info == nil {
		return nil, ErrInvalidToken
	}
	if info.ExpireTime < time.Now().Unix() {
		return nil, ErrExpiredToken
	}
	return info, nil
}

/*Login
 *@Description: 用 uid 签发 token(当前不校验密码,预留给 user.Service.Login 后续补充校验)
 *@param ctx 请求上下文(当前未使用,预留)
 *@param uid 用户 ID
 *@param passwd 密码(当前未校验)
 *@return *user.Info
 *@return string token
 *@return error
 */
func Login(ctx context.Context, uid, passwd string) (*user.Info, string, error) {
	_ = ctx
	_ = passwd // TODO: 对接 user.Service.Login 时校验
	info := &user.Info{
		UID:        uid,
		LoginTime:  time.Now().Unix(),
		ExpireTime: time.Now().Unix() + defaultTokenTTL,
	}
	data, err := json.Marshal(info)
	if err != nil {
		return nil, "", err
	}
	cipher, err := aes.NewCipher(padding(authKey, 16))
	if err != nil {
		return nil, "", err
	}
	data1 := padding(data, cipher.BlockSize())
	dst := make([]byte, len(data1))
	cipher.Encrypt(dst, data1)
	return info, base64.StdEncoding.EncodeToString(dst[:len(data)]), nil
}

// padding 对输入按 blkSize 右侧零填充(AES 要求块对齐)
func padding(src []byte, blkSize int) []byte {
	l := len(src)
	for i := 0; i < blkSize-l%blkSize; i++ {
		src = append(src, byte(0))
	}
	return src
}
