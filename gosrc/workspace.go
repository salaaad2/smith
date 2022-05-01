package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

//
//
func (d DisplayGrid) handleEnter(config Config) {
	switch d.name {
	case "main":
		if d.tree.SelectedNode().Nodes == nil {
			rep, err := sendRequest(d.tree.SelectedNode().Value.String(), config)
			if rep == nil || err != nil {
				d.output.Text = "something went wrong :^{"
			} else {
				displayResponse(d.tree.SelectedNode().Value.String(), rep, &d.output.Text)
				defer rep.Body.Close()
			}
		}
	}
}

// create workspace one, where general information is displayed and fetched
func createMainWorkspace(config Config) *DisplayGrid {
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
				}, {
					Value: nodeValue("Deposit Address"),
					Nodes: nil,
				},
				{
					Value: nodeValue("Daily Snapshot"),
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
	tree := widgets.NewTree()
	tree.WrapText = false
	tree.SetNodes(requestTree)
	tree.TextStyle = ui.NewStyle(ui.ColorYellow)

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
	output.Border = true
	output.TextStyle.Fg = ui.ColorGreen

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
		ui.NewRow(0.1, p2),
		ui.NewRow(0.8,
			ui.NewCol(0.2, tree),
			ui.NewCol(0.8, output),
		),
	)

	grid := DisplayGrid{
		output: output,
		tree: tree,
		grid: main_grid,
		name: "main",
	}
	return &grid
}

// create workspace two, where the user can trade symbols
func createTradeWorkspace(config Config) *DisplayGrid {
	output    := widgets.NewParagraph()
	ticker    := widgets.NewParagraph()
	strategy  := widgets.NewParagraph()
	tree      := widgets.NewTree()
	main_grid := ui.NewGrid()

	ticker.Title   = "symbol"
	ticker.Text   = "<pick a symbol>"
	strategy.Title   = "strategy"
	strategy.Text = "<then, pick a strategy>"

	requestTree := []*widgets.TreeNode{
		{
			Value: nodeValue("Available Symbols"),
			Nodes:[]*widgets.TreeNode{
				{
					Value: nodeValue("BTCETH"),
					Nodes:nil,
				},
			},
		},
		{
			Value: nodeValue("Available Strategies"),
			Nodes:[]*widgets.TreeNode{
				{
					Value: nodeValue("FizzBozzNacci(TM)"),
					Nodes:nil,
				},
			},
		},
	}
	tree.WrapText = false
	tree.SetNodes(requestTree)
	tree.TextStyle = ui.NewStyle(ui.ColorYellow)

	termWidth, termHeight := ui.TerminalDimensions()
	main_grid.SetRect(0, 0, termWidth, termHeight)
	main_grid.Border = true
	main_grid.Set(
		ui.NewRow(0.1, ticker),
		ui.NewRow(0.1, strategy),
		ui.NewRow(0.8,
			ui.NewCol(0.2, tree),
			ui.NewCol(0.8, output),
		),
	)
	grid := DisplayGrid{
		output: output,
		tree: tree,
		grid: main_grid,
		name: "trading",
	}
	return &grid
}
