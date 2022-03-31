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
	ctogo := C.mr_smith()
	box := tview.NewBox().SetBorder(true).SetTitle(C.GoString(ctogo))

	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
