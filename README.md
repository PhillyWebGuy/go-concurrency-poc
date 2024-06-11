# GO-CONCONCURRENCY-POC

The purpose of this POC is to demonstrate the use of:

- Air
- Gin
- REST API that consumes and produces JSON
- Gorm
- Concurrency

# Starting the Application

Create a `.env` file in the root, and include this:

`
PORT=3000
`

Run this command from a terminal:

`go install github.com/air-verse/air@latest`

Run command at top level of project to utilized air

`air`

Point your browser: http://localhost:3000

# Branches / Milestones

## setup

This contains the commits that do the following:

- Add main, git ignore, basic directory structure
- Add postgres support
- Add Gorm support
- Add Gin
- Add godotenv
- Add Air