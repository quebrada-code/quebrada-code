

gen_docs:
	swag init -g internal/app/router/router.go

gen_dependecies:
	wire gen cmd/api/wire.go