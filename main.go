package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// os.Args is []string array.
	if len(os.Args) >= 2 {
		// subcommand 존재시
		switch os.Args[1] {
		case "convert":
			if len(os.Args) <= 5 {
				boldRed("error: Not enough arguments provided\n")
			} else {
				amountStr := os.Args[2]
				sourceCurrency := os.Args[3]
				targetCurrency := os.Args[5]

				isFlag := false
				var flagRateValue float64

				// --rate flag
				if len(os.Args) >= 7 { // len은 1부터이다.
					// rate flag에 값이 제공되지 않으면 그냥 아무 활동도 하지 않음.
					flag := os.Args[6]
					if flag != "--rate" && flag != "-R" { // --rate"가 아니고, 그리고 "-R"도 아닐 때
						boldRed("error: %s is an invalid flag\n", flag)
						return
					}
					if len(os.Args) == 8 {
						isFlag = true
						var err error
						flagRateValue, err = strconv.ParseFloat(os.Args[7], 64) // 무조건 1달러당 원화 // 1300
						if err != nil {
							boldRed("error: Error converting string to float64\n")
							return
						}

					}
				}

				// 지원 통화 체크
				var isValidSourceCurrency bool = false
				var isValidTargetCurrency bool = false
				for _, i := range currencyList {
					if i == sourceCurrency {
						isValidSourceCurrency = true
					}
					if i == targetCurrency {
						isValidTargetCurrency = true
					}
				}
				if !isValidSourceCurrency {
					boldRed("error: The currency %s in <source_currency> is not supported\n", sourceCurrency)
					return // main 함수 종료
				}
				if !isValidTargetCurrency {
					boldRed("error: The currency %s in <target_currency> is not supported\n", targetCurrency)
					return
				}

				// amount: string -> positive int 변환
				// 참고로, amount가 0이면 의미가 없으므로 종료해야 함.
				amount, err := strconv.Atoi(strings.ReplaceAll(amountStr, ",", "")) // amount에 존재하는 쉼표를 모두 삭제함; Atoi는 int(32 or 64) type으로 변환함
				if err != nil || amount <= 0 {
					boldRed("error: <amount> must be positive\n")
					return
				}

				// main logic
				switch targetCurrency {
				case "USD", "usd":
					switch sourceCurrency {
					case "KRW", "krw":
						rate := 1 / ExchangeRates["USD_KRW"]
						time := ExchangeRatesLastUpdated["USD_KRW"].Format("2006-01-02")
						if isFlag {
							rate = 1 / flagRateValue
							time = "Customed"
						}

						amountInUsd := rate * float64(amount) // USD로 환산된 금액; 원화 -> 달러
						fmt.Printf("Last Updated: %s\n", time)
						boldCyan("%s$\n", comma(amountInUsd))
					}
				case "KRW", "krw":
					switch sourceCurrency {
					case "USD", "usd":
						rate := ExchangeRates["USD_KRW"]
						time := ExchangeRatesLastUpdated["USD_KRW"].Format("2006-01-02")
						if isFlag {
							rate = flagRateValue
							time = "Customed"
						}

						amountInKrw := rate * float64(amount)
						fmt.Printf("Last Updated: %s\n", time)
						boldCyan("%s₩\n", comma(amountInKrw))
					}
				}
			}
		case "history":
			if len(os.Args) <= 6 {
				boldRed("error: Not enough arguments provided\n")
			} else {
				amountStr := os.Args[2]
				currency := os.Args[3]
				startYearStr := os.Args[4]
				endYearStr := os.Args[6]

				var isValidCurrency bool = false
				for _, i := range currencyList {
					if i == currency {
						isValidCurrency = true
					}
				}
				if !isValidCurrency {
					boldRed("error: The currency %s in <currency> is not supported\n", currency)
					return
				}

				amount, err := strconv.Atoi(strings.ReplaceAll(amountStr, ",", ""))
				if err != nil || amount <= 0 {
					boldRed("error: <amount> must be positive\n")
					return
				}
				startYear, err := strconv.Atoi(startYearStr)
				if err != nil || startYear <= 0 {
					boldRed("error: <start_year> must be positive\n")
					return
				}
				endYear, err := strconv.Atoi(endYearStr)
				if err != nil || endYear <= 0 {
					boldRed("error: <end_year> must be positive\n")
					return
				}

				// sY, eY는 2024 ~ 1965년만 가능함
				isValidStartYearRange := (1965 <= startYear) && (startYear <= 2024)
				isValidEndYearRange := (1965 <= endYear) && (endYear <= 2024)
				if !isValidStartYearRange || !isValidEndYearRange {
					if !isValidStartYearRange {
						boldRed("error: <start_year> is only valid between 1965 and 2024\n")
					}
					if !isValidEndYearRange {
						boldRed("error: <end_year> is only valid between 1965 and 2024\n")
					}
					return
				}
				// sY와 eY가 같다면 에러를 반환함
				if startYear == endYear {
					boldRed("error: <start_year> and <end_year> must be different\n")
					return
				}

				// main logic
				var fromCPI float64
				var toCPI float64
				switch currency {
				case "KRW", "krw":
					fromCPI = KrwCpiData[startYear]
					toCPI = KrwCpiData[endYear]
				case "USD", "usd":
					fromCPI = UsdCpiData[startYear]
					toCPI = UsdCpiData[endYear]
				}

				var result float64 = float64(amount) * (toCPI / fromCPI)
				var isToPresent bool = endYear > startYear

				if isToPresent {
					fmt.Println("Present Value")
				} else {
					fmt.Println("Past Value")
				}
				boldCyan("%s\n", comma(result))

			}
		case "help":
			help()
		default:
			boldRed("error: %s is an invalid command\n", os.Args[1])
		}
	} else {
		help()
	}
}
