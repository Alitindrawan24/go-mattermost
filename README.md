# Go Mattermost
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go Mattermost is a very simple API for interacting with mattermost accounts

## Tech
- Go

## Package/Frameworks
- Gin

## Getting Started
To get started with Go Mattermost, you will need to have Go installed on your computer. You can download Go from the official website: https://golang.org/dl/

Once you have Go installed, follow these steps:
- Clone the repository to your local machine using
```bash
git clone https://github.com/Alitindrawan24/go-mattermost.git
```
- Navigate to the project directory using the command line.
- Copy .env.example into .env
- Fill the environment value in .env
- Run the app using
```bash
make run
```
Or using
```bash
go run cmd/main.go
```
- Open your web browser and navigate to http://localhost:8080 to access the application.

## Endpoints
- ```GET/```  get a user information
- ```PUT/custom-status```  update the custom status user

## License
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/MarketingPipeline/README-Quotes/blob/main/LICENSE) file for details.