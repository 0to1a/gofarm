# What is GoFarm
**GoFarm** is an Application Development Framework for especially Backend Developer with Golang. Our goal is to develop easier, standardized, and faster than you code from scratch. Simplify interface and logic structure for easy-to-maintain project.

**GoFarm** lets you create better project, decrease learning curve and minimizing the amount of code needed. Let's contribute and make it better :)

## Installation
We need some components to support Build & Deploy working well. Need to install:

For GoFarm Assistant / CLI:
```
Still in progress
```
For database generator:
```
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/blob/master/drivers/sqlboiler-mysql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest
```

## Usage
Git clone from this repository or download here. To build use:
```
go build .
```
For download all dependencies, you can use:
```
go get ./...
```

## Concept
- **Rapid Development:** Analyzed behavior & Designed for Backend Engineer. Make better to maintain and rapid development.
- **Performance First:** Framework with a small footprint, make it not costly for production-based.
- **Dual-Architecture:** Using Microservices modular concept, but can deploy as single apps / Monolith. Just move module to new project, apps will separate & being microservice.
- **Structured:** Single pattern of writing code (like MVC), make other developer read & maintain easily. Minimize learning curve is our focus.
- **Clean Deployment:** Focus on clean structure for make sure better deployment.

## Requirement
Golang version 1.16+ and Go-Mod supported.

## Feature on Next Development
- [ ] Better logging system
- [ ] GoFarm Assistant, like PHP Artisan-Laravel / Lumen
- [x] Continuable framework scaling & dependency module
- [ ] Unit test, framework coverage
- [ ] Unit test, project structure
- [x] Database migration pattern / tools
- [ ] Postgres database connection
- [ ] Sqlite database connection
- [ ] Usage Documentation & Best-practice guidance
- [ ] Support Docker deployment
- [ ] Serverless Architecture deployment
- [x] Redis connection
- [x] Built-in cron job

## Changelog and New Feature
You can find on GitHub's Release

## Looking to contribute? Try to follow these guidelines:
- Use issues for idea, bug, issue, or everything in your mind
- For small change, please use PR
- Help us on test, documentation, and create examples usage
- Please improve this project to make better framework

## Dependencies & Big Thanks for
- github.com/labstack/echo
- github.com/go-sql-driver/mysql
- github.com/doug-martin/goqu
- github.com/jmoiron/sqlx
- github.com/golang-jwt/jwt
- github.com/go-co-op/gocron
- github.com/go-redis/redis
- github.com/volatiletech/sqlboiler
