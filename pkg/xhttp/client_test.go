package xhttp

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/weitrue/Seckill/pkg/utils/convert"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlValuesAdd(t *testing.T) {
	values := make(url.Values)
	values.Add("bankcard", "bankcard1")
	values.Add("idcard", "idcard")
	values.Add("mobile", "mobile")
	values.Add("name", "name")
	values.Add("bankcard", "bankcard2")

	assert.Equal(t, "bankcard=bankcard1&bankcard=bankcard2&idcard=idcard&mobile=mobile&name=name", values.Encode())
}

func TestUrlValuesSet(t *testing.T) {
	values := make(url.Values)
	values.Set("bankcard", "bankcard1")
	values.Set("idcard", "idcard")
	values.Set("mobile", "mobile")
	values.Set("name", "name")
	values.Set("bankcard", "bankcard2")

	assert.Equal(t, "bankcard=bankcard2&idcard=idcard&mobile=mobile&name=name", values.Encode())
}

func TestClient(t *testing.T) {
	type args struct {
		method string
		rawurl string
		header map[string]string
		body   io.Reader
	}
	cases := []struct {
		title string
		args  args
	}{
		{
			title: "do get method",
			args: args{
				method: "GET",
				rawurl: "https://www.httpbin.org/get",
				header: map[string]string{
					"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
					"User-Agent":      "Go-HTTP-Request",
				},
				body: nil,
			},
		},
		{
			title: "do post method",
			args: args{
				method: "POST",
				rawurl: "https://www.httpbin.org/post",
				header: map[string]string{
					"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
					"User-Agent":      "Go-HTTP-Request",
				},
				body: strings.NewReader("a=b&c=d"),
			},
		},
	}

	client := NewDefaultClient()
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			req, err := client.GetRequest(c.args.method, c.args.rawurl, c.args.header, c.args.body)
			assert.NoError(t, err)
			_, resp, err := client.GetResponse(req)
			assert.NoError(t, err)
			t.Log(string(resp))
		})
	}
}

// CommonResp 通用响应
type CommonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// FileInfo 文件响应
type FileInfo struct {
	Name       string `json:"name"`        // 文件名称
	Type       string `json:"type"`        // 文件类型
	Hash       string `json:"hash"`        // 文件hash（sha256(文件内容+文件名称)）
	Url        string `json:"url"`         // 文件url
	Extra      string `json:"extra"`       // 额外信息
	CreateTime int64  `json:"create_time"` // 创建时间
	UpdateTime int64  `json:"update_time"` // 更新时间
}

// FileChunkInfo 文件分块响应
type FileChunkInfo struct {
	CurrentSeq  int64  `json:"current_seq"`  // 当前块序号
	CurrentSize int64  `json:"current_size"` // 当前块容量
	FileName    string `json:"file_name"`    // 文件名称
	FileHash    string `json:"file_hash"`    // 文件hash（sha256(文件内容+文件名称)）
	TotalSeq    int64  `json:"total_seq"`    // 总序号
	TotalSize   int64  `json:"total_size"`   // 总容量
}

// CheckFileChunkResp 检查文件分块响应
type CheckFileChunkResp struct {
	IsDone     bool             `json:"is_done"`     // 是否完成（若文件分块已完成，返回文件信息；否则返回已上传成功的文件分块列表）
	FileHash   string           `json:"file_hash"`   // 文件hash（sha256(文件内容+文件名称)）
	FileInfo   *FileInfo        `json:"file_info"`   // 文件信息
	FileChunks []*FileChunkInfo `json:"file_chunks"` // 已上传成功的文件分块列表
}

// UploadFileChunkReq 上传文件分块请求
type UploadFileChunkReq struct {
	CurrentData string `json:"current_data" validate:"required" label:"当前块数据"`         // 当前块数据（base64编码）
	CurrentSeq  int64  `json:"current_seq" validate:"required" label:"当前块序号"`          // 当前块序号
	CurrentSize int64  `json:"current_size" validate:"required" label:"当前块大小"`         // 当前块大小
	FileName    string `json:"file_name" validate:"required" label:"文件名称"`             // 文件名称
	FileHash    string `json:"file_hash" validate:"len=64,hexadecimal" label:"文件hash"` // 文件hash（sha256(文件内容+文件名称)）
	TotalSeq    int64  `json:"total_seq" validate:"required" label:"总序号"`              // 总序号
	TotalSize   int64  `json:"total_size" validate:"required" label:"总大小"`             // 总大小
}

// UploadFileChunkResp 上传文件分块响应
type UploadFileChunkResp struct {
	CurrentSeq int64  `json:"current_seq"` // 当前块序号
	FileHash   string `json:"file_hash"`   // 文件hash（sha256(文件内容+文件名称)）
}

// MergeFileChunkReq 合并文件分块请求
type MergeFileChunkReq struct {
	FileName string `json:"file_name" validate:"required" label:"文件名称"`             // 文件名称
	FileHash string `json:"file_hash" validate:"len=64,hexadecimal" label:"文件hash"` // 文件hash（sha256(文件内容+文件名称)）
}

// MergeFileChunkResp 上传文件分块响应
type MergeFileChunkResp struct {
	FileHash string `json:"file_hash"` // 文件hash（sha256(文件内容+文件名称)）
	Extra    string `json:"extra"`     // 额外信息
}

func TestClient_Call(t *testing.T) {
	api := "http://172.16.101.87:46792/v2"
	client := NewDefaultClient()

	f, err := os.Open("./testdata/test.pdf")
	assert.NoError(t, err)
	defer f.Close()

	fi, err := f.Stat()
	assert.NoError(t, err)

	s := sha256.New()
	_, err = io.Copy(s, f)
	assert.NoError(t, err)

	s.Write([]byte(fi.Name()))
	sha256Hash := hex.EncodeToString(s.Sum(nil))
	t.Log(sha256Hash, fi.Name())

	var cr CommonResp
	var checkFileChunkResp CheckFileChunkResp
	cr.Data = &checkFileChunkResp

	err = client.Call("GET", api+"/file/chunk/check?file_hash="+sha256Hash, nil, nil, &cr)
	assert.NoError(t, err)
	t.Log(cr, checkFileChunkResp)

	if !checkFileChunkResp.IsDone {
		_, err = f.Seek(0, 0)
		assert.NoError(t, err)

		count := 1
		chunkSize := 4 << 20 // 4MB
		buf := make([]byte, chunkSize)
		totalSize := fi.Size()
		totalSeq := (fi.Size() / int64(chunkSize)) + 1
		for {
			n, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				assert.NoError(t, err)
			}

			uploadFileReq := &UploadFileChunkReq{
				CurrentData: base64.StdEncoding.EncodeToString(buf[:n]),
				CurrentSeq:  int64(count),
				CurrentSize: int64(n),
				FileName:    fi.Name(),
				FileHash:    sha256Hash,
				TotalSeq:    totalSeq,
				TotalSize:   totalSize,
			}
			data, err := json.Marshal(uploadFileReq)
			assert.NoError(t, err)

			var uploadFileResp UploadFileChunkResp
			cr.Data = &uploadFileResp

			err = client.Call("POST", api+"/file/chunk/upload",
				map[string]string{"Content-Type": "application/json"}, bytes.NewReader(data), &cr)
			assert.NoError(t, err)
			t.Log(cr, uploadFileResp)
			count++
		}

		mergeFileChunkReq := &MergeFileChunkReq{
			FileName: fi.Name(),
			FileHash: sha256Hash,
		}
		data, err := json.Marshal(mergeFileChunkReq)
		assert.NoError(t, err)

		var mergeFileChunkResp MergeFileChunkResp
		cr.Data = &mergeFileChunkResp

		err = client.Call("POST", api+"/file/chunk/merge",
			map[string]string{"Content-Type": "application/json"}, bytes.NewReader(data), &cr)
		assert.NoError(t, err)
		t.Log(cr, mergeFileChunkResp)
	}
}

func TestRefreshMetadata(t *testing.T) {
	api := "https://dev.nswap.works/api/multi-chain/56/collection/refresh/0xAeef2b7D89b609b2890e10413D13177fC9301Ee6"
	var cr CommonResp
	body := make(map[string]interface{})
	body["chain"] = 1
	body["address"] = "0xDF557AB50FE8476125fE382e55A02Fa7593089ac"
	tokenIds := []int{108, 139, 155, 260, 274, 369, 412, 418, 467, 613, 621, 640, 722, 737, 745, 756, 701, 858, 862, 714, 735, 907, 908, 914, 893, 171, 191, 193, 305, 309, 317, 511, 528, 607, 717, 721, 952, 981, 1006, 809, 1126, 1140, 1161, 78, 92, 223, 231, 351, 352, 354, 355, 371, 904, 1099, 1191, 73, 93, 280, 296, 798, 940, 1101, 766, 196, 372, 731, 746, 720, 982, 134, 150, 262, 273, 214, 228, 303, 306, 368, 656, 661, 620, 912, 922, 672, 715, 1051, 1225, 1227, 810, 815, 39, 56, 167, 170, 204, 208, 361, 367, 545, 551, 562, 602, 605, 614, 556, 558, 742, 757, 869, 897, 900, 1078, 1134, 1153, 1245, 66, 79, 82, 729, 823, 969, 990, 1172, 363, 698, 540, 818, 548, 345, 901, 831, 1226, 19, 29, 140, 154, 156, 162, 88, 283, 225, 310, 313, 477, 482, 337, 344, 346, 517, 356, 357, 530, 485, 700, 505, 506, 739, 871, 874, 632, 970, 1014, 1028, 1171, 761, 1194, 45, 58, 216, 226, 314, 376, 563, 617, 550, 557, 903, 913, 622, 630, 1068, 1080, 1114, 1221, 1224, 75, 83, 377, 378, 993, 994, 655, 30, 141, 168, 190, 197, 320, 325, 340, 349, 480, 504, 507, 669, 688, 693, 702, 894, 899, 917, 995, 1005, 1027, 1130, 207, 217, 220, 365, 230, 510, 514, 515, 523, 547, 629, 1037, 1049, 1163, 80, 222, 233, 604, 627, 728, 976, 1092, 1248, 623, 133, 138, 144, 166, 172, 182, 279, 281, 210, 213, 439, 336, 343, 444, 456, 518, 531, 634, 538, 542, 765, 767, 866, 876, 822, 886, 919, 928, 949, 719, 67, 205, 308, 516, 522, 527, 544, 730, 755, 827, 91, 97, 353, 366, 816, 819, 820, 826}
	for _, start := range tokenIds {
		//body["token_id"] = convert.ToString(start)
		//byt, _ := json.Marshal(body)
		client := NewClientWithHTTPClient(http.DefaultClient)
		err := client.Call("POST", fmt.Sprintf("%s/%d", api, start), nil, nil, &cr)
		if err != nil {
		}

		t.Log(start, cr.Code, cr.Msg, err)
	}
}

func TestRefreshMetadataByConfig(t *testing.T) {
	var cr CommonResp
	body := make(map[string]interface{})
	body["chain"] = 1
	body["address"] = "0xDF557AB50FE8476125fE382e55A02Fa7593089ac"
	//tokenIds := []int{108, 139, 155, 260, 274, 369, 412, 418, 467, 613, 621, 640, 722, 737, 745, 756, 701, 858, 862, 714, 735, 907, 908, 914, 893, 171, 191, 193, 305, 309, 317, 511, 528, 607, 717, 721, 952, 981, 1006, 809, 1126, 1140, 1161, 78, 92, 223, 231, 351, 352, 354, 355, 371, 904, 1099, 1191, 73, 93, 280, 296, 798, 940, 1101, 766, 196, 372, 731, 746, 720, 982, 134, 150, 262, 273, 214, 228, 303, 306, 368, 656, 661, 620, 912, 922, 672, 715, 1051, 1225, 1227, 810, 815, 39, 56, 167, 170, 204, 208, 361, 367, 545, 551, 562, 602, 605, 614, 556, 558, 742, 757, 869, 897, 900, 1078, 1134, 1153, 1245, 66, 79, 82, 729, 823, 969, 990, 1172, 363, 698, 540, 818, 548, 345, 901, 831, 1226, 19, 29, 140, 154, 156, 162, 88, 283, 225, 310, 313, 477, 482, 337, 344, 346, 517, 356, 357, 530, 485, 700, 505, 506, 739, 871, 874, 632, 970, 1014, 1028, 1171, 761, 1194, 45, 58, 216, 226, 314, 376, 563, 617, 550, 557, 903, 913, 622, 630, 1068, 1080, 1114, 1221, 1224, 75, 83, 377, 378, 993, 994, 655, 30, 141, 168, 190, 197, 320, 325, 340, 349, 480, 504, 507, 669, 688, 693, 702, 894, 899, 917, 995, 1005, 1027, 1130, 207, 217, 220, 365, 230, 510, 514, 515, 523, 547, 629, 1037, 1049, 1163, 80, 222, 233, 604, 627, 728, 976, 1092, 1248, 623, 133, 138, 144, 166, 172, 182, 279, 281, 210, 213, 439, 336, 343, 444, 456, 518, 531, 634, 538, 542, 765, 767, 866, 876, 822, 886, 919, 928, 949, 719, 67, 205, 308, 516, 522, 527, 544, 730, 755, 827, 91, 97, 353, 366, 816, 819, 820, 826}
	//tokenIdsV3 := []int{825, 926, 930, 830, 932, 945, 927, 882}
	tokenIdV4 := []int{7, 1015, 1028}
	for _, start := range tokenIdV4 {
		//body["token_id"] = convert.ToString(start)
		//byt, _ := json.Marshal(body)
		client := NewClientWithHTTPClient(http.DefaultClient)
		err := client.Call("POST", fmt.Sprintf("https://www.nswap.works/api/multi-chain/56/item/lucci/%d/refresh", start), nil, nil, &cr)
		if err != nil {
		}

		t.Log(start, cr.Code, cr.Msg, err)
	}
}

func TestRefreshMetadata2(t *testing.T) {
	api := "https://www.nswap.network/api/multi-chain/10001/item/refresh"
	var cr CommonResp
	body := make(map[string]interface{})
	body["chain"] = 1
	body["address"] = "0xd06521CAD27513a8E635e9933C94Cd9F84472914"
	end := 1810
	for start := 840; start <= end; start++ {
		body["token_id"] = convert.ToString(start)
		byt, _ := json.Marshal(body)
		client := NewClientWithHTTPClient(http.DefaultClient)
		err := client.Call("POST", api, nil, strings.NewReader(string(byt)), &cr)
		if err != nil {
		}

		t.Log(start, cr.Code, cr.Msg, err)
	}
}

func TestRefreshMetadataFromCsv(t *testing.T) {
	api := "https://www.nswap.works/api/multi-chain/56/item/refresh"
	//api := "https://www.nswap.works/api/v2/launchpad/refresh/0xfE542c1FEcc88f1d8B0E53171B8E4d486391fb98"
	var cr CommonResp
	body := make(map[string]interface{})
	body["chain"] = 1
	body["address"] = "0xBC867E76D1360b2D21eB1C2753f94b1C103415e1"
	result := GetFromCsv()
	for _, tokenId := range result {
		body["token_id"] = tokenId
		byt, _ := json.Marshal(body)
		client := NewDefaultClient()
		_ = client.Call("POST", api, nil, strings.NewReader(string(byt)), &cr)

		t.Log(tokenId, cr.Code, cr.Msg)
	}
}

func GetFromCsv() []string {
	f, err := os.Open("/Users/pony/Projects/go/src/Seckill/pkg/xhttp/1667284580042.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true
	result := make([]string, 0)
	//针对大文件，一行一行的读取文件
	content, err := r.ReadAll()
	if err != nil && err != io.EOF {
		log.Fatalf("can not read, err is %+v", err)
	}
	for _, c := range content {
		if !strings.Contains(c[0], "token_id") {
			result = append(result, c[0])
		}
	}

	return result
}

type LeggiMsg struct {
	Address  string   `json:"address"`
	TokenIDs []string `json:"tokenIDs"`
	ChainID  int64    `json:"chainID"`
}

func TestName(t *testing.T) {

	leggiMsg := LeggiMsg{
		Address:  "0xfE542c1FEcc88f1d8B0E53171B8E4d486391fb98",
		TokenIDs: []string{"1", "2", "3", "4", "5", "6"},
		ChainID:  56,
	}

	byt, err := json.Marshal(leggiMsg)
	if err != nil {
		t.Log(err)
	}

	privateKey, err := crypto.HexToECDSA("ad37fd4de7ef09a2fc47cc3a6041e28415108cd08608519035af47ec86958eb5")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(byt)
	fmt.Println(hash.Hex()) // 0xeaf5dbb5ba2041c036919d3e00d112693bbede54d745230cc4d548316371a760

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x751096364b009ae90540368095310e273965e583f48a2b5a4c5d723618b861e24a528564fb6e355e596c8977c0b9ba522cf2ceb142728926887235273757e7da01
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified, string(signatureNoRecoverID)) // true
}
