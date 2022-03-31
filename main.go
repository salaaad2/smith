//   SMITH            (  //       /
//   main              ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	"github.com/rivo/tview"
)

// #include "csrc/smith.h"
// #cgo LDFLAGS: -lsmith -L.
import "C"

func main() {
	smith := tview.NewApplication()
	ctogo := C.mr_smith()

	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, false).
				AddItem(nil, 0, 1, false), width, 1, false).
			AddItem(nil, 0, 1, false)
	}

	box := tview.NewBox().SetBorder(true).SetTitle(C.GoString(ctogo))


	init_form := tview.NewForm().
		AddDropDown("You are", []string{"Salad", "Salade", "Mr Smith"}, 0, nil).
		AddInputField("First Name", "", 20, nil, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			smith.Stop()
		})

	pages := tview.NewPages().
		AddPage("box", box, true, true).
		AddPage("form", modal(init_form, 40, 20), true, true)

	if err := smith.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
