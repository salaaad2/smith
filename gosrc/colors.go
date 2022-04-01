package main

// Theme defines the colors used when primitives are initialized.
type Theme struct {
 colorReset   	string // Main background color for primitives.
 colorRed     	string // Background color for contrasting elements.
 colorGreen   	string // Background color for even more contrasting elements.
 colorYellow  	string // Box borders.
 colorBlue    	string // Box titles.
 colorPurple  	string // Graphics.
 colorCyan    	string // Primary text.
 colorWhite   	string // Secondary text (e.g. labels).
}

// Styles defines the theme for applications. The default is for a black
// background and some basic colors: black, white, yellow, green, cyan, and
// blue.
var Styles = Theme{
 colorReset:   	"\033[0m",
 colorRed :    	"\033[31m",
 colorGreen :  	"\033[32m",
 colorYellow : 	"\033[33m",
 colorBlue :   	"\033[34m",
 colorPurple : 	"\033[35m",
 colorCyan :   	"\033[36m",
 colorWhite :  	"\033[37m",
}
