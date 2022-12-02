package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/weitrue/Seckill/pkg/utils/convert"
	"io/ioutil"
	"os"
	"strings"
)

type TokenUri struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Attributes  []struct {
		Value     string `json:"value"`
		TraitType string `json:"trait_type"`
	} `json:"attributes"`
}

func ReadCsv(fileName string) map[string][]*TokenUri {
	fileName = "/Users/pony/Projects/go/src/Seckill/utils/check_duplicatesV4.csv"
	ret := make(map[string][]*TokenUri)
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil
	}

	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	fmt.Println(sz)
	count := 0

	// 循环取数据
	for i := 1; i < sz; i++ {
		tokenUrl := &TokenUri{
			Image:       "",
			Name:        strings.TrimSpace(ss[i][0]),
			Description: "",
			Attributes: []struct {
				Value     string `json:"value"`
				TraitType string `json:"trait_type"`
			}{
				{
					TraitType: "color",
					Value:     strings.TrimSpace(ss[i][2]),
				},
			},
		}

		if strings.TrimSpace(ss[i][2]) == "purple" {
			count++
		}

		if d, ok := ret[strings.TrimSpace(ss[i][2])]; ok {
			d = append(d, tokenUrl)
			ret[strings.TrimSpace(ss[i][2])] = d
		} else {
			ret[strings.TrimSpace(ss[i][2])] = []*TokenUri{tokenUrl}
		}
	}

	fmt.Println("=======", count)

	return ret
}

func GenerateData(fileName string) {
	red := []string{}
	blue := []string{}
	green := []string{}
	purple := []string{}
	black := []string{"black_7b458dfa7f574b0652b6a08f02149c7c7.png"}
	gold := []string{"gold_4c8c786f7ef531779a3e8ecbc0f741957.png"}
	silver := []string{"silver_c617582c4f88887cafd79ae34bde2bf87.png"}
	yellow := []string{}
	redMap := make(map[int]struct{})
	blueMap := make(map[int]struct{})
	greenMap := make(map[int]struct{})
	purpleMap := make(map[int]struct{})
	blackMap := make(map[int]struct{})
	goldMap := make(map[int]struct{})
	silverMap := make(map[int]struct{})
	yellowMap := make(map[int]struct{})

	data := ReadCsv(fileName)

	tokenIds := make([]string, 0)
	for k, d := range data {
		fmt.Println(k, "start", len(d))
		if k == "red" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/red/%s", red[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				redMap[i] = struct{}{}
			}
			fmt.Println(k, "ok")
			continue
		}

		if k == "blue" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/blue/%s", blue[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				blueMap[i] = struct{}{}
			}
			fmt.Println(k, "ok")
			continue
		}

		if k == "green" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/green/%s", green[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				greenMap[i] = struct{}{}
			}
			fmt.Println(k, "ok")
			continue
		}

		if k == "purple" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/purple/%s", purple[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				purpleMap[i] = struct{}{}
			}
			fmt.Println(k, "ok")
			continue
		}

		if k == "black" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/black/%s", black[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				blackMap[i] = struct{}{}
			}
			fmt.Println(k, "ok")
			continue
		}
		if k == "gold" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/gold/%s", gold[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				goldMap[i] = struct{}{}
			}
			fmt.Println(k, "ok")
			continue
		}
		if k == "silver" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/silver/%s", silver[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				silverMap[i] = struct{}{}
			}
			fmt.Println(k, "ok")
			continue
		}
		if k == "yellow" {
			for i, v := range d {
				v.Image = fmt.Sprintf("https://s3-nswap-base.nswap.com/luccifiles/yellow/%s", yellow[i])
				fs := strings.Split(v.Name, "#")
				tokenIds = append(tokenIds, fs[1])
				writeFIle(fmt.Sprintf("%s.json", fs[1]), v)
				yellowMap[i] = struct{}{}
			}
			continue
		}

		fmt.Println(d)
	}

	tokenIdsStr := ""
	for _, v := range tokenIds {
		tokenIdsStr += "," + v
	}
	fmt.Println(tokenIdsStr)
}

func writeFIle(fileName string, data *TokenUri) {
	str, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(fmt.Sprintf("/Users/pony/Projects/go/src/Seckill/utils/jsonv4/%s", fileName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(string(str))
		panic(err)
	}
	//及时关闭file句柄
	defer file.Close()
	file.WriteString(string(str))
}

func ReadMetadataCsv() {
	fileName := "/Users/pony/Projects/go/src/Seckill/utils/1669341419139.csv"
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, err := r2.ReadAll()
	if err != nil {
		panic(err)
	}

	tokenMap := make(map[string]struct{})
	// 循环取数据
	for i := 1; i < len(ss); i++ {
		tokenMap[ss[i][2]] = struct{}{}
	}

	for i := 0; i < 1756; i++ {
		if _, ok := tokenMap[convert.ToString(i)]; !ok {
			fmt.Println(i)
		}
	}
}
