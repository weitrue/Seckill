package validator

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// 公民身份号码 GB 11643-1999
// 校验规则: http://openstd.samr.gov.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E
// 在线预览：https://zh.wikisource.org/wiki/GB_11643-1999_%E5%85%AC%E6%B0%91%E8%BA%AB%E4%BB%BD%E5%8F%B7%E7%A0%81

// 索引-权重 map
// 其中 key 表示身份证号码数字的索引，value 表示该索引对应的权重
// value = 2 ** （key） % 11
var idCardIndexWeightMap = map[int]int{
	0:  7,
	1:  9,
	2:  10,
	3:  5,
	4:  8,
	5:  4,
	6:  2,
	7:  1,
	8:  6,
	9:  3,
	10: 7,
	11: 9,
	12: 10,
	13: 5,
	14: 8,
	15: 4,
	16: 2,
}

// 加权求和模-校验码 map
// 其中 key 表示加权求和后对 11 取的模，value 表示校验码
var modCheckCodeMap = map[int]uint8{
	0:  '1',
	1:  '0',
	2:  'X',
	3:  '9',
	4:  '8',
	5:  '7',
	6:  '6',
	7:  '5',
	8:  '4',
	9:  '3',
	10: '2',
}

// IdCard 公民身份号码
type IdCard string

// NewIdCard 新建公民身份号码校验器
func NewIdCard(idcard string) IdCard {
	return IdCard(idcard)
}

// IsValid 判断是否合法
func (ic IdCard) IsValid() bool {
	icStr := strings.ToUpper(string(ic))
	if !idcardRegex.Match([]byte(icStr)) {
		return false
	}
	sum := 0
	checkCode := icStr[17]
	for index := range icStr[:17] {
		if a, err := strconv.Atoi(string(icStr[index])); err == nil {
			// 计算加权因子
			w := idCardIndexWeightMap[index]
			// 计算加权和
			sum += a * w
		} else {
			return false
		}
	}
	// fmt.Println(string(modCheckCodeMap[sum%11]), string(checkCode))
	return modCheckCodeMap[sum%11] == checkCode
}

// GetBirthday 获取生日日期
func (ic IdCard) GetBirthday() (time.Time, error) {
	if !ic.IsValid() {
		return time.Time{}, errors.New("validator: invalid idcard")
	}
	icStr := string(ic)
	yearStr, monthStr, dayStr := icStr[6:10], icStr[10:12], icStr[12:14]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, err
	}
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return time.Time{}, err
	}
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

const (
	// Female 女性
	Female = 0
	// Male 男性
	Male = 1
)

// GetGender 获取性别
func (ic IdCard) GetGender() (int, error) {
	if !ic.IsValid() {
		return 0, errors.New("validator: invalid idcard")
	}
	icStr := string(ic)
	numStr := icStr[14:17]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}
	return num % 2, nil
}

// IsMale 判断是否为男性
func (ic IdCard) IsMale() (bool, error) {
	gender, err := ic.GetGender()
	if err != nil {
		return false, err
	}
	return gender == Male, nil
}

// IsFemale 判断是否为女性
func (ic IdCard) IsFemale() (bool, error) {
	gender, err := ic.GetGender()
	if err != nil {
		return false, err
	}
	return gender == Female, nil
}
