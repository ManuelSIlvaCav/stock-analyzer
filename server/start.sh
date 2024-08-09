export ENV=dev
export PORT=3000
export POSTGRES_HOST=localhost

go build .

go run cmd/stock-analyzer/main.go