# cvtr

[![GitHub Release](https://img.shields.io/github/v/release/chebread/cvtr?style=flat-square)](https://github.com/chebread/cvtr/releases)
[![GoReleaser](https://github.com/chebread/cvtr/actions/workflows/release.yml/badge.svg)](https://github.com/chebread/cvtr/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/chebread/cvtr)](https://goreportcard.com/report/github.com/chebread/cvtr)

`cvtr` is a currency converter and historical value translator.

## Table of Contents
- [Features](#features)
- [How to use](#how-to-use)
- [Example](#example)
- [Currently supported currencies](#currently-supported-currencies)
- [Installation](#installation)
- [Updating](#updating)
- [License](#license)

## Features
- Currency conversion
- Historical value calculation

## How to use
### Currency conversion
```shell
cvtr convert <amount> <source_currency> to <target_currency> [--rate|-R <rate_value>]
```
Instantly convert any amount between currencies. By default, the command uses a set of built-in, recent exchange rates. To use a custom rate, you can provide the **--rate or -R flag**. The value for this flag specifies the rate in KRW per 1 USD.

```shell
# To convert 100,000 KRW to USD
> cvtr convert 100000 KRW to USD
Last Updated: 2025-07-30
72.25276908737528$

# To convert 100 USD to KRW with a custom rate of 1400 KRW per 1 USD
> cvtr convert 100 $ to KRW --rate 1400
Last Updated: Customed
140,000₩
```

### Historical value calculation
```shell
cvtr history <amount> <currency> <start_year> to <end_year>
```
Ever wondered what "100 million KRW in the 1980s" would be worth today, or how much "1,000 KRW in 2025" would have been worth in the 2000s? This command estimates the purchasing power of a monetary value from a specified start_year in a target end_year. You can use current for either <start_year> or <end_year> to represent the present day. Please note that the start and end years are limited to the range of **1965-2024**.

```shell
# How much will 1,000 KRW from 1980 be worth in 2024?
> cvtr history 1000 ₩ 1980 to 2024
Present Value
5,438.0952380952385

# How much is 10 USD worth in 2024 in 1980?
> cvtr history 10 USD 2024 to 1980
Past Value
2.6270605599813828
```

### Flexible Amount Input
You can input the <amount> with or without commas in both cvtr convert and cvtr history commands. For example, `1,000`, `10,00`, `100,0`, or even `1,0,0,0` are all valid. The system will automatically remove all commas before processing, using only the numerical value.

### Accumulated Error Display
Instead of exiting immediately on the first error, `cvtr` aggregates all validation issues and displays them at once. This allows users to intuitively identify and fix all input errors in a single run.

```bash
$ cvtr history 10x,00 usda to krwb --rate 30x00
error: Use 'to' keyword, instead of 'krwb' keyword
error: The currency usda in <currency> is not supported
error: <amount> must be positive
error: <start_year> must be positive
error: <end_year> must be positive
error: <start_year> is only valid between 1965 and 2025
error: <end_year> is only valid between 1965 and 2025
error: <start_year> and <end_year> must be different

For more information, try '--help'.
```

## Example
### Live Rates
Get real-time exchange rates using `curl`, `jq` and [exchangerate-api](https://www.exchangerate-api.com/) with the `--rate` flag.

```shell
cvtr convert 1 USD to KRW --rate $(curl -s "https://v6.exchangerate-api.com/v6/YOUR_API_KEY/pair/USD/KRW" | jq -r '.conversion_rate')
```

## Currently supported currencies
- Korean Won: KRW, krw, ₩
- US Dollar: USD, usd, $

## Installation
### On macOS
You can install `cvtr` with Homebrew:
```shell
brew tap chebread/cvtr

brew install cvtr
```

## Updating
### On macOS
If you installed `cvtr` using Homebrew, you can easily upgrade to the latest version when it's released.

```shell
brew upgrade cvtr
```

## License
MIT LICENSE &copy; 2025-2026 Cha Haneum
