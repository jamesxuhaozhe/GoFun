package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errorPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errorUser(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errorNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errorPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errorUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprint(writer, "no error")
	return nil
}

var tests = []struct {
	handler appHandler
	code    int
	message string
}{
	{errorPanic, 500, "Internal Server Error"},
	{errorUser, 400, "user error"},
	{errorNotFound, 404, "Not Found"},
	{errorPermission, 403, "Forbidden"},
	{errorUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrorWrapper(t *testing.T) {
	for _, testCase := range tests {
		f := errWrapper(testCase.handler)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)

		verifyResponse(response.Result(), testCase.code, testCase.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.handler)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMessage string, t *testing.T) {
	byteSlice, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(byteSlice), "\n")
	if resp.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("Expected (%d, %s); "+"got (%d, %s)", expectedCode, expectedMessage, resp.StatusCode, body)
	}
}
