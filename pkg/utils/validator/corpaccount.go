package validator

import (
	"strings"
)

// CorpAccount 企业对公账户
type CorpAccount string

// NewCorpAccount 新建企业对公账户校验器
func NewCorpAccount(corpaccount string) CorpAccount {
	return CorpAccount(corpaccount)
}

// IsValid 判断是否合法
func (ca CorpAccount) IsValid() bool {
	caStr := strings.ToUpper(string(ca))
	if !corpaccountRegex.Match([]byte(caStr)) {
		return false
	}
	return true
}
