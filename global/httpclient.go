//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package global

import (
	"crypto/tls"
	"net/http"
)

// Client the global http client object
var Client *http.Client

func initHttpClient() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // http client with skip verify of CA
		},
		DisableCompression: false,
		DisableKeepAlives:  true,
	}

	Client = &http.Client{Transport: tr}
}
