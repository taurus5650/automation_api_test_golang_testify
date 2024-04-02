# Testing Restful API with Golang Testify

## Testdata
### Purpose
- For testing, which using
  - Golang test
  - Testify
- Support testing
  - Restful api testing (Testing resource : https://fakeapi.platzi.com/en/about/introduction/)

## Directory Structure
```
git ls-tree -r --name-only HEAD | tree --fromfile

├── README.md
├── api
│   └── product
... ...
│       └── const.go
├── go.mod
├── go.sum
├── http_client
│   └── client.go
├── readme
│   └── getproduct_result.png
└── test_suites
    └── product
        ├── deleteProduct_test.go
        ├── getProduct_test.go
        └── postProduct_test.go
        ... ...
```

## Getting started
Pls input the command before testing.
```
$ go mod init automation_api_test_golang_testify
$ go mod tidy
```

## Testing
After the code completed, then run the command to launch the API test.
```
$ go test -v ./test_suites/product/getProduct_test.go
```
![getproduct_result.png](readme%2Fgetproduct_result.png)
