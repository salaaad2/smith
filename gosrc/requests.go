//   SMITH            (  //       /
//   requests          ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"strconv"
)

// #include "../csrc/smith.h"
// #cgo LDFLAGS: -lsmith -L../
import "C"

func sendRequest(node_name string, config Config) (*http.Response, error) {
	endpoint := getRequestType(node_name)
	if len(endpoint) == 0 {
		return nil, nil
	}

	// create payload
	body := makeBody(endpoint)
	signature := signRequest(body, config.Secret_key)
	url := "https://" + config.Mirror + "/" + endpoint + "?" + body + "&signature=" + signature
	// make a request out of it
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-MBX-APIKEY", config.Public_key)
	if err != nil {
		log.Fatal("error: creating http request ", err)
	}

	// send it
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("error: making http request ", err)
		return response, err
	}
	return response, nil
}

// return request path from tree node name
func getRequestType(name string) string {
	switch (name) {
		case "Account Status":
			return GET_Targets.status
		case "Available Coins":
			return GET_Targets.getall
		case "Deposit Address":
			return GET_Targets.address
		case "Daily Snapshot":
			return GET_Targets.snapshot
		default :
			return ""
	}
}

// display something cool
func displayResponse(node_name string, response *http.Response, output *string) {
	requestType := getRequestType(node_name)
	body, _ := io.ReadAll(response.Body)
	var outputFormatted string

	switch (requestType) {
		case GET_Targets.status:
			var structuredRep AccountStatusResponse
			json.Unmarshal(body, &structuredRep)
			outputFormatted =
				"Status     [" + structuredRep.Data + "]\n"
		case GET_Targets.address:
			var structuredRep DepositAddressResponse
			json.Unmarshal(body, &structuredRep)
			outputFormatted =
				"Coin     [" + structuredRep.Coin + "]\n" +
				"Address  [" + structuredRep.Address + "]\n" +
				"Url      [" + structuredRep.Url + "]\n" +
				"Tag      [" + structuredRep.Tag + "]\n"
		case GET_Targets.snapshot:
			var structuredRep AccountSnapshotResponseMain
			json.Unmarshal(body, &structuredRep)
			outputFormatted = "code  [" + strconv.FormatFloat(structuredRep.Code, 'f', 2, 64) + "]\n"
		default:
			outputFormatted = string(body)
	}
	*output = outputFormatted
}

// Sign payload using openssl
func signRequest(body string, key string) string {
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

// return body value from endpoint
func makeBody(endpoint string) string {
	ret := "timestamp=" + C.GoString(C.get_timestamp()) + "&recvWindow=50000"

	switch (endpoint) {
		case GET_Targets.getall , GET_Targets.status:
			break
		case GET_Targets.snapshot:
			ret += "&type=SPOT"
		case GET_Targets.address:
			ret += "&coin=BNB"
	}
	return ret
}
