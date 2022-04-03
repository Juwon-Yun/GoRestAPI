#### this repo is service for Restful API

```go

if run main.go {
    go run main.go && GO111MODULE=off go run main.go
}

if init develop {
    go mod init && go mod verify && go run main.go
}

dependency(
    "go install github.com/smartystreets/goconvey"
    "go get -u github.com/gofiber/fiber/v2"
    "go get -u gorm.io/gorm"
    "go get -u gorm.io/driver/mysql"
    "go get -u golang.org/x/crypto/bcrypt"
)

```
