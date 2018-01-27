package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/usamaiqbal83/developer-test-1/externalservice"
)

type MockClient struct {
	mockGet func (id int) (*externalservice.Post, error)
	mockPost func (id int, post *externalservice.Post) (*externalservice.Post, error)
}

func (mockClient *MockClient) GET(id int) (*externalservice.Post, error) {
	if mockClient.mockGet != nil {
		return mockClient.mockGet(id)
	}
	return nil, nil
}

func (mockClient *MockClient) POST(id int, post *externalservice.Post) (*externalservice.Post, error) {
	if mockClient.mockPost != nil {
		return mockClient.mockPost(id, post)
	}
	return nil, nil
}

func TestPOSTCallsAndReturnsJSONfromExternalServicePOST(t *testing.T) {

	// Description
	//
	// Write a test that accepts a POST request on the server and sends it the
	// fake external service with the posted form body return the response.
	//
	// Use the externalservice.Client interface to create a mock and assert the
	// client was called <n> times.
	//
	// ---
	//
	// Server should receive a request on
	//
	//  [POST] /api/posts/:id
	//  application/json
	//
	// With the form body
	//
	//  application/x-www-form-urlencoded
	//	title=Hello World!
	//	description=Lorem Ipsum Dolor Sit Amen.
	//
	// The server should then relay this data to the external service by way of
	// the Client POST method and return the returned value out as JSON.
	//
	// ---
	//
	// Assert that the externalservice.Client#POST was called 1 times with the
	// provided `:id` and post body and that the returned Post (from
	// externalservice.Client#POST) is written out as `application/json`.

	var callCount = 0
	e := echo.New()
	mockClient := MockClient{mockPost: func (id int, post *externalservice.Post) (*externalservice.Post, error) {
		callCount += 1
		post.ID = id

		return post, nil
	}}

	server := Server{Echo: e, Client: &mockClient}

	req := httptest.NewRequest(echo.POST, "/api/posts/43", nil)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")

	if err := req.ParseForm(); err != nil {
		panic(err.Error())
	}

	req.Form.Add("title", "Hello World!")
	req.Form.Add("description", "Lorem Ipsum Dolor Sit Amen.")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("43")

	// execute get post method
	server.AddPost(c)

	// Assertions
	// response code testing
	if rec.Code != http.StatusCreated {
		t.Fatalf("Response code mismatch => got %d : expected %d ", rec.Code, http.StatusCreated)
	}

	// response body testing
	var wantBody = "{\"data\":{\"id\":43,\"title\":\"Hello World!\",\"description\":\"Lorem Ipsum Dolor Sit Amen.\"},\"message\":\"Post Created\",\"code\":201}"
	if wantBody != rec.Body.String() {
		t.Fatalf("Response body mismatch => got %s : expected %d ", rec.Body.String(), wantBody)
	}

	// Post method call count testing
	if callCount != 1 {
		t.Fatalf("Get Method call count not correct => got %d : expected %d ", callCount, 1)
	}
}

func TestPOSTCallsAndReturnsErrorAsJSONFromExternalServiceGET(t *testing.T) {
	// Description
	//
	// Write a test that accepts a GET request on the server and returns the
	// error returned from the external service.
	//
	// Use the externalservice.Client interface to create a mock and assert the
	// client was called <n> times.
	//
	// ---
	//
	// Server should receive a request on
	//
	//	[GET] /api/posts/:id
	//
	// The server should then return the error from the external service out as
	// JSON.
	//
	// The error response returned from the external service would look like
	//
	//	400 application/json
	//
	//	{
	//		"code": 400,
	//		"message": "Bad Request"
	//	}
	//
	// ---
	//
	// Assert that the externalservice.Client#GET was called 1 times with the
	// provided `:id` and the returned error (above) is output as the response
	// as
	//
	//	{
	//		"code": 400,
	//		"message": "Bad Request",
	//		"path": "/api/posts/:id
	//	}
	//
	// Note: *`:id` should be the actual `:id` in the original request.*

	var callCount = 0
	var gotId = 0
	e := echo.New()
	mockClient := MockClient{mockGet: func (id int) (*externalservice.Post, error) {
		callCount += 1
		gotId = id
		return nil, errors.New("error")
	}}

	server := Server{Echo: e, Client: &mockClient}

	req := httptest.NewRequest(echo.GET, "/api/posts/43", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("43")

	// execute get post method
	server.GetPost(c)

	// Assertions
	// response code testing
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("Response code mismatch => got %d : expected %d ", rec.Code, http.StatusBadRequest)
	}

	// response body testing
	var wantBody = "{\"path\":\"/api/posts/43\",\"message\":\"Bad Request\",\"code\":400}"
	if wantBody != rec.Body.String() {
		t.Fatalf("Response body mismatch => got %s : expected %d ", rec.Body.String(), wantBody)
	}

	// Get method call count testing
	if callCount != 1 {
		t.Fatalf("Get Method call count not correct => got %d : expected %d ", callCount, 1)
	}

	// Get method call with given id testing
	if gotId != 43 {
		t.Fatalf("Get Method call given id passing not correct => got %d : expected %d ", gotId, 43)
	}
}
