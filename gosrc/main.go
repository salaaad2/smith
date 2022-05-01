//   SMITH            (  //       /
//   main              ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
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

	ui_loop(config)
	defer ui.Close()
	return
}

// draw stuff and get input
func ui_loop(config Config) error {
	if err := ui.Init(); err != nil {
		log.Fatal("error: failed to initialize termui", err)
	}

	// create default grid
	active_grid := createMainWorkspace(config)
	// ui update loop
	ui.Render(active_grid.grid)
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
					active_grid.grid.SetRect(0, 0, payload.Width, payload.Height)
					ui.Clear()
				case "j", "<Down>":
					active_grid.tree.ScrollDown()
				case "k", "<Up>":
					active_grid.tree.ScrollUp()
				case "e":
					active_grid.tree.ToggleExpand()
				case "<Home>":
					active_grid.tree.ScrollTop()
				case "G", "<End>":
					active_grid.tree.ScrollBottom()
				case "E":
					active_grid.tree.ExpandAll()
				case "C":
					active_grid.tree.CollapseAll()
				case "<Enter>":
					active_grid.handleEnter(config)
				case "2":
					active_grid = createTradeWorkspace(config)
					ui.Clear()
					ui.Render(active_grid.grid)
					ui_events = ui.PollEvents()
			}
			case <-ticker:
				ticker_count++
			default:
				ui.Render(active_grid.grid)
		}
	}
}
