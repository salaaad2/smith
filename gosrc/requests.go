//   SMITH            (  //       /
//   requests          ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"fmt"
	"log"
	"strings"
)

// sign request with given private key
func sign_request(body string, key string) string {
	fmt.Println("signing request : ", body)

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
func make_body(order string, ticker string, price string) string {
	ret := "GET /sapi/v1/capital/config/getall"
	return ret
}
