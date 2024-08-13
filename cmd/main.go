package main

import (
	"algo_trading/pkg/assets"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	stock := assets.ParseConfig()

	fmt.Println("in main: ", stock)

	// 设置代理
	proxyURL, _ := url.Parse("http://192.168.152.1:7890")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: transport,
	}

	// 发送请求
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// 读取响应的 body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取 body 失败: %v\n", err)
		return
	}

	// 打印响应的 body
	fmt.Println(string(body))

	// url := "http://www.google.com"

	// os.Setenv("HTTP_PROXY", "http://127.0.0.1:9090")
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// // 检查响应状态码
	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
	// 	return
	// }

	// // 读取响应的 body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("读取 body 失败: %v\n", err)
	// 	return
	// }

	// // 打印响应的 body
	// fmt.Println(string(body))

	//fmt.Println("assets is ", config.Assets, " ", config.Assets.Datas[1].Name, "'s time interval is ", config.Assets.Datas[1].Time_Interval)

}
