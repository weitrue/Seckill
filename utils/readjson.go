package utils

import (
	"encoding/json"
	"fmt"
	"github.com/weitrue/Seckill/pkg/evm/eip"
	"github.com/weitrue/Seckill/pkg/utils/excel"
	"github.com/weitrue/Seckill/pkg/utils/math"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

func ReadSnapshot() map[string]int64 {
	filePtr, err := os.OpenFile("/Users/pony/Projects/go/src/Seckill/utils/snapshot-1665705774.json", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return nil
	}
	defer filePtr.Close()
	info := make(map[string]int64)
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	count := int64(0)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		for _, i := range info {
			count += i
		}
		fmt.Println("解码成功", count)
	}

	return info
}

func ReadWhitelist() map[string]int64 {
	filePtr, err := os.OpenFile("/Users/pony/Projects/go/src/Seckill/utils/whitelist.json", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return nil
	}
	defer filePtr.Close()
	info := make(map[string]int64)
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	count := int64(0)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		for _, i := range info {
			count += i
		}
		fmt.Println("解码成功", count)
	}

	return info
}

func ReadTargetSnapshot() map[string]int64 {
	filePtr, err := os.OpenFile("/Users/pony/Projects/go/src/Seckill/utils/snapshot.prod.json", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return nil
	}
	defer filePtr.Close()
	info := make(map[string]int64)
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	count := int64(0)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		for _, i := range info {
			count += i
		}
		fmt.Println("解码成功", count)
	}

	return info
}

func GenerateTargetList(total int64) {
	snapshot := ReadSnapshot()
	whitelist := ReadWhitelist()
	whitelistNum := int64(0)
	for _, v := range whitelist {
		whitelistNum += v
	}

	targetMap := whitelist
	randomList := make([]string, 0)
	for k, v := range snapshot {
		if _, ok := whitelist[k]; !ok {
			for ; v > 0; v-- {
				randomList = append(randomList, k)
			}
		}
	}

	for i := int64(0); i < total-whitelistNum; i++ {
		fmt.Println(len(randomList))
		rand := math.RandInt(0, len(randomList))
		if v, ok := targetMap[randomList[rand]]; ok {
			targetMap[randomList[rand]] = v + 1
		} else {
			targetMap[randomList[rand]] = 1
		}

		randomList = append(randomList[:rand], randomList[rand+1:]...)
	}

	fmt.Println(len(targetMap))
	whitelistNum = 0
	for _, v := range targetMap {
		whitelistNum += v
	}
	fmt.Println(whitelistNum)
	str, err := json.Marshal(targetMap)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("./snapshot.prod.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(string(str))
		panic(err)
	}
	//及时关闭file句柄
	defer file.Close()
	file.WriteString(string(str))
}

func CheckTargetList() {
	snapshot := ReadTargetSnapshot()
	whitelist := ReadWhitelist()

	count := make([]string, 0)
	for k, _ := range whitelist {
		if _, ok := snapshot[k]; !ok {
			count = append(count, k)
		}
	}

	fmt.Println(count)
}

func WriteJsonFromXlsx() {
	//list1 := GetIn()
	list1 := make(map[string]int64)
	list2 := GetOut()
	fmt.Println("Get from excel", len(list2))
	for k, v := range list2 {
		_, err := eip.ToCheckSumAddress(k)
		if err != nil {
			fmt.Println("Invalid address", k)
			continue
		}
		if count, ok := list1[k]; ok {
			list1[k] = v + count
		} else {
			list1[k] = v
		}
	}

	fmt.Println(len(list1))
	whitelistNum := int64(0)
	for k, v := range list1 {
		whitelistNum += v

		if v > 1 {
			fmt.Println("Mint greater than 1,", k)
		}
	}
	fmt.Println("Mint num:", whitelistNum)
	str, err := json.Marshal(list2)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("/Users/pony/Projects/go/src/Seckill/utils/snapshot.add.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(string(str))
		panic(err)
	}
	//及时关闭file句柄
	defer file.Close()
	file.WriteString(string(str))
}

func GetIn() map[string]int64 {
	f, err := os.Open("/Users/pony/Projects/go/src/Seckill/utils/nin.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := make(map[string]int64)
	err = excel.ReadRows(f, "in", func(rowNum int, columns []string) bool {
		if rowNum == 35 {
			return false
		}

		for _, column := range columns {
			if len(column) == 0 {
				continue
			}

			if strings.HasPrefix(column, "0x") {
				if v, ok := result[column]; ok {
					result[column] = v + 1
				} else {
					result[column] = 1
				}
			}
		}

		return true
	})
	if err != nil {
		panic(err)
	}

	return result
}

func GetOut() map[string]int64 {
	f, err := os.Open("/Users/pony/Projects/go/src/Seckill/utils/nout.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := make(map[string]int64)
	err = excel.ReadRows(f, "out", func(rowNum int, columns []string) bool {
		if rowNum == 1521 {
			return false
		}

		for _, column := range columns {
			if len(column) == 0 {
				continue
			}

			if strings.HasPrefix(column, "0x") {
				if v, ok := result[column]; ok {
					result[column] = v + 1
				} else {
					result[column] = 1
				}
			}
		}

		return true
	})
	if err != nil {
		panic(err)
	}

	return result
}

func WritePubPriFromXlsx() {
	list := GetPubPri()
	fmt.Println(len(list))
	public := make([]string, 0)
	private := make([]string, 0)
	str := ""
	for k, v := range list {
		public = append(public, k)
		private = append(private, v)
		str += "'" + k + "',"
	}

	fmt.Println(str)
	//file1, err := os.OpenFile("/Users/pony/Projects/go/src/Seckill/utils/snapshot.transfer.public.json", os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(len(public))
	//publicByt, _ := json.Marshal(public)
	////及时关闭file句柄
	//defer file1.Close()
	//file1.WriteString(string(publicByt))
	//
	//file2, err := os.OpenFile("/Users/pony/Projects/go/src/Seckill/utils/snapshot.transfer.private.json", os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(len(private))
	//privateByt, _ := json.Marshal(private)
	////及时关闭file句柄
	//defer file2.Close()
	//file2.WriteString(string(privateByt))
}

func GetPubPri() map[string]string {
	f, err := os.Open("/Users/pony/Projects/go/src/Seckill/utils/nout.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := make(map[string]string)
	err = excel.ReadRows(f, "out", func(rowNum int, columns []string) bool {
		if rowNum == 73 {
			return false
		}

		add, _ := eip.ToCheckSumAddress(columns[0])
		result[add] = columns[1]
		return true
	})
	if err != nil {
		panic(err)
	}

	return result
}

func GetTransTarget() map[string]int64 {
	f, err := os.Open("/Users/pony/Projects/go/src/Seckill/utils/nout.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := make(map[string]int64)
	err = excel.ReadRows(f, "out", func(rowNum int, columns []string) bool {
		if rowNum == 73 {
			return false
		}

		add, _ := eip.ToCheckSumAddress(columns[0])

		result[add] = 1
		return true
	})
	if err != nil {
		panic(err)
	}

	return result
}

func CreateXlsx() {
	f1 := excelize.NewFile()
	defer f1.Close()
	in := f1.NewSheet("in")
	f1.SetActiveSheet(in)
	if err := f1.SaveAs("/Users/pony/Projects/go/src/Seckill/utils/nin.xlsx"); err != nil {
		fmt.Println(err)
	}

	f2 := excelize.NewFile()
	defer f2.Close()
	out := f2.NewSheet("out")
	f2.SetActiveSheet(out)
	if err := f2.SaveAs("/Users/pony/Projects/go/src/Seckill/utils/nout.xlsx"); err != nil {
		fmt.Println(err)
	}
}
