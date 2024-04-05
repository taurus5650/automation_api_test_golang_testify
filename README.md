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

.
├── .gitignore
├── README.md
├── api
│   ├── product
│   │   └── const.go
│   └── users
│       └── const.go
├── deployments
│   └── Dockerfile
├── go.mod
├── go.sum
├── http_client
│   └── client.go
├── readme
│   └── getproduct_result.png
└── test_suites
    ├── product
    │   ├── deleteProduct_test.go
    │   ├── getProduct_test.go
    │   └── postProduct_test.go
    └── users
        └── isAvailable_test.go

        ... ...
```

## Step-by-step

### Before Testing
#### Option 1. Build up a simple Docker
Input the command
```
docker build -t testify_learn_image -f ./deployments/Dockerfile .

```

#### Option 2. Set up on local
Pls input the command before testing.
```
$ go mod init automation_api_test_golang_testify
$ go mod tidy
$ go mod download
```

###  Testing
After the code completed, then run the command to launch the API test.
```
$ go test -v ./...
OR
$ go test -v ./test_suites/product/getProduct_test.go
```
![getproduct_result.png](readme%2Fgetproduct_result.png)
