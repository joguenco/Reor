# Reor
Application for send email with electronic documents and store in the file system and the database

## Requirements
- Go 1.25
- Fiber
- Gorm

## Install dependencies
```
go mod tidy
```
## Formatter
```
gofmt -w ./src
```
## Run
```
go run src/main.go
```
## Hot Reload
```
air
```
## Build
```
go build -o Reor ./src/main.go
```
## Database
```
CREATE ROLE reor WITH LOGIN NOSUPERUSER CREATEDB NOCREATEROLE INHERIT NOREPLICATION CONNECTION LIMIT -1 PASSWORD 'r'
```
```
pgcli -U reor -d postgres -h localhost -W
```
```
create database reor
```