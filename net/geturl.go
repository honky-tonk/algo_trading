package net

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

type ProxySet struct {
	Proxyurl string `yaml:"Proxy_Url"`
}

func SetProxy(URL string) *http.Client {
	proxyurl, err := url.Parse(URL)
	if err != nil {
		log.Fatalf("Can't Parse url: %v", err.Error())
	}

	proxy := http.ProxyURL(proxyurl)
	transport := &http.Transport{Proxy: proxy, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &http.Client{Transport: transport}
}

func GetFromUrlByProxy(URL string, ProxyServerUrl string) (*http.Response, error) {
	client := SetProxy(ProxyServerUrl)
	resp, err := client.Get(URL)
	if err != nil {
		return nil, err
	}

	return resp, nil //一定要close当调用者不用的时候
}

func GetFromUrl(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	//defer resp.Body.Close()
	return resp, nil //一定要close当调用者不用的时候
}

func Get(URL string) (*http.Response, error) {
	var proxyset ProxySet

	//read config from local
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	//parse yaml to struct
	if err = yaml.Unmarshal(file, &proxyset); err != nil {
		return nil, err
	}

	//check if use proxy
	if proxyset.Proxyurl == "" { //not fill proxyurl in yaml
		return GetFromUrl(URL)
	} else {
		return GetFromUrlByProxy(URL, proxyset.Proxyurl)
	}
}
