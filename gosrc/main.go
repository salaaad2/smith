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
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
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

	if err != nil {
		log.Fatal("error: request failed %s", "qwe", err)
		return
	}
	// make a cool ui
	ui_loop(config)

	fmt.Println("qwe")
	defer ui.Close()
	return
}

// draw stuff
func ui_loop(config Config) error {
	if err := ui.Init(); err != nil {
		log.Fatal("error: failed to initialize termui", err)
	}

	// set widgets
	p1 := widgets.NewParagraph()
	p2 := widgets.NewParagraph()
	p3 := widgets.NewParagraph()
	p4 := widgets.NewParagraph()
	p5 := widgets.NewParagraph()

	p1.Text = config.Mirror
	p2.Text = "1000.0 BTC"
	p3.Text = "hello"
	p4.Text = "cool"
	p5.Text = "__MR_SMITH_V001__"
	p5.TextStyle.Fg = ui.ColorGreen

	p1.Border = true
	p2.Border = true
	p3.Border = true
	p4.Border = true

	p1.Title = "Active Mirror"
	p2.Title = "Balance"
	p3.Title = "public"
	p4.Title = "Key"
	p5.Title = "hello"

	main_grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	main_grid.SetRect(0, 0, termWidth, termHeight)
	main_grid.Border = true

	// add items to grid
	main_grid.Set(
		ui.NewRow(0.1, p5),
		ui.NewRow(1.0/8, p2),
		ui.NewRow(1.0/2, p1),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, p3),
			ui.NewCol(1.0/2, p4),
		),
	)

	// ui update loop
	ui.Render(main_grid)
	ui_events := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	ticker_count := 0
	for {
		select {
		case e := <-ui_events:
			switch e.ID {
				case "<C-c>":
					return nil
				case "<Resize>":
					payload := e.Payload.(ui.Resize)
					main_grid.SetRect(0, 0, payload.Width, payload.Height)
					ui.Clear()
					ui.Render(main_grid)
				default:
					p3.Text = e.ID
			}
		case <-ticker:
			ui.Render(main_grid)
			ticker_count++
		}
	}
}

// make request body and sign it
// body := make_body()
// signature := sign_request(body, config.Secret_key)
// res, err := send_request(body, signature, Endpoints.getall, config)
