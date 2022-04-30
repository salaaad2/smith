package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

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


func createMainWorkspace(config Config) DisplayGrid {
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
				{
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
			ui.NewCol(0.2, l),
			ui.NewCol(0.8, output),
		),
	)
	return DisplayGrid{
		output: output,
		tree: l,
		grid: main_grid,
		name: "main",
	}
}
