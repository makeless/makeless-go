# go-saas

Extendable SaaS Application

## Usage

```go
func main() {
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

	// api
	api := &saas_api.Api{
		Logger:   logger,
		Database: database,
		Tls:      nil,
		Port:     os.Getenv("API_PORT"),
		Mode:     os.Getenv("API_MODE"),
		RWMutex:  new(sync.RWMutex),
	}

	// extend
	api.Extend(func(api *saas_api.Api) {
		api.GetEngine().GET("test", func(context *gin.Context) {
			context.JSON(http.StatusOK, api.Response(nil, "test"))
		})
	})

	saas := &go_saas.Saas{
		License:  "license",
		Logger:   logger,
		Database: database,
		Api:      api,
		RWMutex:  new(sync.RWMutex),
	}

	if err := saas.Run(); err != nil {
		saas.GetLogger().Fatal(err)
	}
}
```