//   SMITH            (  //       /
//   main              ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

	// get user config from json file
	config_path := "./config.json"
	for i, a := range os.Args[1:] {
		if a == "-c" || a == "--config" {
			config_path = os.Args[i + 2]
		}
	}
	config_content, err := ioutil.ReadFile(config_path)
	if err != nil {
		log.Fatal("error: config file not found", err)
	}
	var config Config
	err = json.Unmarshal(config_content, &config)
	if err != nil {
		log.Fatal("error: marshall() ", err)
	}

	make_ui()

	fmt.Println("\n" +
		Styles.colorBlue + "public_key: " + Styles.colorReset + config.Public_key + "\n" +
		Styles.colorBlue + "secret_key: " + Styles.colorReset + config.Secret_key + "\n" +
		Styles.colorBlue + "mirror:     " + Styles.colorReset + config.Mirror)

	// make request body and sign it
	body := make_body("qwe", "qwe", "qwe")
	fmt.Println(body)
	signature := sign_request(body, config.Secret_key)
	fmt.Println("[" + signature + "]")
}

func make_ui() error {
	if err := ui.Init(); err != nil {
		log.Fatal("error: failed to initialize termui", err)
	}

	p:= widgets.NewParagraph()
	p.Text = "hello"
	p.SetRect(0, 0, 25, 5)
	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			return nil
		}
	}

	return nil
}
