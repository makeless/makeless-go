# Makeless

Makeless - SaaS Framework - Golang Implementation

[![License](https://img.shields.io/badge/license-polyform:noncommercial-blue)](https://polyformproject.org/licenses/noncommercial/1.0.0/)
[![Build Status](https://ci.loeffel.io/api/badges/makeless/makeless-go/status.svg)](https://ci.loeffel.io/makeless/makeless-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/makeless/makeless-go)](https://goreportcard.com/report/github.com/makeless/makeless-go)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmakeless%2Fmakeless-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmakeless%2Fmakeless-go?ref=badge_shield)
[![Discord](https://img.shields.io/discord/775684445314744341?label=discord)](https://discord.gg/K7Une7gndt) 

- Based on Golang ([gin](https://github.com/gin-gonic/gin) & [gorm](https://github.com/go-gorm/gorm))
- Concurrency safe & scalable
- Super clean and small
- Fully customizable and configurable
- Multilingual
- State of the art Authentication with JWT HttpOnly Cookies
- User management
- Team management
- Token management for users and teams
- Realtime events
- Subscriptions and Per-Seat Payments out of the box (coming soon)

## Preview

<img src="https://raw.githubusercontent.com/makeless/makeless-ui/master/preview.png" alt="preview">

## Frontend

- TypeScript & Vue.js: [https://github.com/makeless/makeless-ui](https://github.com/makeless/makeless-ui)

## Demo

- Go + TypeScript & Vue.js: [https://github.com/makeless/makeless-demo](https://github.com/makeless/makeless-demo)

## Usage

```go
package main

import (
	"os"
	"strings"
	"sync"
	"time"

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
	"gorm.io/driver/mysql"
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

	// router
	router := &makeless_go_http_basic.Router{
		RWMutex: new(sync.RWMutex),
	}

	// http
	http := &makeless_go_http_basic.Http{
		Router:        router,
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

	makeless := &makeless_go.Makeless{
		Config:   config,
		Logger:   logger,
		Mailer:   mailer,
		Database: database,
		Http:     http,
		RWMutex:  new(sync.RWMutex),
	}

	if err := makeless.Init(mysql.Open(database.GetConnectionString()), "./makeless.json"); err != nil {
		makeless.GetLogger().Fatal(err)
	}

	if err := makeless.Run(); err != nil {
		makeless.GetLogger().Fatal(err)
	}
}
```

## Config

Demo [makeless.json](https://github.com/makeless/makeless-demo/blob/master/makeless.json)

## License

- Makeless is licensed under the [Polyform Noncommercial](https://polyformproject.org/licenses/noncommercial/1.0.0/) license.  
- Exemption: Every contributor gets free access to a commercial license.  
- Please contact lucas@loeffel.io for a commercial license.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmakeless%2Fmakeless-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmakeless%2Fmakeless-go?ref=badge_large)
