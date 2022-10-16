[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=0to1a_gofarm&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=0to1a_gofarm)
![Coverage](https://img.shields.io/badge/Coverage-85.9%25-brightgreen)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=0to1a_gofarm&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=0to1a_gofarm)
[![GitHub issues](https://img.shields.io/github/issues/0to1a/gofarm)](https://github.com/0to1a/gofarm/issues)
[![GitHub license](https://img.shields.io/github/license/0to1a/gofarm)](https://github.com/0to1a/gofarm/blob/master/LICENSE)
![Supported Go Versions](https://img.shields.io/badge/Go-1.16-lightgrey.svg)

# What is GoFarm
**GoFarm**  is an Application Development Framework for especially a Backend Developer with Golang. Our goal is to develop easier, standardized, and faster than you code from scratch. Simplify the interface and logic structure for an easy-to-maintain project.

**GoFarm** lets you create a better project, decreasing the learning curve and minimizing the amount of code needed. Let's contribute and make it better :)
## Installation
We need some components to support Build & Deploy working well. Need to install:

For GoFarm Assistant / CLI:
```
go install github.com/0to1a/gofarm-assistant@latest
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
- **Rapid Development:** Analyzed behavior & Designed for Backend Engineer. Make better to maintain and rapid development. Reuse your modules and connect the dot quickly.
- **Performance First:** Framework with a small footprint, make it not costly for production-grade.
- **Dual-Architecture:** Using Microservices concept, but can deploy as single apps / Monolith. Just move the module to a new project, and the app will separate & become a microservice.
- **Structured:** Single pattern of writing code (like MVC), makes other developers read & maintain easily. Minimizing the learning curve is our focus.
- **Clean Deployment:** Focus on a clean structure to make the sure better deployment.

## Requirement
Golang version 1.16+ and Go-Mod supported.

## Feature on Next Development
- [ ] Better logging system
- [ ] GoFarm Assistant, like PHP Artisan-Laravel / Lumen
- [x] Continuable framework scaling & dependency module
- [x] Unit test, framework coverage
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
- github.com/golang-jwt/jwt
- github.com/go-co-op/gocron
- github.com/go-redis/redis
- github.com/volatiletech/sqlboiler
- github.com/golang-migrate/migrate
