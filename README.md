# cvtr
`cvtr` is a currency converter and historical value translator.

## Table of Contents
- [Features](#features)
- [How to use](#how-to-use)
- [Currently supported currencies](#currently-supported-currencies)
- [Installation](#installation)
- [License](#license)

## Features
- Real-time currency conversion
- Historical value calculation

## How to use
### Real-time currency conversion
```shell
cvtr convert <amount> <source_currency> to <target_currency>
```
Instantly convert any amount between currencies using the latest exchange rates.

```shell
# To convert 100,000 KRW to USD:
cvtr convert 100000 KRW to USD
```

### Historical value calculation
```shell
cvtr history <amount> <currency> <start_year> to <end_year>
```
Ever wondered what "100 million KRW in the 1980s" would be worth today, or how much "1,000 KRW in 2025" would have been worth in the 2000s? This command estimates the purchasing power of a monetary value from a specified year1 in a target year2. You can use current for either <start_year> or <end_year> to represent the present day. Please note that the start and end years are limited to the range of **1965-2024**.

```shell
# How much will 1,000 KRW from 1980 be worth in 2024?
cvtr history 1000 KRW 1980 to 2024

# How much is 1,000 KRW worth in 2024 in 1980?
cvtr history 1000 KRW 2024 to 1980
```

## Currently supported currencies
- Korean Won (KRW)
- US Dollar (USD)

## Installation
1. Visit the GitHub Releases page for `cvtr`.
2. Download the appropriate file for your operating system and architecture.
3. Unachive the downloaded file.
4. Execute the `cvtr` executable file.
5. For easier access, consider adding `cvtr` executable file to your system's PATH environment variable.

## License
MIT LICENSE &copy; 2025 Cha Haneum