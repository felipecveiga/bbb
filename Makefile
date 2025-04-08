run:
	@echo "Iniciando..."
	@go run main.go

setup:
	@echo "Instalando dependencias"
	@go get -u gorm.io/gorm
	@go get -u gorm.io/driver/mysql
	@go get github.com/labstack/echo/v4
	@go get github.com/labstack/echo
	@go get github.com/joho/godotenv
	@go install go.uber.org/mock/mockgen@latest
	@go get github.com/DATA-DOG/go-sqlmock

mock-generate:
	@go generate ./...

test:
	@go test -v ./service
	