# Golang + Gin + PostgreSQL

## initiate
```sh
go mod init test_crud
```

## libs
```sh
go install github.com/gin-gonic/gin@latest
go get github.com/gin-contrib/cors
go get github.com/lib/pq
go get github.com/jinzhu/gorm
go get github.com/olahol/melody
```

## run
```sh
# normal
go run .
# detach
go run . &
```

## create struct from ddl
https://qiita.com/neras_1215/items/f5b6e29c9fb870f1b4e3#go

# [OverView] package structure

## package:model
- db entity struct
- api request struct
- api response struct

## package:service
api implementation

## package:sql
query, ddl

# [OverView]Design
## IFs
- usr create/read/update/delete
- category create/read/update/delete
- tag create/read/update/delete
- contents create/read/update/delete
- chatroom create/read/update/delete
- chatmsg create+broadcast/read
- wheather get

## tables
detail in ddl.sql
![ER](./ER.png)

