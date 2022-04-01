package main

import (
	"fmt"
	"log"
)

func sign_request(body string, key string) {
	fmt.Println("signing request : ", body)

	out2, err := RunStrings("/usr/bin/echo", "-n", body, "|", "/usr/bin/openssl", "dgst", "-sha256", "-hmac", key)
	if err != nil {
		log.Fatal("hwhat")
	}

	fmt.Println(out2)
	// here
}

func make_body(order string, ticker string, price string) {
	ret := "GET /sapi/v1/capital/config/getall"
	sign_request(ret, "heheheheh")
}
