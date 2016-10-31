package developertest

import (
	"testing"
)

func TestPOSTCallsAndReturnsJSONfromExternalServicePOST(t *testing.T) {
	// Descirption
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
}
