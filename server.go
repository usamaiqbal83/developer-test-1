package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/usamaiqbal83/developer-test-1/externalservice"
)

type ClientImpl struct {

}

func (c *ClientImpl) GET(id int) (*externalservice.Post, error) {
	//return nil, nil
	return nil, errors.New("post not found")
}

func (c *ClientImpl) POST(id int, post *externalservice.Post) (*externalservice.Post, error) {
	post.ID = id
	return post, nil
}

type Server struct {
	Client externalservice.Client
	Echo *echo.Echo
}

func main() {

	// initialization
	e := echo.New()
	client := ClientImpl{}
	server := Server{Client: &client, Echo: e}

	// add routes
	server.AddRoutes()

	// start server
	server.Run()
}

func (server *Server) Run() {
	server.Echo.Logger.Fatal(server.Echo.Start(":8080"))
}

func (server *Server) AddRoutes() {
	api := server.Echo.Group("/api")

	// routes
	api.POST("/posts/:id", server.AddPost)
	api.GET("/posts/:id", server.GetPost)
}

func (server *Server) AddPost(c echo.Context) error {
	// extract required values from request
	id := c.Param("id")
	idInt, er := strconv.Atoi(id)
	title := c.FormValue("title")
	desc := c.FormValue("description")

	// return error if id is not a valid integer
	if er != nil || title == "" || desc == ""{
		return server.CreateBadResponse(&c, http.StatusBadRequest, "Bad Request")
	}

	// create post
	post := externalservice.Post{Title: title, Description: desc}

	// process post create using client
	uPost, err := server.Client.POST(idInt, &post)
	if err != nil {
		return server.CreateBadResponse(&c, http.StatusBadRequest, "Bad Request")
	}

	// Marshal provided interface into JSON structure
	postData, err := json.Marshal(uPost)
	if err != nil {
		return server.CreateBadResponse(&c, http.StatusBadRequest, "Bad Request")
	}

	// on success
	return server.CreateSuccessResponse(&c, http.StatusCreated, "Post Created", postData)
}

func (server *Server) GetPost(c echo.Context) error {
	// extract required values from request
	id := c.Param("id")
	idInt, er := strconv.Atoi(id)

	// return error if id is not a valid integer
	if er != nil {
		return server.CreateBadResponse(&c, http.StatusBadRequest, "Bad Request")
	}

	// process post get using client
	post, err := server.Client.GET(idInt)
	if err != nil {
		return server.CreateBadResponse(&c, http.StatusBadRequest, "Bad Request")
	}

	// Marshal provided interface into JSON structure
	postData, err := json.Marshal(post)
	if err != nil {
		return server.CreateBadResponse(&c, http.StatusBadRequest, "Bad Request")
	}

	// on success
	return server.CreateSuccessResponse(&c, http.StatusCreated, "Post Created", postData)
}

func (server *Server) CreateSuccessResponse(c *echo.Context, responseCode int, message string, data []byte) error {

	localC := *c
	response := fmt.Sprintf("{\"data\":%s,\"message\":%q,\"code\":%d}", data, message, responseCode)
	return localC.JSONBlob(responseCode, []byte(response))
}

func (server *Server) CreateBadResponse(c *echo.Context, responseCode int, message string) error {
	localC := *c
	response := fmt.Sprintf("{\"path\":%q,\"message\":%q,\"code\":%d}", localC.Request().URL.Path, message, responseCode)
	return localC.JSONBlob(responseCode, []byte(response))
}
