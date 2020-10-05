# go-saas

Go SaaS Framework - Golang Implementation

[![Build Status](https://ci.loeffel.io/api/badges/makeless/makeless-go/status.svg)](https://ci.loeffel.io/makeless/makeless-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/makeless/makeless-go)](https://goreportcard.com/report/github.com/makeless/makeless-go)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgo-saas%2Fgo-saas.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgo-saas%2Fgo-saas?ref=badge_shield)

- Based on Golang ([gin](https://github.com/gin-gonic/gin) & [gorm](https://github.com/go-gorm/gorm))
- Super clean and small
- Fully customizable and configurable
- State of the art Authentication with JWT HttpOnly Cookies
- User management
- Team management
- Token management for users and teams
- Subscriptions and Per-Seat Payments out of the box (coming soon)

## Preview

<img src="https://raw.githubusercontent.com/makeless/makeless-go-ui/master/preview.png" alt="logo">

## Frontend

- TypeScript & Vue.js: [https://github.com/makeless/makeless-go-ui](https://github.com/makeless/makeless-go-ui)

## Demo

- Go + TypeScript & Vue.js: [https://github.com/makeless/makeless-go-demo](https://github.com/makeless/makeless-go-demo)

## Usage

```go
package main

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go"
	"github.com/makeless/makeless-go/authenticator/basic"
	"github.com/makeless/makeless-go/config/basic"
	"github.com/makeless/makeless-go/database/basic"
	"github.com/makeless/makeless-go/event/basic"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/http/basic"
	"github.com/makeless/makeless-go/logger/basic"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/mailer/basic"
	"github.com/makeless/makeless-go/security/basic"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// logger
	logger := new(makeless_go_logger_basic.Logger)

	// config
	config := &makeless_go_config_basic.Config{
		RWMutex: new(sync.RWMutex),
	}

	// mailer
	mailer := &makeless_go_mailer_basic.Mailer{
		Handlers: make(map[string]func(data map[string]interface{}) (makeless_go_mailer.Mail, error)),
		Host:     os.Getenv("MAILER_HOST"),
		Port:     os.Getenv("MAILER_PORT"),
		Identity: os.Getenv("MAILER_IDENTITY"),
		Username: os.Getenv("MAILER_USERNAME"),
		Password: os.Getenv("MAILER_PASSWORD"),
		RWMutex:  new(sync.RWMutex),
	}

	// database
	database := &makeless_go_database_basic.Database{
		Dialect:  "mysql",
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		RWMutex:  new(sync.RWMutex),
	}

	// security
	security := &makeless_go_security_basic.Security{
		Database: database,
		RWMutex:  new(sync.RWMutex),
	}

	// event hub
	hub := &makeless_go_event_basic.Hub{
		List:    new(sync.Map),
		RWMutex: new(sync.RWMutex),
	}

	// event
	event := &makeless_go_event_basic.Event{
		Hub:     hub,
		Error:   make(chan error),
		RWMutex: new(sync.RWMutex),
	}

	// jwt authenticator
	authenticator := &makeless_go_authenticator_basic.Authenticator{
		Security:    security,
		Realm:       "auth",
		Key:         os.Getenv("JWT_KEY"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		RWMutex:     new(sync.RWMutex),
	}

	// http
	http := &makeless_go_http_basic.Http{
		Router:        gin.Default(),
		Handlers:      make(map[string]func(http makeless_go_http.Http) error),
		Logger:        logger,
		Event:         event,
		Authenticator: authenticator,
		Security:      security,
		Database:      database,
		Mailer:        mailer,
		Tls:           nil,
		Origins:       strings.Split(os.Getenv("ORIGINS"), ","),
		Headers:       []string{"Team"},
		Port:          os.Getenv("API_PORT"),
		Mode:          os.Getenv("API_MODE"),
		RWMutex:       new(sync.RWMutex),
	}

	saas := &makeless.Saas{
		Config:   config,
		Logger:   logger,
		Mailer:   mailer,
		Database: database,
		Http:     http,
		RWMutex:  new(sync.RWMutex),
	}

	if err := saas.Init("./../go-saas.json"); err != nil {
		saas.GetLogger().Fatal(err)
	}

	if err := saas.Run(); err != nil {
		saas.GetLogger().Fatal(err)
	}
}
```

## Config

go-saas.json

```json
{
  "name": "Go SaaS",
  "logo": null,
  "locale": "en",
  "host": "http://localhost:3000",
  "mail": {
    "name": "Go SaaS",
    "logo": null,
    "from": "Go SaaS <info@go-saas.io>",
    "link": "https://localhost",
    "buttonColor": "#4586ab",
    "buttonTextColor": "#FFFFFF",
    "texts": {
      "en": {
        "greeting": "Hello",
        "signature": "Best Regards",
        "copyright": "Copyright © 2020 Go SaaS. All rights reserved."
      }
    }
  },
  "tokens": true,
  "teams": {
    "tokens": true
  },
  "navigation": {
    "left": {
      "en": [
        {
          "label": "Dashboard",
          "to": "dashboard"
        }
      ]
    },
    "right": {
      "en": [
        {
          "label": "GitHub",
          "to": "https://github.com/go-saas",
          "external": true
        },
        {
          "label": "Login",
          "to": "login"
        },
	{
	  "label": "Register",
	  "to": "registration"
	}
      ]
    }
  }
}
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgo-saas%2Fgo-saas.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgo-saas%2Fgo-saas?ref=badge_large)
