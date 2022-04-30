//   SMITH            (  //       /
//   structs           ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

//
// displayed contents
//  having a tree and an output makes sense because it just does, ok
//
type DisplayGrid struct {
	output *widgets.Paragraph
	tree *widgets.Tree
	grid *ui.Grid
	name string
}

// config.json content
type Config struct {
    Public_key string
	Secret_key string
	Mirror string
}

//
// ---- responses ----
//
type AccountStatusResponse struct {
    Data string
}

type AccountSnapshotResponseMain struct {
    Code float64
	Msg string
	SnapshotVos map[string]interface{}
}

type AccountSnapshotResponseVos struct {
	Data map[string]interface{}
	Type string
	UpdateTime float64
}

type AccountSnapshotResponseData struct {
	Balances []interface{}
	TotalAssetOfBtc float64
}

type DepositAddressResponse struct {
	Address string
	Coin string
	Tag string
	Url string
}

//
// urls to hit on mirror
//
type Targets struct {
	getall		string
	address		string
	status		string
	snapshot	string
	null		string
}

var GET_Targets = Targets {
	getall:		"/sapi/v1/capital/config/getall",
	address:	"/sapi/v1/capital/deposit/address",
	status:		"/sapi/v1/account/status",
	snapshot:	"/sapi/v1/accountSnapshot",
	null:		"/null",
}

var POST_Targets = Targets {
	getall:		"/null",
	address:	"/null",
	status:		"/null",
	snapshot:	"/null",
	null:		"/null",
}
