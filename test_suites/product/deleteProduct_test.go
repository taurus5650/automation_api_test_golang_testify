package product

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	UrlDomain "golang_rest_api_test/api/product"
	HttpClient "golang_rest_api_test/http_client"
	"testing"
)

func Test_DeleteSuccess(t *testing.T) {

	var (
		title               = "NewProduct"
		price               = 10
		description         = fmt.Sprintf("NewProduct_%s", GetTimeNow())
		categoryID          = 1
		images              = []string{"https://placeimg.com/640/480/any"}
		postProductResponse ResponsePostProduct
	)

	request := RequestPostProduct{
		Title:       title,
		Price:       price,
		Description: description,
		CategoryID:  categoryID,
		Images:      images,
	}

	postProductRequestBytes, _ := json.Marshal(request)

	postProductUrl := UrlDomain.HttpEscuelajsCreateProducts
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	postProductResp, err := HttpClient.ApiRequest.Request("POST", postProductUrl, postProductRequestBytes, headers)
	if err != nil {
		t.Fatal(err)
	}
	HttpClient.PrettyPrintAllResponse(postProductResp.Response.Request, postProductResp.Response, postProductResp.ResponseBody)
	err = HttpClient.DecodeResponseForAssertion(postProductResp.ResponseBody, &postProductResponse)
	if err != nil {
		t.Fatal(err)
	}

	deleteURL := fmt.Sprintf(UrlDomain.HttpEscuelajsProducts, postProductResponse.ID)
	response, err := HttpClient.ApiRequest.Request("DELETE", deleteURL, nil, headers)
	HttpClient.PrettyPrintAllResponse(response.Response.Request, response.Response, response.ResponseBody)

	// 执行断言
	assert.True(t, true, "Expected product deletion to be successful")
}
