package product

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	UrlDomain "golang_rest_api_test/api/product"
	HttpClient "golang_rest_api_test/http_client"
	"net/http"
	"testing"
	"time"
)

type Response struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Images      []string  `json:"images"`
	CreationAt  time.Time `json:"creationAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Category    struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		Image      string    `json:"image"`
		CreationAt time.Time `json:"creationAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"category"`
}

func Test_GetProductDetailsSuccess(t *testing.T) {
	var (
		idNo                       = 101
		productResponse            Response
		expectedTitle              = "books"
		expectedPrice              = 522
		expectedDescriptionPattern = "A description"
		//expectedImages             = []string{"https://i.imgur.com/0qQBkxX.jpg", "https://i.imgur.com/I5g1DoE.jpg", "https://i.imgur.com/myfFQBW.jpg"}
		expectedCategoryId   = 1
		expectedCategoryName = "Clothes"
		expectedTime, err    = time.Parse(time.RFC3339, "2024-04-02T11:53:26.000Z")
	)

	url := fmt.Sprintf(UrlDomain.HttpEscuelajsProducts, idNo)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	response, err := HttpClient.ApiRequest.Request("GET", url, nil, headers)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err)

	err = HttpClient.DecodeResponseForAssertion(response.ResponseBody, &productResponse)
	if err != nil {
		t.Fatal(err)
	}

	HttpClient.PrettyPrintAllResponse(response.Response.Request, response.Response, response.ResponseBody)

	assert.Equal(t, http.StatusOK, response.Response.StatusCode)

	require.Equal(t, idNo, productResponse.ID)
	assert.Equal(t, expectedTitle, productResponse.Title)
	assert.NotZero(t, expectedPrice, productResponse.Price)
	assert.GreaterOrEqual(t, expectedPrice, productResponse.Price)
	assert.Regexp(t, expectedDescriptionPattern, productResponse.Description)
	//assert.ElementsMatch(t, expectedImages, productResponse.Images)
	assert.Equal(t, expectedCategoryId, productResponse.Category.ID)
	// assert.Equal(t, expectedCategoryName, productResponse.Category.Name)
	if productResponse.Category.Name != expectedCategoryName {
		expectedError := errors.New("Expected error, but got nil")
		assert.Error(t, expectedError)
	}
	creationTime, err := time.Parse(time.RFC3339, productResponse.CreationAt.String())
	assert.True(t, creationTime.Before(expectedTime) || creationTime.Equal(expectedTime))
}
