docker run --name=todo-list -e POSTGRES_PASSWORD='123456' -p 5432:54321 -d --rm postgres

migrate -path ./schema -database postgres://postgres:123456@localhost:54321/postgres?sslmode=disable up

basePath: /
definitions:
  handler.errorResponse:
    properties:
      message:
        type: string
    type: object
  handler.signInInput:
    properties:
      password:
        type: string
      username:
        type: string
    required:
    - password
    - username
    type: object
host: localhost:8000
info:
  contact: {}
  description: API server for TodoList Application
  title: Todo App API
  version: "1.0"
paths:
  /auth/sign-in:
    post:
      consumes:
      - application/json
      description: signin
      operationId: sign-in
      parameters:
      - description: account info
        in: body
        name: input
        required: true
        schema:
          $ref: '#/definitions/handler.signInInput'
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            type: integer
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/handler.errorResponse'
        "404":
          description: Bad Request
          schema:
            $ref: '#/definitions/handler.errorResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/handler.errorResponse'
        default:
          description: ""
          schema:
            $ref: '#/definitions/handler.errorResponse'
      summary: SignIn
      tags:
      - auth
swagger: "2.0"


SWAGGER DOCUMENTATION