# cvtr

[[한국어](README.kr.md)]

`cvtr` is a currency converter and historical value translator.

## Table of Contents
- [Features](#features)
- [How to use](#how-to-use)
- [Currently supported currencies](#currently-supported-currencies)
- [Installation](#installation)
- [Updating](#updating)
- [License](#license)

## Features
- Real-time currency conversion
- Historical value calculation

## How to use
### Real-time currency conversion
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
> cvtr convert 100 USD to KRW --rate 1400
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
> cvtr history 1000 KRW 1980 to 2024
Present Value
5,438.0952380952385

# How much is 10 USD worth in 2024 in 1980?
> cvtr history 10 USD 2024 to 1980
Past Value
2.6270605599813828
```

### Flexible Amount Input
You can input the <amount> with or without commas in both cvtr convert and cvtr history commands. For example, `1,000`, `10,00`, `100,0`, or even `1,0,0,0` are all valid. The system will automatically remove all commas before processing, using only the numerical value.

## Currently supported currencies
- Korean Won (KRW)
- US Dollar (USD)

## Installation
### On macOS
You can install `cvtr` with Homebrew:
```shell
brew tap chebread/cvtr

brew install cvtr
```

### For other OS
1. Visit [the GitHub Releases page](https://github.com/chebread/cvtr/releases) for `cvtr`.
2. Download the appropriate file for your operating system and architecture.
3. Unachive the downloaded file.
4. Execute the `cvtr` executable file.
5. For easier access, consider adding `cvtr` executable file to your system's PATH environment variable.

## Updating
### On macOS
If you installed `cvtr` using Homebrew, you can easily upgrade to the latest version when it's released.

```shell
brew upgrade cvtr
```

### For Other OS
For other OS, you will need to download the new version from [the GitHub Releases page](https://github.com/chebread/cvtr/releases) for `cvtr`.

Download the latest release for your system and replace your old executable file with the new one.

## License
MIT LICENSE &copy; 2025 Cha Haneum
