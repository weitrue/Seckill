package vrcode

import (
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/weitrue/Seckill/pkg/xhttp"
)

var vrCode = VRCode{
	domain:   "https://api-java.vrcode.com",
	userName: "testyx1",
	password: "3@saPR93hkd",
	mode:     16,
	client: xhttp.NewClientWithHTTPClient(&http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}),
}

func TestGetQRCode(t *testing.T) {
	dir, _ := os.Getwd()
	dir = strings.ReplaceAll(dir, "pkg/qrcode/vrcode", "docs/images")
	src := dir + "/src.jpg"
	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	qrCode, err := vrCode.GetQRCode(base64.StdEncoding.EncodeToString(bytes), "firstTest")
	if err != nil {
		log.Fatal(err)
	}

	if qrCode != "" {
		bytes, err = base64.StdEncoding.DecodeString(qrCode)
		if err != nil {
			log.Fatal(err)
		}

		vrImages := dir + "/vr.jpg"
		f, _ := os.OpenFile(vrImages, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		_, err = f.Write(bytes)
		if err != nil {
			return
		}
	}

	// _, _ = vrCode.GetQRCode(base64.StdEncoding.EncodeToString(bytes), "firstTest")
}
