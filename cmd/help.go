package main

import (
	"fmt"

	"github.com/fatih/color"
)

var h = color.New(color.Bold, color.Underline).SprintFunc()
var b = color.New(color.Bold).SprintFunc()

func help() {
	// \n\n을 사용하면 fmt.Println() 호출을 줄여 코드를 더 간결하게 만들 수 있습니다.
	fmt.Printf("currency converter and historical value translator\n\n")

	fmt.Printf("%s\n", h("Usage:"))
	fmt.Printf("  cvtr convert <amount> <source_currency> to <target_currency> [-R <rate>]\n")
	fmt.Printf("  cvtr history <amount> <currency> <start_year> to <end_year>\n")
	fmt.Printf("  cvtr [-h | -V]\n\n")

	fmt.Printf("%s\n", h("Commands:"))
	// 긴 단어 뒤에는 \t 한 번, 짧은 단어 뒤에는 \t 두 번을 써서 세로선을 맞춥니다.
	fmt.Printf("  %s\t\tConvert currency rates\n", b("convert"))
	fmt.Printf("  %s\t\tTranslate historical monetary value via CPI\n\n", b("history"))

	fmt.Printf("%s\n", h("Options:"))
	fmt.Printf("  %s, %s %s\tApply a custom exchange rate\n", b("-R"), b("--rate"), b("<value>"))
	fmt.Printf("  %s, %s\t\tPrint help information\n", b("-h"), b("--help"))
	fmt.Printf("  %s, %s\t\tPrint version information\n\n", b("-V"), b("--version"))

	fmt.Printf("%s\n", h("Currencies:"))
	fmt.Printf("  %s, %s, %s\t\tKorean Won (CPI: 1965-2025)\n", b("KRW"), b("krw"), b("₩"))
	fmt.Printf("  %s, %s, %s\t\tUnited States Dollar (CPI: 1965-2025)\n", b("USD"), b("usd"), b("$"))
}
