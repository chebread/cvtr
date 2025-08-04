# cvtr

[[English](README.md)]

`cvtr` 은 통화 변환 및 화폐의 역사적 가치를 계산하는 CLI 도구입니다.

## 목차

  - [주요 기능](https://www.google.com/search?q=%23%EC%A3%BC%EC%9A%94-%EA%B8%B0%EB%8A%A5)
  - [사용법](https://www.google.com/search?q=%23%EC%82%AC%EC%9A%A9%EB%B2%95)
  - [지원 통화](https://www.google.com/search?q=%23%EC%A7%80%EC%9B%90-%ED%86%B5%ED%99%94)
  - [설치 방법](https://www.google.com/search?q=%23%EC%84%A4%EC%B9%98-%EB%B0%A9%EB%B2%95)
  - [라이선스](https://www.google.com/search?q=%23%EB%9D%BC%EC%9D%B4%EC%84%A0%EC%8A%A4)

## 주요 기능
  - 실시간 환율 변환
  - 역사적 화폐 가치 계산

## 사용법
### 실시간 환율 변환
```shell
cvtr convert <금액> <소스_통화> to <타겟_통화> [--rate|-R <환율 값>]
```

내장된 최근의 환율 정보를 사용하여 통화 간 금액을 즉시 변환합니다. `--rate` 또는 `-R` 플래그를 사용하여 사용자 지정 환율을 적용할 수 있습니다. 이 플래그에 사용되는 값은 1달러당 원화(KRW)를 기준으로 합니다.

```shell
# 100,000원을 달러로 변환하기
$ cvtr convert 100000 KRW to USD
Last Updated: 2025-07-30
72.25$

# 1달러당 1400원의 사용자 지정 환율로 100달러를 원화로 변환하기
$ cvtr convert 100 USD to KRW --rate 1400
Last Updated: Customed
140,000₩
```

### 역사적 화폐 가치 계산
```shell
cvtr history <금액> <통화> <시작_연도> to <종료_연도>
```

과거 특정 연도의 화폐가 현재 어느 정도의 구매력을 갖는지, 혹은 현재의 화폐가 과거에는 어느 정도의 가치였는지 계산합니다. `<시작_연도>` 또는 `<종료_연도>`에 `current`를 사용하여 현재 시점을 나타낼 수 있습니다. 사용 가능한 연도는 **1965-2024**로 제한됩니다.

```shell
# 1980년의 1000원이 2024년에는 얼마의 가치를 가질까?
$ cvtr history 1000 KRW 1980 to 2024
Present Value
5,438.095

# 2024년의 10달러가 1980년에는 얼마의 가치였을까?
$ cvtr history 10 USD 2024 to 1980
Past Value
2.627
```

### 쉼표(,)를 포함한 금액 입력
`convert` 및 `history` 명령어 모두 금액을 입력할 때 쉼표(,)를 포함하거나 제외할 수 있습니다. 예를 들어, `1,000`, `10,00`, `100,0` 등은 모두 유효한 입력입니다. 시스템은 처리 전에 모든 쉼표를 자동으로 제거하고 숫자 값만 사용합니다.

## 지원 통화
  - 대한민국 원 (KRW)
  - 미국 달러 (USD)

## 설치 방법
1.  `cvtr`의 GitHub Releases 페이지를 방문합니다.
2.  사용 중인 운영체제와 아키텍처에 맞는 파일을 다운로드합니다.
3.  다운로드한 파일의 압축을 해제합니다.
4.  `cvtr` 실행 파일을 실행합니다.
5.  더 쉬운 접근을 위해 `cvtr` 실행 파일이 있는 경로를 시스템의 `PATH` 환경 변수에 추가하는 것을 권장합니다.

## 라이선스
MIT LICENSE © 2025 Cha Haneum