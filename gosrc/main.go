//   SMITH            (  //       /
//   main              ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// #include "../csrc/smith.h"
// #cgo LDFLAGS: -lsmith -L../
import "C"

type Config struct {
    Public_key string
	Secret_key string
	Mirror string
}

func main() {
	fmt.Println("Welcome, traveller, my name is " + Styles.colorGreen + C.GoString(C.mr_smith()) + Styles.colorReset +
	"\nI will try to help you get the bag...")

	config_path := "./config.json"
	config_content, err := ioutil.ReadFile(config_path)
	if err != nil {
		log.Fatal("config file not found")
	}

	var config Config
	err = json.Unmarshal(config_content, &config)

	if err != nil {
		log.Fatal("error during marshall() ", err)
	}

	fmt.Println("\n" +
		Styles.colorBlue + "public_key: " + Styles.colorReset + config.Public_key + "\n" +
		Styles.colorBlue + "secret_key: " + Styles.colorReset + config.Secret_key + "\n" +
		Styles.colorBlue + "mirror:     " + Styles.colorReset + config.Mirror)
}
