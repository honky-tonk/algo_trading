package net

import (
	"crypto/tls"
	"fmt"
	"io"
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

func GetFromUrlByProxy(URL string, ProxyServerUrl string) {
	client := SetProxy(ProxyServerUrl)
	resp, err := client.Get(URL)
	if err != nil {
		log.Fatalf("Get URL error from GetFromUrlByProxy: %v", err.Error())
	}
	defer resp.Body.Close()

	bodybyte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read from body error: %v", err.Error())
	}

	fmt.Println(string(bodybyte))
}

func GetFromUrl(URL string) {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Get URL error from GetFromUrlByProxy: %v", err.Error())
	}
	defer resp.Body.Close()

	bodybyte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read from body error: %v", err.Error())
	}

	fmt.Println(string(bodybyte))
}

func Get(URL string) {
	var proxyset ProxySet

	//read config from local
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Can't read config.yaml file: %v", err.Error())
	}

	//parse yaml to struct
	if err = yaml.Unmarshal(file, &proxyset); err != nil {
		log.Fatalf("Can't unmarshal yaml file to struct: %v", err.Error())
	}

	//check if use proxy
	if proxyset.Proxyurl == "" { //not fill proxyurl in yaml
		GetFromUrl(URL)
	} else {
		GetFromUrlByProxy(URL, proxyset.Proxyurl)
	}
}
