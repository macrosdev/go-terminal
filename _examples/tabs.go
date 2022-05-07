// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/extra"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	header := ui.NewPar("Press q to quit, Press h or l to switch tabs")
	header.Height = 1
	header.Width = 50
	header.Border = false
	header.TextBgColor = ui.ColorBlue

	tab1 := extra.NewTab("pierwszy")
	par2 := ui.NewPar("Press q to quit\nPress h or l to switch tabs\n")
	par2.Height = 5
	par2.Width = 37
	par2.Y = 0
	par2.BorderLabel = "Keys"
	par2.BorderFg = ui.ColorYellow
	tab1.AddBlocks(par2)

	tab2 := extra.NewTab("drugi")
	bc := ui.NewBarChart()
	data := []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.BorderLabel = "Bar Chart"
	bc.Data = data
	bc.Width = 26
	bc.Height = 10
	bc.DataLabels = bclabels
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorRed
	bc.NumColor = ui.ColorYellow
	tab2.AddBlocks(bc)

	tab3 := extra.NewTab("trzeci")
	tab4 := extra.NewTab("żółw")
	tab5 := extra.NewTab("four")
	tab6 := extra.NewTab("five")

	tabpane := extra.NewTabpane()
	tabpane.Y = 1
	tabpane.Width = 30
	tabpane.Border = true

	tabpane.SetTabs(*tab1, *tab2, *tab3, *tab4, *tab5, *tab6)

	ui.Render(header, tabpane)

	for {
		e := <-ui.PollEvent()
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.SetActiveLeft()
			ui.Clear()
			ui.Render(header, tabpane)
		case "l":
			tabpane.SetActiveRight()
			ui.Clear()
			ui.Render(header, tabpane)
		}
	}
}
