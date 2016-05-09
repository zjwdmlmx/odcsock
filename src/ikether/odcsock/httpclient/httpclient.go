package httpclient

import (
	"crypto/tls"
	"net/http"
)

var Client *http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableCompression: false,
		DisableKeepAlives:  true,
	}

	Client = &http.Client{Transport: tr}
}
