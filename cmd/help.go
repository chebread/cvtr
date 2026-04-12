package main

import (
	"fmt"

	"github.com/fatih/color"
)

var h = color.New(color.Bold, color.Underline).SprintFunc()
var b = color.New(color.Bold).SprintFunc()

func help() {
	fmt.Println("A dead simple currency tool for your terminal")
	fmt.Println()

	fmt.Printf("%s\n  cvtr <COMMAND> [OPTIONS]\n\n", h("Usage:"))

	fmt.Printf("%s\n", h("Commands:"))
	fmt.Printf("  %s\n", b("convert"))
	fmt.Println("          Convert currency using internal, custom, or live exchange rates")
	fmt.Printf("  %s\n", b("history"))
	fmt.Println("          Translate currency value between years based on CPI data")
	fmt.Println()

	fmt.Printf("%s\n", h("Arguments:"))
	fmt.Printf("  %s\n", b("convert"))
	fmt.Printf("    %s\n          Amount to convert\n", b("<amount>"))
	fmt.Printf("    %s\n          Source currency code\n", b("<source>"))
	fmt.Printf("    %s\n          Target currency code\n", b("<target>"))
	fmt.Println()
	fmt.Printf("  %s\n", b("history"))
	fmt.Printf("    %s\n          Amount to translate\n", b("<amount>"))
	fmt.Printf("    %s\n          Currency code\n", b("<currency>"))
	fmt.Printf("    %s\n          The year to start from 1965-2025\n", b("<start_year>"))
	fmt.Printf("    %s\n          The year to translate to 1965-2025\n", b("<end_year>"))
	fmt.Println()

	fmt.Printf("%s\n", h("Options:"))
	fmt.Printf("  %s\n", b("convert"))
	fmt.Printf("    %s, %s %s\n", b("-R"), b("--rate"), b("<value>"))
	fmt.Println("            Apply a custom exchange rate for conversion")
	fmt.Println()
	fmt.Printf("  %s\n", b("global"))
	fmt.Printf("    %s, %s\n", b("-h"), b("--help"))
	fmt.Println("            Print help information")
	fmt.Printf("    %s, %s\n", b("-V"), b("--version"))
	fmt.Println("            Print version information")
}
