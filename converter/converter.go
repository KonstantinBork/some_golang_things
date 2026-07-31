package main

import (
	"api"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Not the right amount of arguments, please specify the converter and a value!")
		os.Exit(1)
	}

	converter := args[0]
	value, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Error parsing the value, please provide a number!")
		os.Exit(1)
	}

	switch converter {
	case "celsius-to-fahrenheit":
		fmt.Println(celsiusToFahrenheit(value))
	case "fahrenheit-to-celsius":
		fmt.Println(fahrenheitToCelsius(value))
	}
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}

func dollarToEuro(dollar float64) float64 {
	fromCurrency := "USD"
	toCurrency := "EUR"

	requestUrl := buildCurrencyConverterRequestUrl(dollar, fromCurrency, toCurrency)
	response := api.Get(requestUrl)

	return 0
}

func euroToDollar(euro float64) float64 {
	fromCurrency := "EUR"
	toCurrency := "USD"

	requestUrl := buildCurrencyConverterRequestUrl(euro, fromCurrency, toCurrency)

	return 0
}

func buildCurrencyConverterRequestUrl(value float64, fromCurrency, toCurrency string) string {
	baseUrl := "https://openexchangerates.org/api/convert/"
	apiKey := "31aafbeb2cc143a28c6571d9e91cbbdb"

	return fmt.Sprintf("%s/%s/%s/%s?app_id=%s", baseUrl, value, fromCurrency, toCurrency, apiKey)
}
