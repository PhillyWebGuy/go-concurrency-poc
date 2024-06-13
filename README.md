# GO-CONCONCURRENCY-POC

The purpose of this POC is to demonstrate the use of:

- Air
- Gin
- REST API that consumes and produces JSON
- Gorm
- Concurrency

# Starting the Application

Create a folder named `.local` at the root, then create a `.env` file and place in that folder. Include this in the file:

```
ENV=DEV

DATABASE_PASSWORD=password
DATABASE_USER=user 
DATABASE_NAME=database
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_CONNECTION_RETRY=5

HTTP_PORT=:3001
HTTP_DOMAIN=localhost
```

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

In order to understand how to build out the complete project, which is merged into develop, follow the evolution
of the repo buy reviewing each branch/milestone in the following order.

## setup

This contains the commits that do the following:

- Add main, git ignore, basic directory structure
- Add postgres support
- Add Gorm support
- Add Gin
- Add godotenv
- Add Air

## postgres

This contains the commits from previous milestones, plus the following:

- Script to initialize Postgres container
- Creation of makefile to start/stop db container, start Gin

## infrastructure

This contains the commits from previous milestones, plus the following:

- Refactor of HTTP related code to increase abstraction
- Make use of viper to load and read configuration
- Expanded use of .env settings
- Introduction of infrastructure as abstraction layer