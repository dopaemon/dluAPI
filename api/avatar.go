package api

import (
	// "fmt"
	"io"
	"log"
	"net/http"
)

func ApiAvatar(token string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://portal-api.dlu.edu.vn/api/student/GetAvatar", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "portal-api.dlu.edu.vn")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "vi-VN,vi;q=0.9,en;q=0.8,zh-CN;q=0.7,zh;q=0.6")
	req.Header.Set("apikey", "pscRBF0zT2Mqo6vMw69YMOH43IrB2RtXBS0EHit2kzvL2auxaFJBvw==")
	req.Header.Set("authorization", "Bearer " + token)
	req.Header.Set("clientid", "vhu")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://online.dlu.edu.vn")
	req.Header.Set("referer", "https://online.dlu.edu.vn/")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)
	return bodyText
}
