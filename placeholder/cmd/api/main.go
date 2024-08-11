package main

import (
	"log"

	"{{module_name}}/internal"
	"{{module_name}}/src/handlers"
	"{{module_name}}/src/http"
	"{{module_name}}/src/repositories"
	"{{module_name}}/src/services"
)

const app = "{{app_name}}"


func main() {
	cfg := internal.GetConfig()

	if cfg.Database.IsMigration {
		internal.InitDatabaseClient()
	} else {
		_, err := internal.Init(&cfg.Database)
		if err != nil {
			log.Panicf("Failed to connect to database. Error: %s", err.Error())
		}
	}

	repoContainer := repositories.NewContainer()
	servicesWithRepos := services.GetServices(repoContainer)
	handlersWithServices := handlers.NewHandlerContainer(servicesWithRepos)

	http.NewServer(handlersWithServices).Start()
}
