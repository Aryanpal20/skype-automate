# Skype Automation using Selenium in Golang

This project is an example of how to automate Skype using Selenium in Golang.


## Stack

```
1. Using Golang
2. Using Selenium 
```

## Installation

### Prerequisite

First you need to have the following in your system:

1. Install [Golang](https://go.dev/dl/go1.20.2.linux-amd64.tar.gz) in your system.
2. Install [chromedriver](https://chromedriver.storage.googleapis.com/111.0.5563.64/chromedriver_linux64.zip) for doing automation testing by chrome using linux.

To run this project, you need to have the following Modules:

1. Install [selenium](github.com/tebeka/selenium) for using selenium automation testing in golang.
```
  go get github.com/tebeka/selenium
```
2. Install [godotenv](github.com/joho/godotenv) for fetch the data from environment file.
```
  go get github.com/joho/godotenv
```
3. Install [gocron](github.com/go-co-op/gocron) for automatically run commands on a schedule time.
```
  go get github.com/go-co-op/gocron
```

## Environment Setup 

  Create a .env file in the project root directory For example:

```
  USER_EMAIL = "<Your_Skype_Email_Or_Phone_no>"
  USER_PASSWORD = "<Ypur_Skype_Password>"
```

## Run Automation Script

  To run the main file of your app.
```
go run main.go
```