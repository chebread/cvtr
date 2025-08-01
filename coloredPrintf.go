package main

import "github.com/fatih/color"

var boldCyan = color.New(color.FgCyan, color.Bold).PrintfFunc() // cyan colored printf
var boldRed = color.New(color.FgRed, color.Bold).PrintfFunc()
