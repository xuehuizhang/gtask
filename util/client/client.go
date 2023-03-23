package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	GClient *http.Client
)

func InitClient() {
	newTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //TCP超时时间
			KeepAlive: 30 * time.Second, //TCP KeepAlive时间
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,              //最大空闲连接数
		IdleConnTimeout:       90 * time.Second, //空闲连接数超时时间
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	GClient = &http.Client{
		Transport: newTransport,
		Timeout:   60 * time.Second, //http超时时间
	}
}

func Get(url string, params map[string]string) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	//设置header todo

	//设置请求参数
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := GClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func Post(url string, params string) {

}
