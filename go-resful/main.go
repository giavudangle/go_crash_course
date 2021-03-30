package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1 // increment
)

//----------
// Handlers
//----------

func createUser(ctx echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := ctx.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return ctx.JSON(http.StatusCreated, u)
}

func getUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return ctx.JSON(http.StatusOK, users[id])
}

func updateUser(ctx echo.Context) error {
	u := new(user)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	users[id].Name = "Vudang"
	return ctx.JSON(http.StatusOK, users[id])
}

func deleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	delete(users, id)
	return ctx.NoContent(http.StatusNoContent)
}

func main() {
	app := echo.New()

	// Middlewares
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	// Routers
	app.POST("/users", createUser)
	app.GET("/users/:id", getUser)
	app.PUT("/users/:id", updateUser)
	app.DELETE("/users/:id", deleteUser)

	// Start server
	app.Logger.Fatal(app.Start(":10000"))
}
