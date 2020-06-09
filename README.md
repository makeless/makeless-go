# go-saas

SaaS Golang Framework

[![Build Status](https://ci.loeffel.io/api/badges/go-saas/go-saas/status.svg)](https://ci.loeffel.io/go-saas/go-saas)

## Usage

```go
func main() {
<<<<<<< Updated upstream
    // logger
    logger := new(saas_logger_stdio.Stdio)
    
    // database
    database := &saas_database.Database{
        Dialect:  "mysql",
        Host:     os.Getenv("DB_HOST"),
        Database: os.Getenv("DB_NAME"),
        Port:     os.Getenv("DB_PORT"),
        Username: os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASS"),
        RWMutex:  new(sync.RWMutex),
    }
    
    // security
    security := &saas_security_basic.Basic{
        Database: database,
        RWMutex:  new(sync.RWMutex),
    }
    
    // jwt
    jwt := &saas_api.Jwt{
        Key:     os.Getenv("JWT_KEY"),
        RWMutex: new(sync.RWMutex),
    }
    
    // event
    event := &saas_event_basic.Event{
        Hub:     new(saas_event_basic.Hub).Init(),
        RWMutex: new(sync.RWMutex),
    }
    
    // api
    api := &saas_api.Api{
        Logger:   logger,
        Event:    event,
        Security: security,
        Database: database,
        Origins:  strings.Split(os.Getenv("ORIGINS"), ","),
        Jwt:      jwt,
        Tls:      nil,
        Port:     os.Getenv("API_PORT"),
        Mode:     os.Getenv("API_MODE"),
        RWMutex:  new(sync.RWMutex),
    }
    
    saas := &go_saas.Saas{
        License:  "abc",
        Logger:   logger,
        Database: database,
        Api:      api,
        RWMutex:  new(sync.RWMutex),
    }
    
    if err := saas.Run(); err != nil {
        saas.GetLogger().Fatal(err)
    }
=======
	// logger
	logger := new(go_saas_basic_logger.Logger)

	// config
	config := &go_saas_basic_config.Config{
		RWMutex: new(sync.RWMutex),
	}

	// database
	database := &go_saas_basic_database.Database{
		Dialect:  "mysql",
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		RWMutex:  new(sync.RWMutex),
	}

	// security
	security := &go_saas_basic_security.Security{
		Database: database,
		RWMutex:  new(sync.RWMutex),
	}

	// event hub
	hub := &go_saas_basic_event.Hub{
		List:    make(map[uint]map[uint]chan sse.Event),
		RWMutex: new(sync.RWMutex),
	}

	// event
	event := &go_saas_basic_event.Event{
		Hub:     hub,
		RWMutex: new(sync.RWMutex),
	}

	// jwt authenticator
	authenticator := &go_saas_basic_authenticator.Authenticator{
		Security:    security,
		Realm:       "auth",
		Key:         os.Getenv("JWT_KEY"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		RWMutex:     new(sync.RWMutex),
	}

	// http
	http := &go_saas_basic_http.Http{
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
>>>>>>> Stashed changes
}
```