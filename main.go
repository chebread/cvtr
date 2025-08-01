// TODO: amount에서 쉼표 사용할 수 있게 하기. 쉼표는 1,000 10,00 100,0 1,0,0,0 든 아무런 상관 없음.
// -> 쉼표는 오직 사용자가 입력시 혼란을 방지하기 위한 것임. 모든 입력된 쉼표는 모두 삭제되어 정수만 남을 것임.
// TODO: usd history 추가

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var currencyList = [...]string{
	"KRW",
	"krw",
	"USD",
	"usd",
}

// KRW CPI
// https://www.index.go.kr/unify/idx-info.do?idxCd=4226
var KrwCpiData = map[int]float64{
	2024: 114.2, 2023: 111.6, 2022: 107.7, 2021: 102.5, 2020: 100.0,
	2019: 99.5, 2018: 99.1, 2017: 97.6, 2016: 95.8, 2015: 94.9,
	2014: 94.2, 2013: 93.0, 2012: 91.8, 2011: 89.9, 2010: 86.4,
	2009: 83.9, 2008: 81.7, 2007: 78.0, 2006: 76.1, 2005: 74.4,
	2004: 72.4, 2003: 69.9, 2002: 67.5, 2001: 65.7, 2000: 63.2,
	1999: 61.8, 1998: 61.3, 1997: 57.0, 1996: 54.6, 1995: 52.0,
	1994: 49.8, 1993: 46.8, 1992: 44.7, 1991: 42.1, 1990: 38.5,
	1989: 35.4, 1988: 33.5, 1987: 31.3, 1986: 30.4, 1985: 29.6,
	1984: 28.8, 1983: 28.2, 1982: 27.3, 1981: 25.4, 1980: 21.0,
	1979: 16.3, 1978: 13.8, 1977: 12.0, 1976: 10.9, 1975: 9.5,
	1974: 7.6, 1973: 6.1, 1972: 5.9, 1971: 5.3, 1970: 4.7,
	1969: 4.0, 1968: 3.6, 1967: 3.2, 1966: 2.9, 1965: 2.6,
}

// USD CPI
// https://data.bls.gov/timeseries/CUUR0000SA0
var UsdCpiData = map[int]float64{
	1965: 31.508, 1966: 32.458, 1967: 33.358, 1968: 34.783, 1969: 36.683,
	1970: 38.825, 1971: 40.492, 1972: 41.817, 1973: 44.400, 1974: 49.308,
	1975: 53.817, 1976: 56.908, 1977: 60.608, 1978: 65.233, 1979: 72.575,
	1980: 82.408, 1981: 90.925, 1982: 96.500, 1983: 99.600, 1984: 103.900,
	1985: 107.550, 1986: 109.600, 1987: 113.650, 1988: 118.250, 1989: 124.000,
	1990: 130.650, 1991: 136.200, 1992: 140.300, 1993: 144.500, 1994: 148.250,
	1995: 152.350, 1996: 156.850, 1997: 160.550, 1998: 163.000, 1999: 166.600,
	2000: 172.200, 2001: 177.050, 2002: 179.900, 2003: 183.950, 2004: 188.900,
	2005: 195.300, 2006: 201.600, 2007: 207.343, 2008: 215.303, 2009: 214.537,
	2010: 218.055, 2011: 224.939, 2012: 229.594, 2013: 232.957, 2014: 236.736,
	2015: 237.017, 2016: 240.007, 2017: 245.120, 2018: 251.107, 2019: 255.658,
	2020: 258.811, 2021: 270.969, 2022: 292.655, 2023: 304.702, 2024: 313.689,
}

// 합리적인 시점의 환율 데이터를 코드에 변수로 내장함.
var ExchangeRatesLastUpdated = map[string]time.Time{
	"USD_TO_KRW": time.Date(2025, 7, 30, 0, 0, 0, 0, time.UTC), // KRW_TO_USD = 1 / USD_TO_KRW
}
var ExchangeRates = map[string]float64{
	"USD_TO_KRW": 1384.03, // 1달러당 원화
}

func main() {
	// os.Args is []string array.
	if len(os.Args) >= 2 {
		// subcommand 존재시
		switch os.Args[1] {
		case "convert":
			// USAGE: convert <amount> <source_currency> to <target_currency>
			if len(os.Args) <= 5 {
				fmt.Println("error: 충분한 인수가 제공되지 않음.")
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
						fmt.Printf("error: %v 잘못된 플래그.\n", flag)
					}
					if len(os.Args) == 8 {
						isFlag = true
						var err error
						flagRateValue, err = strconv.ParseFloat(os.Args[7], 64) // 무조건 1달러당 원화 // 1300
						if err != nil {
							fmt.Println("error: string to float64.")
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
					fmt.Printf("error: sourceCurrency %s 통화는 지원되지 않음.\n", sourceCurrency)
					return // main 함수 종료
				}
				if !isValidTargetCurrency {
					fmt.Printf("error: targetCurrency %s 통화는 지원되지 않음.\n", targetCurrency)
					return
				}

				// amount: string -> positive int 변환
				// 참고로, amount가 0이면 의미가 없으므로 종료해야 함.
				amount, err := strconv.Atoi(strings.ReplaceAll(amountStr, ",", "")) // amount에 존재하는 쉼표를 모두 삭제함; Atoi는 int(32 or 64) type으로 변환함
				if err != nil || amount <= 0 {
					fmt.Println("error: amount는 양수여야 함.")
				}

				// main logic
				switch targetCurrency {
				case "USD", "usd":
					switch sourceCurrency {
					case "KRW", "krw":
						rate := 1 / ExchangeRates["USD_TO_KRW"]
						time := ExchangeRatesLastUpdated["USD_TO_KRW"].Format("2006-01-02")
						if isFlag {
							rate = 1 / flagRateValue
							time = "Customed"
						}

						amountInUsd := rate * float64(amount) // USD로 환산된 금액; 원화 -> 달러
						fmt.Println("USD_KRW Last Updated:", time)
						fmt.Printf("원화 -> 달러: %v$\n", amountInUsd)
					}
				case "KRW", "krw":
					switch sourceCurrency {
					case "USD", "usd":
						rate := ExchangeRates["USD_TO_KRW"]
						time := ExchangeRatesLastUpdated["USD_TO_KRW"].Format("2006-01-02")
						if isFlag {
							rate = flagRateValue
							time = "Customed"
						}

						amountInKrw := rate * float64(amount)
						fmt.Println("USD_KRW Last Updated:", time)
						fmt.Printf("달러 -> 원화: %v₩\n", amountInKrw)
					}
				}
			}
		case "history":
			// USAGE: history <amount> <currency> <start_year> to <end_year>
			if len(os.Args) <= 6 {
				fmt.Println("error: 충분한 인수가 제공되지 않음.")
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
					fmt.Printf("error: currency %s 통화는 지원되지 않음.\n", currency)
					return
				}

				amount, err := strconv.Atoi(strings.ReplaceAll(amountStr, ",", ""))
				if err != nil || amount <= 0 {
					fmt.Println("error: amount는 양수여야 함.")
					return
				}
				startYear, err := strconv.Atoi(startYearStr)
				if err != nil || startYear <= 0 {
					fmt.Println("error: startYear는 양수여야 함.")
					return
				}
				endYear, err := strconv.Atoi(endYearStr)
				if err != nil || endYear <= 0 {
					fmt.Println("error: startYear는 양수여야 함.")
					return
				}

				// sY, eY는 2024 ~ 1965년만 가능함
				isValidStartYearRange := (1965 <= startYear) && (startYear <= 2024)
				isValidEndYearRange := (1965 <= endYear) && (endYear <= 2024)
				if !isValidStartYearRange || !isValidEndYearRange {
					if !isValidStartYearRange {
						fmt.Println("error: startYear 1965년 ~ 2024년만 가능함.")
					}
					if !isValidEndYearRange {
						fmt.Println("error: endYear 1965년 ~ 2024년만 가능함.")
					}
					return
				}
				// sY와 eY가 같다면 에러를 반환함
				if startYear == endYear {
					fmt.Println("error: startYear와 endYear는 달라야 함.")
					return
				}

				// main logic
				var fromCPI float64 = KrwCpiData[startYear]
				var toCPI float64 = KrwCpiData[endYear]
				var result float64 = float64(amount) * (toCPI / fromCPI)
				var isToPresent bool = endYear > startYear

				if isToPresent {
					fmt.Printf("현재 가치: %v\n", result)
				} else {
					fmt.Printf("과거 가치: %v\n", result)
				}
			}
		case "help":
			help()
		default:
			fmt.Printf("error: %s 잘못된 명령어.\n", os.Args[1])
		}
	} else {
		help()
	}
}

// 도움말
func help() {
	fmt.Println("help")
}
