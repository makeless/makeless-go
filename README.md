# go-saas

Go SaaS Framework - Golang Implementation

[![Build Status](https://ci.loeffel.io/api/badges/go-saas/go-saas/status.svg)](https://ci.loeffel.io/go-saas/go-saas)

- Based on Golang ([gin](https://github.com/gin-gonic/gin) & [gorm](https://github.com/go-gorm/gorm))
- Super clean and small
- Fully customizable and configurable
- State of the art Authentication with JWT HttpOnly Cookies
- User management
- Team management
- Token management for users and teams
- Subscriptions and Per-Seat Payments out of the box (coming soon)

## Preview

<img src="https://raw.githubusercontent.com/go-saas/go-saas-ui/master/preview.png" alt="logo">

## Frontend

- TypeScript & Vue.js: [https://github.com/go-saas/go-saas-ui](https://github.com/go-saas/go-saas-ui)

## Demo

- Go + TypeScript & Vue.js: [https://github.com/go-saas/go-saas-demo](https://github.com/go-saas/go-saas-demo)

## Usage

```go
package main

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas"
	"github.com/go-saas/go-saas/authenticator/basic"
	"github.com/go-saas/go-saas/config/basic"
	"github.com/go-saas/go-saas/database/basic"
	"github.com/go-saas/go-saas/event/basic"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/http/basic"
	"github.com/go-saas/go-saas/logger/basic"
	"github.com/go-saas/go-saas/security/basic"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// logger
	logger := new(go_saas_logger_basic.Logger)

	// config
	config := &go_saas_config_basic.Config{
		RWMutex: new(sync.RWMutex),
	}

	// database
	database := &go_saas_database_basic.Database{
		Dialect:  "mysql",
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		RWMutex:  new(sync.RWMutex),
	}

	// security
	security := &go_saas_security_basic.Security{
		Database: database,
		RWMutex:  new(sync.RWMutex),
	}

	// event hub
	hub := &go_saas_event_basic.Hub{
		List:    make(map[uint]map[uint]chan sse.Event),
		RWMutex: new(sync.RWMutex),
	}

	// event
	event := &go_saas_event_basic.Event{
		Hub:     hub,
		RWMutex: new(sync.RWMutex),
	}

	// jwt authenticator
	authenticator := &go_saas_authenticator_basic.Authenticator{
		Security:    security,
		Realm:       "auth",
		Key:         os.Getenv("JWT_KEY"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		RWMutex:     new(sync.RWMutex),
	}

	// http
	http := &go_saas_http_basic.Http{
		Router:        gin.Default(),
		Handlers:      make(map[string]func(http go_saas_http.Http) error),
		Logger:        logger,
		Event:         event,
		Authenticator: authenticator,
		Security:      security,
		Database:      database,
		Tls:           nil,
		Origins:       strings.Split(os.Getenv("ORIGINS"), ","),
		Headers:       []string{"Team"},
		Port:          os.Getenv("API_PORT"),
		Mode:          os.Getenv("API_MODE"),
		RWMutex:       new(sync.RWMutex),
	}

	saas := &go_saas.Saas{
		Config:   config,
		Logger:   logger,
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
  "tokens": true,
  "teams": {
    "tokens": false
  },
  "navigation": {
    "left": {
      "en": [
        {"label": "Dashboard","to": "dashboard"}
      ]
    },
    "right": {
      "en": [
        {"label": "Login","to": "login"}
      ]
    }
  }
}
```
