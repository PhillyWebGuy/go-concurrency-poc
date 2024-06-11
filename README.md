# GO-CONCONCURRENCY-POC

The purpose of this POC is to demonstrate the use of:

- CompileDaemon
- GIN
- REST API that consumes and produces JSON
- Gorm
- Concurrency

# Starting the Application

cmd/http/CompileDaemon -build="go build ./main.go" -command="./main"

Point your browser: http://localhost:8080

# Branches / Milestones

## setup

This contains the commits that do the following:

- Add main, git ignore, basic directory structure
- Add postgres support
- Add Gorm support
- Add Gin
- Add godotenv
- Add CompileDaemon