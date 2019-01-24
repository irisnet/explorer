package utils

import (
	"crypto/tls"
	"github.com/irisnet/explorer/backend/logger"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var client http.Client
var timeout = 5 * time.Second
var keepAlive = 30 * time.Second

func init() {
	client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
			}).DialContext,
		},
	}
}

func Get(url string) (bz []byte, err error) {
	start := time.Now()
	logger.Info("http Get url", logger.String("url", url))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error("req error", logger.Any("err", err))
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("req error", logger.Any("err", err))
		return
	}
	if resp.StatusCode != http.StatusOK {
		logger.Error("req error", logger.Any("http_Status", resp.Status))
		return
	}

	bz, err = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err != nil {
		logger.Error("ioutil.ReadAll err", logger.Any("io", err))
	}
	end := time.Now()
	logger.Info("http get coast second", logger.Int64("time", end.Unix()-start.Unix()))
	return
}

func Forward(req *http.Request, url string) (bz []byte, err error) {
	r := forkRequest(req, url)
	res, err := client.Do(r)
	if err != nil || res.StatusCode != 200 {
		logger.Error("Forward err", logger.String("err", err.Error()))
		return bz, err
	}

	defer res.Body.Close()

	bz, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("Forward err", logger.String("err", err.Error()))
		return bz, err
	}
	return bz, err
}

func forkRequest(req *http.Request, url string) *http.Request {
	r, _ := http.NewRequest(req.Method, url, req.Body)
	for k, v := range req.Header {
		for _, vv := range v {
			r.Header.Add(k, vv)
		}
	}
	r.RemoteAddr = req.RemoteAddr
	r.Form = req.Form
	r.PostForm = req.PostForm
	r.MultipartForm = req.MultipartForm
	r.Proto = req.Proto
	r.ProtoMajor = req.ProtoMajor
	r.ProtoMinor = req.ProtoMinor
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return r
}

func GetIpAddr(req *http.Request) string {
	logger.Info("http header", logger.Any("header", req.Header))
	addrs := req.Header["X-Forwarded-For"]
	if len(addrs) > 0 && "unknown" != addrs[0] {

		return addrs[0]
	}

	addrs = req.Header["Proxy-Client-IP"]
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	addrs = req.Header["WL-Proxy-Client-IP"]
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	addrs = req.Header["X-Real-IP"]
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	addrs = strings.Split(req.RemoteAddr, ":")
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	return ""
}
