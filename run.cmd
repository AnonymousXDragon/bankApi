REM windows

if "%1" == "" (
    echo No Command specified select : [ migrate:up / migrate:down / create:migration / sqlc:genrate ]
    exit /b 1
)

REM migrate the db into database
if "%1" == "migrate:up" (
    echo migration running
    migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/bankdb?sslmode=disable" -verbose up
    echo migration successfully finished
    exit /b 0
)

REM removing the mogration
if "%1" == "migrate:down" (
    echo removing migration 
    migrate -path db/migration -database "postgres://postgres:password@localhost:5432/bankdb?sslmode=disable" -verbose down
    echo migration successfully removed 
    exit /b 0
)

REM create migration
if "%1" == "create:migration"(
    migrate -ext postgres -dir db/migration -seq init_schema
    exit /b 0
)

REM genrate queries
if "%1" == "sqlc:generate" (
    echo generating queries
    slqc generate
    exit /b 0
)

if "%1" == "test" (
    echo processing all unit tests
    go test -v -cover ./...
)