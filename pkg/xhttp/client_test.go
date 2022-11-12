package xhttp

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"github.com/weitrue/Seckill/pkg/utils/convert"
	"io"
	"log"
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
	api := "https://www.nswap.network/api/multi-chain/10001/item/refresh"
	client := NewDefaultClient()
	var cr CommonResp
	body := make(map[string]interface{})
	body["chain"] = 1
	body["address"] = "0xDF557AB50FE8476125fE382e55A02Fa7593089ac"
	end := 50728
	for start := 1; start <= end; start++ {
		body["token_id"] = convert.ToString(start)
		byt, _ := json.Marshal(body)
		err := client.Call("POST", api, nil, strings.NewReader(string(byt)), &cr)
		if err != nil {
		}

		t.Log(start, cr.Code, cr.Msg, err)
	}
}

func TestRefreshMetadataFromCsv(t *testing.T) {
	api := "https://www.nswap.network/api/multi-chain/10001/item/refresh"
	client := NewDefaultClient()
	var cr CommonResp
	body := make(map[string]interface{})
	body["chain"] = 1
	body["address"] = "0xF6EAC0712c52433451a6BBCF0eF576fa7e44bCe1"
	result := GetFromCsv()
	for _, tokenId := range result {
		body["token_id"] = tokenId
		byt, _ := json.Marshal(body)
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
