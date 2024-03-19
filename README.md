# Real-Time Currency Converter

![Demo](https://private-user-images.githubusercontent.com/72413805/314117478-3c85c273-aa3d-4e2c-8831-456daa11097f.gif?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA4NjE0MjIsIm5iZiI6MTcxMDg2MTEyMiwicGF0aCI6Ii83MjQxMzgwNS8zMTQxMTc0NzgtM2M4NWMyNzMtYWEzZC00ZTJjLTg4MzEtNDU2ZGFhMTEwOTdmLmdpZj9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDAzMTklMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwMzE5VDE1MTIwMlomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTRlNThmNDg5MzZlMzhiYjZlZjgxNjFhYzk0NWY1OGZmYWE0NGIxYmNhYTA0YjI1ZTBhNWE5YWE4ZjNjNDM2ZjUmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.KXs1hUl44T9EqEi2FLjVpbrUXrKCwVCfYj05PBdUH6A)

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