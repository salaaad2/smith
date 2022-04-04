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
	null		string
}

var Endpoints = Targets {
	getall:		"/sapi/v1/capital/config/getall",
	null:		"/null",
}

