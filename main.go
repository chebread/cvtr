package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	currencyList := [...]string{
		"KRW",
		"USD",
	}
	const UsdToKrw = 1384.03  // 1 달러 = 1388.50 원
	const KrwToUsd = 0.00072  // 1 원 = 0.00072 달러
	_, _ = UsdToKrw, KrwToUsd // handle unused error; 튜플 할당

	// os.Args is []string array.
	if len(os.Args) >= 2 {
		// subcommand 존재시
		switch os.Args[1] {
		case "convert":
			// USAGE: convert <amount> <source_currency> to <target_currency>
			if len(os.Args) <= 5 {
				fmt.Println("error: 충분한 인수가 제공되지 않음.")
			} else {
				// TDOO: type이 안맞으면 오류 내보내기. => os.Args는 모두 string으로 취급된다.
				// => 그러면, string에서 타입 변환시 에러 체크하기.
				amountStr := os.Args[2]
				sourceCurrency := os.Args[3]
				targetCurrency := os.Args[5]

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
					fmt.Printf("error: sourceCurrency %s 통화는 지원되지 않음.\n", sourceCurrency)
					return // main 함수 종료
				}
				if !isValidTargetCurrency {
					fmt.Printf("error: targetCurrency %s 통화는 지원되지 않음.\n", targetCurrency)
					return
				}

				// amount: string -> positive int 변환
				// 참고로, amount가 0이면 의미가 없으므로 종료해야 함.
				amount, err := strconv.Atoi(amountStr)
				if err != nil || amount <= 0 {
					fmt.Println("error: amount는 양수여야 함.")
				}

				// main logic
				switch targetCurrency {
				case "USD":
					switch sourceCurrency {
					case "KRW":
						// convert 3000 KRW to USD
						amountInUsd := KrwToUsd * float64(amount) // USD로 환산된 금액; 원화 -> 달러
						fmt.Printf("원화 -> 달러: %v$\n", amountInUsd)
					}
				case "KRW":
					switch sourceCurrency {
					case "USD":
						amountInKrw := UsdToKrw * float64(amount)
						fmt.Printf("달러 -> 원화: %v₩\n", amountInKrw)
					}
				}
			}
		case "history":
			// USAGE: history <amount> <currency> <start_year> to <end_year>
			fmt.Println("history")
			if len(os.Args) <= 6 {
				fmt.Println("error: 충분한 인수가 제공되지 않음.")
			} else {
				amountStr := os.Args[2]
				currency := os.Args[3]
				startYearStr := os.Args[4]
				endYearStr := os.Args[6]
				fmt.Println(amountStr, currency, startYearStr, endYearStr)

				var isValidCurrency bool = false
				for _, i := range currencyList {
					if i == currency {
						isValidCurrency = true
					}
				}
				if !isValidCurrency {
					fmt.Printf("error: currency %s 통화는 지원되지 않음.\n", currency)
					return
				}

				amount, err := strconv.Atoi(amountStr)
				if err != nil || amount <= 0 {
					fmt.Println("error: amount는 양수여야 함.")
				}
				startYear, err := strconv.Atoi(startYearStr)
				if err != nil || startYear <= 0 {
					fmt.Println("error: startYear는 양수여야 함.")
				}
				endYear, err := strconv.Atoi(endYearStr)
				if err != nil || endYear <= 0 {
					fmt.Println("error: startYear는 양수여야 함.")
				}

				// main logic
			}
		default:
			fmt.Printf("error: %s 잘못된 명령어.\n", os.Args[1])
		}
	}
}
