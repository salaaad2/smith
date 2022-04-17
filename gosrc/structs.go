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
	status		string
	null		string
}

var Endpoints = Targets {
	getall:		"/sapi/v1/capital/config/getall",
	status:		"/sapi/v1/account/status",
	null:		"/null",
}

