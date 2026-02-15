# Go Spine Backend Clean Architecture


## Framework / Library
- **[Spine]** 
- **Bun**
- **testify**
- **mockery**
- **viper**
- **PostgreSQL**
- **jwt**
- **bcrypt**
- Check more packages in go.mod.

## How to Run

- clone

```bash
cd your-workspace

git clone https://github.com/janghanul090801/spine-clean-architecture.git

cd spine-clean-architecture
```
- run
```bash
make run 
```

### Run Test

```bash
make test
```

### The Complete Project Folder Structure

```
.
├── api/
│   ├── controller/
│   │   ├── auth_controller.go
│   │   ├── profile_controller.go
│   │   ├── profile_controller_test.go
│   │   └── task_controller.go
│   └── route/
│       ├── auth_route.go
│       ├── profile_route.go
│       └── task_route.go
├── cmd/
│   └── main.go
├── config/
│   └── env.go
├── domain/
│   ├── mocks/
│   │   ├── AuthUseCase.go
│   │   ├── ProfileUsecase.go
│   │   ├── TaskRepository.go
│   │   ├── TaskUsecase.go
│   │   └── UserRepository.go
│   ├── auth.go
│   ├── domain.go
│   ├── error_response.go
│   ├── jwt_custom.go
│   ├── profile.go
│   ├── success_response.go
│   ├── task.go
│   └── user.go
├── infra/
│   ├── database/
│   │   └── database.go
│   ├── migrations/
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── 00000000_init.go
│   │   └── migrations.go
│   ├── model/
│   │   ├── task_model.go
│   │   └── user_model.go
│   └── repository/
│       ├── task_repository.go
│       ├── user_repository.go
│       └── user_repository_test.go
├── interceptor/
│   ├── auth_interceptor.go
│   ├── cors_interceptor.go
│   ├── error_interceptor.go
│   ├── logging_interceptor.go
│   ├── rate_limiter.go
│   ├── security_headers.go
│   └── tx_interceptor.go
├── internal/
│   ├── fakeutil/
│   │   └── fakeutil.go
│   ├── logger/
│   │   └── logger.go
│   └── token/
│       └── token.go
├── usecase/
│   ├── auth_usecase.go
│   ├── profile_usecase.go
│   ├── task_usecase.go
│   └── task_usecase_test.go
├── .env.example
├── .git/...
├── .gitignore
├── .idea/...
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── help.ps1
├── LICENSE
├── Makefile
└── README.md
```

### API Request & Response

- signup

  - request

  ```
  curl --location --request POST 'http://localhost:8080/signup' \
  --data-urlencode 'email=test@gmail.com' \
  --data-urlencode 'password=test' \
  --data-urlencode 'name=Test Name'
  ```

  - response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- login

  - request

  ```
  curl --location --request POST 'http://localhost:8080/login' \
  --data-urlencode 'email=test@gmail.com' \
  --data-urlencode 'password=test'
  ```

  - response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- profile

  - request

  ```
  curl --location --request GET 'http://localhost:8080/profile' \
  --header 'Authorization: Bearer access_token'
  ```

  - response

  ```json
  {
    "name": "Test Name",
    "email": "test@gmail.com"
  }
  ```

- task create

  - request

  ```
  curl --location --request POST 'http://localhost:8080/task' \
  --header 'Authorization: Bearer access_token' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'title=Test Task'
  ```

  - response

  ```json
  {
    "message": "Task created successfully"
  }
  ```

- task fetch

  - request

  ```
  curl --location --request GET 'http://localhost:8080/task' \
  --header 'Authorization: Bearer access_token'
  ```

  - response

  ```json
  [
    {
      "title": "Test Task"
    },
    {
      "title": "Test Another Task"
    }
  ]
  ```

- refresh token

  - request

  ```
  curl --location --request POST 'http://localhost:8080/refresh_token' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'refreshToken=refresh_token'
  ```

  - response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

[Spine]: https://github.com/NARUBROWN/spine