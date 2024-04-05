FROM golang:1.21.5

WORKDIR /app

CMD go mod init automation_api_test_golang_testify

CMD go mod tidy

COPY . .

ENV RUNNING_IN_DOCKER=true


CMD ["go", "test", "-v", "./..."]

CMD tail -f /dev/null
