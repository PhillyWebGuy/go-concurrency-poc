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

In addition to haveing Go and Docker/Rancher installed, you will also need Air.

If you don't know what Air is, or have it installed, run this command from a terminal:

`go install github.com/air-verse/air@latest`

Now you should be ready to run your application.

First, run the PostgreSQL database by executing this command from the root of the project:

`make up`

Then, run the web server:

`make http`

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

## gorm

This containes the commits from previous milestones, plus the following:

- Add support for PostgreSQL database