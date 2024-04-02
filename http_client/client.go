package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpApiRequestStruct struct {
}

type ApiResponse struct {
	Response     *http.Response
	ResponseBody []byte
	Err          error
}

type HttpApiRequestInterface interface {
	Request(method, url string, body []byte, headers map[string]string) (*ApiResponse, error)
	Get(method, url string, headers map[string]string) (*ApiResponse, error)
	Post(method, url string, body []byte, headers map[string]string) (*ApiResponse, error)
	Delete(method, url string, body []byte, headers map[string]string) (*ApiResponse, error)
	Put(method, url string, body []byte, headers map[string]string) (*ApiResponse, error)
}

var (
	ApiRequest HttpApiRequestInterface = &HttpApiRequestStruct{}
)

func (req *HttpApiRequestStruct) Request(method, url string, body []byte, headers map[string]string) (*ApiResponse, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	readResponseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &ApiResponse{
		Response:     response,
		ResponseBody: readResponseBody,
		Err:          nil,
	}, nil

}

func (req *HttpApiRequestStruct) Get(method, url string, headers map[string]string) (*ApiResponse, error) {
	return req.Request("GET", url, nil, headers)
}

func (req *HttpApiRequestStruct) Post(method, url string, body []byte, headers map[string]string) (*ApiResponse, error) {
	return req.Request("POST", url, body, headers)
}

func (req *HttpApiRequestStruct) Delete(method, url string, body []byte, headers map[string]string) (*ApiResponse, error) {
	return req.Request("DELETE", url, body, headers)
}

func (req *HttpApiRequestStruct) Put(method, url string, body []byte, headers map[string]string) (*ApiResponse, error) {
	return req.Request("PUT", url, body, headers)
}

func DecodeResponseForAssertion(rawResponseBody []byte, target interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(rawResponseBody))
	return decoder.Decode(target)
}

func PrettyPrintJson(jsonBody []byte) error {
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, jsonBody, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(prettyJson.String())
	return nil
}

func PrettyPrintAllResponse(request *http.Request, response *http.Response, responseBody []byte) {
	fmt.Println("************************************")
	fmt.Println("🐞 debug prints")
	fmt.Println("--------------------------------")
	fmt.Printf("* datetime: %s\n", time.Now().Format("2006/01/02 15:04:05"))
	fmt.Printf("* request: [%s] %s\n", request.Method, request.URL)
	fmt.Println("* headers:", request.Header)
	if request.Body != nil {
		bodyBytes, _ := ioutil.ReadAll(request.Body)
		fmt.Println("* body:")
		PrettyPrintJson(bodyBytes)
	}
	fmt.Printf("* status: %d\n", response.StatusCode)
	fmt.Println("* response:")
	PrettyPrintJson(responseBody)
	fmt.Println("************************************")
}
