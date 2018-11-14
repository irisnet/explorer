package utils

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"
	"net/http"
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
	log.Println("http Get url:" + url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("req error:" + err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("do error,err:" + err.Error())
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("do Get err:" + resp.Status)
		return
	}

	bz, err = ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err != nil {
		log.Println("ioutil.ReadAll err:" + err.Error())
	}

	return
}
