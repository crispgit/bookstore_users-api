package app

import (
	"github.com/crispgit/bookstore_users-api/controllers/ping"
	"github.com/crispgit/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	//PUT update all the json, PATCH update part of json
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
