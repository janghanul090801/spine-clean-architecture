# Go 백엔드 클린 아키텍처

Fiber, PostgreSQL, Ent ORM, JWT 인증 미들웨어, 테스트 및 Docker를 사용한 Go(Golang) 백엔드 클린 아키텍처 프로젝트입니다.

![Go 백엔드 클린 아키텍처](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-backend-clean-architecture.png?raw=true)

**이 프로젝트를 템플릿으로 사용하여 Go 언어로 백엔드 프로젝트를 구축할 수 있습니다.**

## 프로젝트의 아키텍처 계층

- 라우터 (Router)
- 컨트롤러 (Controller)
- 유스케이스 (Usecase)
- 리포지토리 (Repository)
- 도메인 (Domain)

![Go 백엔드 클린 아키텍처 다이어그램](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-backend-arch-diagram.png?raw=true)


- **Fiber**: Go 언어용 웹 프레임워크, Go에서 가장 빠른 HTTP 엔진인 Fasthttp를 기반으로 구축됨
- **Ent**: Go 언어용 엔티티 프레임워크
- **testify**: Testify는 테스트를 위한 여러가지 패키지가 포함된 툴
- **mockery**: 테스트에 사용되는 Golang용 모의 코드 자동 생성기
- **viper**: `.env` 파일에서 구성을 로드하는 데 사용
- **PostgreSQL**
- **jwt**
- **bcrypt**
- 기타 패키지는 `go.mod` 를 참고

### JWT 인증 미들웨어가 없는 공개 API 요청 흐름

![공개 API 요청 흐름](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-public-api-request-flow.png?raw=true)

### JWT 인증 미들웨어가 있는 비공개 API 요청 흐름

> 액세스 토큰 유효성 검사를 위한 JWT 인증 미들웨어.

![비공개 API 요청 흐름](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-private-api-request-flow.png?raw=true)

### 실행 방법

해당 프로젝트는 Docker를 사용하거나 사용하지 않고 실행할 수 있습니다. 여기서는 두 가지 방법을 모두 제공합니다.

- clone 

```bash
# 작업 공간으로 이동
cd your-workspace

# 프로젝트를 작업 공간으로 복제
git clone https://github.com/janghanul090801/go-backend-clean-architecture-fiber.git

# 프로젝트 루트 디렉토리로 이동
cd go-backend-clean-architecture-fiber
```

#### Docker 없이 실행

- 루트 디렉터리에 `.env.example`를 복사해 `.env` 파일을 만들고 값을 입력
- `go`가 설치되어 있지 않으면 설치
- `PostgreSQL`이 설치되어 있지 않으면 설치
- `.env` 파일에서 `DB_HOST`를 `localhost`로 변경(`DB_HOST=localhost`)
- `go run cmd/main.go` 또는 `make run`를 실행
- `http://localhost:8080`로 접속

#### Docker로 실행

- 루트 디렉터리에 `.env.example`를 복사해 `.env` 파일을 만들고 값을 입력
- Docker 및 Docker Compose를 설치
- `docker-compose up -d` 또는 `make compose-up`를 실행
- `http://localhost:8080`로 접속

### 테스트 실행

```bash
# 모든 테스트 실행
go test ./...
```

### 모의 코드 생성

테스트를 위해 Usecase, Repository 및 데이터베이스에 대한 모의 코드를 생성해야 합니다.

```bash
# Usecase 및 Repository 에 대한 모의 코드 생성
mockery --dir=domain --output=domain/mocks --outpkg=mocks --all
```

Usecase, Repository 또는 데이터베이스의 인터페이스를 변경할 때마다 해당 명령을 실행하여 테스트용 모의 코드를 다시 생성해야 합니다.

### 전체 프로젝트 폴더 구조

```
.
├── api/
│   ├── controller/
│   │   ├── login_controller.go
│   │   ├── profile_controller_test.go
│   │   ├── profile_controller.go
│   │   ├── refresh_token_controller.go
│   │   ├── signup_controller.go
│   │   └── task_controller.go
│   ├── middleware/
│   │   └── jwt_auth_middleware.go
│   └── route/
│       ├── login_route.go
│       ├── profile_route.go
│       ├── refresh_token_route.go
│       ├── signup_route.go
│       └── task_route.go
├── bootstrap/
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd/
│   └── main.go
├── domain/
│   ├── domain.go
│   ├── error_response.go
│   ├── jwt_custom.go
│   ├── login.go
│   ├── profile.go
│   ├── refresh_token.go
│   ├── signup.go
│   ├── success_response.go
│   ├── task.go
│   ├── user.go
│   └── mocks/
│       ├── LoginUsecase.go
│       ├── ProfileUsecase.go
│       ├── RefreshTokenUsecase.go
│       ├── SignupUsecase.go
│       ├── TaskRepository.go
│       ├── TaskUsecase.go
│       └── UserRepository.go
├── ent/
│   ├── client.go
│   ├── ent.go
│   ├── generate.go
│   ├── mutation.go
│   ├── runtime.go
│   ├── task_create.go
│   ├── task_delete.go
│   ├── task_query.go
│   ├── task_update.go
│   ├── task.go
│   ├── tx.go
│   ├── user_create.go
│   ├── user_delete.go
│   ├── user_query.go
│   ├── user_update.go
│   ├── user.go
│   ├── enttest/
│   │   └── enttest.go
│   ├── hook/
│   │   └── hook.go
│   ├── migrate/
│   │   ├── migrate.go
│   │   └── schema.go
│   ├── predicate/
│   │   └── predicate.go
│   ├── runtime/
│   │   └── runtime.go
│   ├── schema/
│   │   ├── task.go
│   │   └── user.go
│   ├── task/
│   │   ├── task.go
│   │   └── where.go
│   └── user/
│       ├── user.go
│       └── where.go
├── internal/
│   ├── fakeutil/
│   │   └── fakeutil.go
│   └── tokenutil/
│       └── tokenutil.go
├── repository/
│   ├── task_repository.go
│   ├── user_repository_test.go
│   └── user_repository.go
├── usecase/
│   ├── login_usecase.go
│   ├── profile_usecase.go
│   ├── refresh_token_usecase.go
│   ├── signup_usecase.go
│   ├── task_usecase_test.go
│   └── task_usecase.go
├── assets/
│   ├── button-view-api-docs.png
│   ├── go-arch-private-api-request-flow.png
│   ├── go-arch-public-api-request-flow.png
│   ├── go-backend-arch-diagram.png
│   └── go-backend-clean-architecture.png
├── .env.example
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── help.ps1
├── LICENSE
├── Makefile
└── README.md
```

### API 요청 및 응답 예시

- 회원가입 (signup)

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

- 로그인 (login)

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

- 프로필 (profile)

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

- 작업 생성 (task create)

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

- 작업 가져오기 (task fetch)

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

- 토큰 갱신 (refresh token)

  - request

  ```
  curl --location --request POST 'http://localhost:8080/refresh' \
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

