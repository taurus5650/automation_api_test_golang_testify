package users_test

import (
	UrlDomain "automation_api_test_golang_testify/api/users"
	HttpClient "automation_api_test_golang_testify/http_client"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ReqeustIsAvailable struct {
	Email string `json:"email"`
}

type ResponseIsAvailable struct {
	IsAvailable bool `json:"isAvailable"`
}

func Test_EmailIsAvailableFalse(t *testing.T) {
	var (
		email               = "false@cannot.com"
		expectedIsAvailable = false
		responseIsAvailable ResponseIsAvailable
	)

	request := ReqeustIsAvailable{
		Email: email,
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	url := UrlDomain.HttpEscuelajsUsersIsAvailable
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := HttpClient.ApiRequest.Request("POST", url, requestBytes, headers)
	if err != nil {
		t.Fatal(err)
	}
	HttpClient.PrettyPrintAllResponse(response.Response.Request, response.Response, response.ResponseBody)

	err = HttpClient.DecodeResponseForAssertion(response.ResponseBody, &responseIsAvailable)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedIsAvailable, responseIsAvailable.IsAvailable)

}
