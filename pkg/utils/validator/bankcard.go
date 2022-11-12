package validator

import (
	"strconv"
	"strings"
)

// 发卡行识别码 ISO/IEC 7812
// 详情信息：https://zh.wikipedia.org/wiki/%E5%8F%91%E5%8D%A1%E8%A1%8C%E8%AF%86%E5%88%AB%E7%A0%81
// 校验规则: https://zh.wikipedia.org/wiki/Luhn%E7%AE%97%E6%B3%95
// 参考标准：https://zh.wikipedia.org/wiki/ISO/IEC_7812

// 数值-sum(数值 * 2) map
// 其中 key 表示银行卡卡号的数值，value 表示该数值 * 2 后，各位相加后的数值
var valueDoubleSumMap = map[int]int{
	0: 0, // 0 * 2 = 0, 0 = 0
	1: 2, // 1 * 2 = 2, 2 = 2
	2: 4, // 2 * 2 = 4, 4 = 4
	3: 6, // 3 * 2 = 6, 6 = 6
	4: 8, // 4 * 2 = 8, 8 = 8
	5: 1, // 5 * 2 = 10, 1 + 0 = 1
	6: 3, // 6 * 2 = 12, 1 + 2 = 3
	7: 5, // 7 * 2 = 14, 1 + 4 = 5
	8: 7, // 8 * 2 = 16, 1 + 6 = 7
	9: 9, // 9 * 2 = 18, 1+ 8 = 9
}

// BankCard 发卡行识别码
type BankCard string

// NewBankCard 新建发卡行识别码校验器
func NewBankCard(bankcard string) BankCard {
	return BankCard(bankcard)
}

// IsValid 判断是否合法
func (bc BankCard) IsValid() bool {
	bcStr := strings.ToUpper(string(bc))
	if !bankcardRegex.Match([]byte(bcStr)) {
		return false
	}
	sum := 0
	length := len(bcStr)
	for index := length - 1; index >= 0; index-- {
		if value, err := strconv.Atoi(string(bcStr[index])); err == nil {
			if (length-index)%2 == 0 {
				sum += valueDoubleSumMap[value]
			} else {
				sum += value
			}
		} else {
			return false
		}
	}
	return sum%10 == 0
}
