package common

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response is very weird in this application because it was matched
// to mimic the previous service which was made in Java's Spring boot.
// The HTTP status will always be 200 irrespective of the error. But
// the Code inside Response will be the actual HTTP status code you
// wanted to convey. Dumb, I know.
//
// Errors can be used to return validation errors or any error that
// has to be in a data structure for itself. If the error you want to
// return is more like a string of text, then use Message with correct
// Status and Code when constructing Response struct
type Response struct {
	Status  string       `json:"status"`
	Message string       `json:"statusMessage"`
	Code    int          `json:"statusCode"`
	Data    ResponseData `json:"data,omitempty"`
	Errors  interface{}  `json:"errors,omitempty"`
}

// ResponseData will house the real Content of the application. If the
// app does not have any response, then make sure you don't send messages
// here instead use Response.Message.
//
// Pagination is a whole different story, checkout Pagination docs for more.
type ResponseData struct {
	Content    interface{}        `json:"content"`
	Pagination PaginationResponse `json:"pagination,omitempty"`
}

// ResponseType is an int64 with two possible values. SUCCESS and ERROR.
// Since golang has nothing called enums, we have to make do with manually
// Build and convert it into a string that can be sent to Response.Status
type ResponseType int64

const (
	SUCCESS ResponseType = iota
	ERROR
)

// Build is responsible to convert ResponseType which is of type int64 into
// a string which can be passed along to Response.Status
func (t ResponseType) Build() *Response {
	switch t {
	case SUCCESS:
		return &Response{Status: "SUCCESS"}
	case ERROR:
		return &Response{Status: "ERROR"}
	default:
		return &Response{Status: "Unknown"}
	}
}

// NewResponse is the constructor for the whole Response builder.
// It takes minimal required values and builds the most minimal
// Response struct that can be sent to client
func NewResponse(status ResponseType, message string) (res *Response) {
	res = status.Build()
	res.Message = message

	return
}

// SetContent sets ResponseData.Content inside Response.Data
func (r *Response) SetContent(data interface{}) *Response {
	r.Data.Content = data

	return r
}

// SetError sets errors in your application. Got an error?
// This is the guy you need to depend on.
func (r *Response) SetError(errors interface{}) *Response {
	r.Errors = errors

	return r
}

// SetPagination sets pagination data. Maybe you somehow figured
// out to properly paginate? Good luck, have fun this :)
func (r *Response) SetPagination(pagination PaginationResponse) *Response {
	r.Data.Pagination = pagination

	return r
}

// Respond is the guy who takes out the whole Response struct,
// assumes you have set all required data into it, marshals the
// whole thing into JSON and pushes it to client.
//
// It will set content-type of the response, set the header.
//
// This will be the last step in response building. Thus, it's the
// only method that can communicate with http.ResponseWriter
func (r *Response) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Status {
	case "SUCCESS":
		r.Code = http.StatusOK
	case "ERROR":
		r.Code = http.StatusBadRequest
	}

	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Fatal("Unable to build response and send")
	}
}
