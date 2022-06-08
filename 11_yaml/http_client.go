package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

var client *fasthttp.Client

func LoadClient() *fasthttp.Client {
	if client != nil {
		return client
	}
	readTimeout, _ := time.ParseDuration(GlobalConfig.Http.ReadTimeout)
	writeTimeout, _ := time.ParseDuration(GlobalConfig.Http.WriteTimeout)
	maxIdleConnDuration, _ := time.ParseDuration(GlobalConfig.Http.MaxIdleConnDuration)
	client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: 10 * time.Minute,
		}).Dial,
	}
	return client
}

func Get(addr string) []byte {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(addr)
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		log.Panicln(err)
	}
	return resp.Body()
}

func Post(addr string, body []byte) []byte {
	reqTimeout := time.Duration(100) * time.Millisecond

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(addr)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetBodyRaw(body)
	resp := fasthttp.AcquireResponse()

	err := client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		log.Panicln(err)
	}
	return resp.Body()
}
