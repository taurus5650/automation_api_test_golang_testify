package product

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	UrlDomain "golang_rest_api_test/api/product"
	HttpClient "golang_rest_api_test/http_client"
	"testing"
	"time"
)

type RequestPostProduct struct {
	Title       string   `json:"title"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	CategoryID  int      `json:"categoryId"`
	Images      []string `json:"images"`
}

type ResponsePostProduct struct {
	Title       string   `json:"title"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Category    struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		Image      string    `json:"image"`
		CreationAt time.Time `json:"creationAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"category"`
	ID         int       `json:"id"`
	CreationAt time.Time `json:"creationAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func GetTimeNow() string {
	now := time.Now().Format("2006-01-02 15:04:05")
	return now
}

func Test_CreateNewProductSuccess(t *testing.T) {
	var (
		title               = "NewProduct"
		price               = 10
		description         = fmt.Sprintf("NewProduct_%s", GetTimeNow())
		categoryID          = 1
		images              = []string{"https://placeimg.com/640/480/any"}
		responsePostProduct ResponsePostProduct
	)

	request := RequestPostProduct{
		Title:       title,
		Price:       price,
		Description: description,
		CategoryID:  categoryID,
		Images:      images,
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	url := UrlDomain.HttpEscuelajsCreateProducts
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := HttpClient.ApiRequest.Request("POST", url, requestBytes, headers)
	if err != nil {
		t.Fatal(err)
	}
	HttpClient.PrettyPrintAllResponse(response.Response.Request, response.Response, response.ResponseBody)

	err = HttpClient.DecodeResponseForAssertion(response.ResponseBody, &responsePostProduct)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, title, responsePostProduct.Title)

	// TearDown
	deleteUrl := fmt.Sprintf(UrlDomain.HttpEscuelajsProducts, responsePostProduct.ID)
	deleteProduct, err := HttpClient.ApiRequest.Request("DELETE", deleteUrl, nil, headers)
	HttpClient.PrettyPrintAllResponse(deleteProduct.Response.Request, deleteProduct.Response, deleteProduct.ResponseBody)

}
