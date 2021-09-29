package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpRequest(method, url, params string, header map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(params)))
	if err != nil {
		panic(err)
	}
	for i, v := range header {
		request.Header.Set(i, v)
	}
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Do(request)
	return response, err
}

func ParseRspToString(successCode int, response *http.Response) ([]byte, error) {
	rspByte, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if response.StatusCode == successCode {
		return rspByte, err
	} else {
		return make([]byte, 0, 0), errors.New(fmt.Sprintf("获取accessToken异常，返回状态码：%d，err message：%s", response.StatusCode,string(rspByte)))
	}
}

func UrlJoin(url string)  string{
	return fmt.Sprintf("http://%s%s",viper.GetString("host"),url)
}
