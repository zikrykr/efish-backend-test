# e-Fishery Backend Engineer Test #
Backend engineer test project

## Apps

Auth App
* [NodeJS](https://nodejs.org/en/)
* [Express](https://expressjs.com/)

Fetch App
* [Go Lang](https://go.dev)
* [Echo](https://echo.labstack.com/)

## Installation

Clone this repository
```
https://github.com/zikrykr/efish-backend-test.git
```

### Auth App

Go to directory repo
```bash
cd auth
```

To set up the application, you need to have `docker` installed in your machine or `node:17.8.0` if you want to run the app directly in your local machine.

Once you have them installed, create .env file inside `auth` folder
```
PORT=<must_define>
JWT_SECRET=<must_define>
```

change all value must-define with value that you wanted

There are 2 ways to run the app,
First method is to run the application with `docker-compose`
```
docker compose up --build
```

Second, is to run the application with `npm start`
```
npm start
```

open one of the api:</br>
http://localhost:port/api/v1/auth/register</br>
http://localhost:port/api/v1/auth/login</br>
http://localhost:port/api/v1/auth/verify-token

### Fetch App

Go to directory repo
```bash
cd fetch
```

To set up the application, you need to have `docker` installed in your machine or `go1.13` if you want to run the app directly in your local machine.

Once you have them installed, create .env file inside `fetch` folder
```
APP_PORT=<must_define>
GIN_PORT=<must_define_different_port>

SECRET_KEY=<must_define>

RESOURCE_URL=https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list
CURRENCY_CONVERTER_URL=https://free.currconv.com/api/v7/convert
CURRENCY_CONVERTER_API_KEY=8e3cda6dc61a6c6f13e6

CACHE_DURATION=15 #in minutes
```
change all value must-define with value that you wanted

There are 2 ways to run the app,
First method is to run the application with `docker-compose`
```
docker compose up --build
```

Second, you can run directly without docker
```
$ go mod tidy
$ go run main.go
```
open one of the api:</br>
http://localhost:port/api/v1/fetch/resources</br>
http://localhost:port/api/v1/fetch/resources/aggregate</br>
http://localhost:port/api/v1/fetch/verify-token


## Apps Diagram

![Context Diagram](https://raw.githubusercontent.com/zikrykr/efish-backend-test/main/app%2Bdiagram.jpg)
