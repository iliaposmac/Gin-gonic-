docker run --name=todo-list -e POSTGRES_PASSWORD='123456' -p 5432:54321 -d --rm postgres

migrate -path ./schema -database postgres://postgres:123456@localhost:54321/postgres?sslmode=disable up

SWAGGER DOCUMENTATION