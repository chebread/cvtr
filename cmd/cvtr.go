package main

import (
	"github.com/chebread/cvtr/internal/db"
	"github.com/chebread/cvtr/internal/lib"

	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// map 공부용

var ProgramName = os.Args[0]
var ProgramVersion = "development" // It will be replaced with a git tag.

var boldCyan = color.New(color.FgCyan, color.Bold).PrintfFunc() // cyan colored printf
var boldRed = color.New(color.FgRed, color.Bold).PrintfFunc()

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	// subcommand 존재시
	mode := os.Args[1]
	switch mode {
	case "convert":
		// USAGE:
		// cvtr convert: args[1] l2 <금액: args[2] l3> <소스_통화: args[3] l4> to: args[4] l5 <타겟_통화: args[5] l6> [--rate|-R: args[6] l7 <환율_값: args[7] l8>]

		// 6개 미만일 때 오류
		if len(os.Args) <= 5 {
			boldRed("error: Not enough arguments provided\n")
			break
		}

		var amountStr string = os.Args[2]
		var sourceCurrency string = os.Args[3]
		var keyword string = os.Args[4]
		var targetCurrency string = os.Args[5]
		var hasError bool

		// to keyword 아닐때 오류
		if keyword != "to" {
			boldRed("error: Use 'to' keyword, instead of '%s' keyword\n", keyword)
			// return
			hasError = true
		}
		// flag는 len 7에 저장되어 있다.
		// len 7일 때 정의되지 않은 다른 flag 사용시 에러를 발생 시킨다.
		// == 7 가 아니라 >= 7 인 이유는, len 8 이상으로 갈 수도 있기 때문임. 그래서 7이상으로 범위를 정한 것임.
		if len(os.Args) >= 7 {
			// "--rate" 그리고 "-R"도 아닐 때 == "rate, R" 외의 다른 flag 일 때
			if flag := os.Args[6]; flag != "--rate" && flag != "-R" {
				boldRed("error: %s is an invalid flag\n", flag)
				// return
				hasError = true
			}
		}
		//TODO: 지금은 이렇게 안하지만 rate flag가 선언되고, 꼭 value가 동봉되어야 한다. l8까지 꼭 와야 한다는 것이다. l7에서 일단 rate flag name이 맞는지 체크하고, -> 응 l7까지만 해야 한다. 왜? rate Flag만 하는게 아니잖아. 미래적으로 봤을 때는.
		// 지원 통화 체크
		// TODO: 이 코드도 좋지만, tidy에서 알려주는 것처럼, 이런 형태는 헬퍼 함수를 만들어서 활용하라.
		var isValidSourceCurrency bool = false
		var isValidTargetCurrency bool = false
		for _, i := range db.CurrencyList {
			if i == sourceCurrency {
				isValidSourceCurrency = true
			}
			if i == targetCurrency {
				isValidTargetCurrency = true
			}
		}
		if !isValidSourceCurrency {
			boldRed("error: The currency %s in <source_currency> is not supported\n", sourceCurrency)
			// return // main 함수 종료
			hasError = true
		}
		if !isValidTargetCurrency {
			boldRed("error: The currency %s in <target_currency> is not supported\n", targetCurrency)
			// return
			hasError = true
		}
		// amount: string -> positive int 변환
		// 참고로, amount가 0이면 의미가 없으므로 종료해야 함.
		amount, err := strconv.Atoi(strings.ReplaceAll(amountStr, ",", "")) // amount에 존재하는 쉼표를 모두 삭제함; Atoi는 int(32 or 64) type으로 변환함
		if err != nil || amount <= 0 {
			boldRed("error: <amount> must be positive\n")
			// return
			hasError = true
		}
		// --rate flag
		var isFlag bool
		var flagRateValue float64
		// len 8 이상이면 rate flag에 value가 있는 거다
		if len(os.Args) >= 8 {
			isFlag = true // rate flag가 있다고 확신가능하다. 왜? 위에서 이미 방어코드를 사용해 전제 조건으로 rate, R을 검사했기 때문이다.
			var err error
			flagRateValue, err = strconv.ParseFloat(os.Args[7], 64) // 무조건 1달러당 원화 // 1300
			if err != nil {
				boldRed("error: rate flag argument must be an integer\n")
				hasError = true
			}
		}
		// 오류 메시지 연속적으로 모두 출력하기 위해서.
		if hasError {
			break
		}

		// main logic
		switch targetCurrency {
		case "USD", "usd", "$":
			switch sourceCurrency {
			case "KRW", "krw", "₩":
				var rate float64
				var time string

				if isFlag {
					rate = 1 / flagRateValue
					time = "Customed"
				} else {
					rate = 1 / db.ExchangeRates["USD_KRW"] // TODO: 이게 부동소수점 error는 안나는가?
					time = db.ExchangeRatesLastUpdated["USD_KRW"].Format("2006-01-02")
				}

				amountInUsd := rate * float64(amount) // USD로 환산된 금액; 원화 -> 달러
				fmt.Printf("Last Updated: %s\n", time)
				boldCyan("%s$\n", lib.Comma(amountInUsd))
			default:
				boldRed("error: The currencies of <source_currency> and <target_currency> must be different\n")
			}
		case "KRW", "krw", "₩":
			switch sourceCurrency {
			case "USD", "usd", "$":
				var rate float64
				var time string

				if isFlag {
					rate = flagRateValue
					time = "Customed"
				} else {
					rate = db.ExchangeRates["USD_KRW"]
					time = db.ExchangeRatesLastUpdated["USD_KRW"].Format("2006-01-02")
				}

				amountInKrw := rate * float64(amount)
				fmt.Printf("Last Updated: %s\n", time)
				boldCyan("%s₩\n", lib.Comma(amountInKrw))
			default:
				boldRed("error: The currencies of <source_currency> and <target_currency> must be different\n")
			}
		}

	case "history":
		if len(os.Args) <= 6 {
			boldRed("error: Not enough arguments provided\n")
			break
		}

		var hasError bool
		var amountStr string = os.Args[2]
		var currency string = os.Args[3]
		var startYearStr string = os.Args[4]
		var keyword string = os.Args[5]
		var endYearStr string = os.Args[6]

		// 'to' keyword
		if keyword != "to" {
			boldRed("Error: Use 'to' keyword, instead of '%s' keyword\n", keyword)
			hasError = true
		}
		var isValidCurrency bool = false
		for _, i := range db.CurrencyList {
			if i == currency {
				isValidCurrency = true
			}
		}
		if !isValidCurrency {
			boldRed("error: The currency %s in <currency> is not supported\n", currency)
			hasError = true
		}
		amount, err := strconv.Atoi(strings.ReplaceAll(amountStr, ",", ""))
		if err != nil || amount <= 0 {
			boldRed("error: <amount> must be positive\n")
			hasError = true
		}
		startYear, err := strconv.Atoi(startYearStr)
		if err != nil || startYear <= 0 {
			boldRed("error: <start_year> must be positive\n")
			hasError = true
		}
		endYear, err := strconv.Atoi(endYearStr)
		if err != nil || endYear <= 0 {
			boldRed("error: <end_year> must be positive\n")
			hasError = true
		}
		// sY, eY는 2025 ~ 1965년만 가능함
		isValidStartYearRange := (1965 <= startYear) && (startYear <= 2025)
		isValidEndYearRange := (1965 <= endYear) && (endYear <= 2025)
		if !isValidStartYearRange || !isValidEndYearRange {
			if !isValidStartYearRange {
				boldRed("error: <start_year> is only valid between 1965 and 2025\n")
			}
			if !isValidEndYearRange {
				boldRed("error: <end_year> is only valid between 1965 and 2025\n")
			}
			hasError = true
		}
		// sY와 eY가 같다면 에러를 반환함
		if startYear == endYear {
			boldRed("error: <start_year> and <end_year> must be different\n")
			hasError = true
		}
		if hasError {
			break
		}

		// main logic
		var fromCPI float64
		var toCPI float64
		switch currency {
		case "KRW", "krw", "₩":
			fromCPI = db.KrwCpiData[startYear]
			toCPI = db.KrwCpiData[endYear]
		case "USD", "usd", "$":
			fromCPI = db.UsdCpiData[startYear]
			toCPI = db.UsdCpiData[endYear]
		}

		var result float64 = float64(amount) * (toCPI / fromCPI)
		var isToPresent bool = endYear > startYear

		switch isToPresent {
		case true:
			fmt.Println("Present Value")
		case false:
			fmt.Println("Past Value")
		}
		boldCyan("%s\n", lib.Comma(result))

	case "help", "-h", "--help":
		help()
	case "version", "-V", "--version":
		fmt.Printf("%s %s\n", ProgramName, ProgramVersion)
	default:
		boldRed("error: %s is an invalid command\n", os.Args[1])
	}
}
