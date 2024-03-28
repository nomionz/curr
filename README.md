# Real-Time Currency Converter

![Demo](https://github.com/nomionz/currency-converter/assets/72413805/d72117c5-3208-466e-a64a-a00fd1aeb5bd)

## Description

This is a real-time currency converter that uses the [European Central Bank exchange rates data](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html) to convert between different currencies. The user can input the amount of money they want to convert and select the currency they want to convert from and to. The app will then display the converted amount in real-time.

## How to run

1. Clone the repository
2. Run the following command to start the server
```bash
go run .
```
3. Open your browser and go to `http://localhost`

## Technologies

- [HTMX](https://htmx.org/) latest CDN
- [Tailwind CSS](https://tailwindcss.com/) latest CDN
- [Golang](https://golang.org/) 1.22
- [GoFiber](https://gofiber.io/) v2