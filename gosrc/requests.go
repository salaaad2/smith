//   SMITH            (  //       /
//   requests          ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"log"
	"net/http"
	"strings"
)

// #include "../csrc/smith.h"
// #cgo LDFLAGS: -lsmith -L../
import "C"

// send request to mirror and return response
// @param body      ?qwe=asd&foo=bar
// @param signature duh
// @param endpoint  /sapi/v1/capital/config/qwe
// @param config    duhh
func send_request(body string, signature string, endpoint string, config Config) (*http.Response, error) {
	// new client to allow defering of requests
	client := &http.Client{}
	switch (endpoint) {
	case "Account Status":
		endpoint = Endpoints.status
	case "Available Coins":
		endpoint = Endpoints.getall
	case "Buy X coin ":
		return nil, nil
	}
	url := "https://" + config.Mirror + "/" + endpoint + "?" + body + "&signature=" + signature

	// add some ifs here
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("error: creating http request ", err)
	}
	req.Header.Add("X-MBX-APIKEY", config.Public_key)

	// finished bakiong request. send it
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("error: making http request ", err)
	}

	return response, nil
}

// sign request with given private key
func sign_request(body string, key string) string {

	// run pipeline
	out2, err := RunStrings("/usr/bin/echo", "-n", body, "|", "/usr/bin/openssl", "dgst", "-sha256", "-hmac", key)
	if err != nil {
		log.Fatal("error: failed to sign request", err)
	}

	// remove unwanted characters
	tok := strings.Index(out2, "(")
	last := len(out2) - 1
	first := tok + len("(stdin)= ")
	out2 = out2[first:last]
	return out2
}

// create body given choice ex :
// GET /sapi/v1/capital/config/getall ||
// POST /sapi/v1/asset/dust-btc
func make_body() string {
	ret := "timestamp=" + C.GoString(C.get_timestamp()) + "&recvWindow=50000"
	return ret
}
