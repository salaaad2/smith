//   SMITH            (  //       /
//   structs           ( )/       /
//   by salade         )(/        /
//  ________________  ( /)        /
// ()__)____________)))))   :^}   /

package main

// config.json content
type Config struct {
    Public_key string
	Secret_key string
	Mirror string
}

// urls to hit on mirror
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
