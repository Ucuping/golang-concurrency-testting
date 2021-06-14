package service

import (
	"errors"
	"go-concurrency-testting/server/client"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getRequestFunc func(url string) (*http.Response, error)
)

type clientMock struct{}

// mocking the client call, so we dont hit the real endpoint
func (cm *clientMock) GetValue(url string) (*http.Response, error) {
	return getRequestFunc(url)
}

func TestUsernameCheck_Success(t *testing.T) {
	urls := []string{
		"http://twitter.com/ucuping",
		"http://instagram.com/_ucuping",
		"http://hackerone.com/ucuping",
		"http://github.com/ucuping",
	}

	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}
	client.ClientCall = &clientMock{}

	result := UsernameService.UsernameCheck(urls)

	assert.NotNil(t, result)
	assert.EqualValues(t, len(result), 3)
}

func TestUsernameCheck_No_Match(t *testing.T) {
	urls := []string{
		"http://twitter.com/ucuping_no_match",
		"http://instagram.com/ucuping_no_match",
		"http://hackerone.com/ucuping_no_match",
		"http://github.com/ucuping_no_match",
	}
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
		}, nil
	}
	client.ClientCall = &clientMock{}

	result := UsernameService.UsernameCheck(urls)

	assert.EqualValues(t, len(result), 0)
}

func TestingUsernameCheck_Url_Invalid(t *testing.T) {
	urls := []string{
		"http://wrong.com/ucuping_no_match",
		"http://wrong.com/ucuping_no_match",
		"http://wrong.com/ucuping_no_match",
		"http://wrong.com/ucuping_no_match",
	}
	getRequestFunc = func(url string) (*http.Response, error) {
		return nil, errors.New("cant_access_resource")
	}
	client.ClientCall = &clientMock{}

	result := UsernameService.UsernameCheck(urls)

	assert.EqualValues(t, len(result), 0)
}
