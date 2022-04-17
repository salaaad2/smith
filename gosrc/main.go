//   SMITH            (  //       /
//   main              ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

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
		log.Fatalf("error: request failed %s", err)
		return
	}

	fmt.Println("hello")
	ui_loop(config)
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
	requestTree := []*widgets.TreeNode{
		{
			Value: nodeValue("Information"),
			Nodes:[]*widgets.TreeNode{
				{
					Value: nodeValue("Account Status"),
					Nodes:nil,
				},
				{
					Value: nodeValue("Available Coins"),
					Nodes: nil,
				},
			},
		},
		{
			Value: nodeValue("Actions"),
			Nodes:[]*widgets.TreeNode{
				{
					Value: nodeValue("Buy x coin"),
					Nodes:nil,
				},
			},
		},
	}
	l := widgets.NewTree()
	l.WrapText = false
	l.SetNodes(requestTree)
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	p2 := widgets.NewParagraph()
	output := widgets.NewParagraph()
	p5 := widgets.NewParagraph()

	p1.Text = config.Mirror
	p2.Text = "1000.0 BTC"
	output.Text = "hello"
	p5.Text = "__MR_SMITH_V001__"
	p5.TextStyle.Fg = ui.ColorGreen

	p1.Border = true
	p2.Border = true
	output.Border = false

	p1.Title = "Active Mirror"
	p2.Title = "Balance"
	output.Title = "Output"
	p5.Title = "hello"

	main_grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	main_grid.SetRect(0, 0, termWidth, termHeight)
	main_grid.Border = true

	// add items to grid
	main_grid.Set(
		ui.NewRow(0.1, p5),
		ui.NewRow(1.0/8, p2),
		ui.NewRow(1.0/2, l),
		ui.NewRow(1.0/2, output),
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
				case "j", "<Down>":
					l.ScrollDown()
				case "k", "<Up>":
					l.ScrollUp()
				case "e":
					l.ToggleExpand()
				case "<Home>":
					l.ScrollTop()
				case "G", "<End>":
					l.ScrollBottom()
				case "E":
					l.ExpandAll()
				case "C":
					l.CollapseAll()
			case "<Enter>":
				if l.SelectedNode().Nodes == nil {
					body := make_body()
					signature := sign_request(body, config.Secret_key)
					rep, err := send_request(body, signature, l.SelectedNode().Value.String(), config)
					if rep == nil || err != nil {
						output.Text = "something went wrong :^{"
					} else {
						body, _ := io.ReadAll(rep.Body)
						output.Text = string(body)
						defer rep.Body.Close()
					}
				}
			}
			case <-ticker:
				ticker_count++
			default:
				ui.Render(main_grid)
		}
	}
}

