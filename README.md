# observation-backend
定点観測システムのバックエンド。
## 構成
```
observation-backend
 ├──docker
 |   └──go
 |      └──Dockerfile
 ├──docs
 |   ├─migration
 |   |  ├──000001_create_example.up.sql
 |   |  └──000001_create_example.down.sql
 |   └──swagger
 |      ├──docs.go
 |      ├──swagger.json
 |      └──swagger.yaml
 ├──domain(Entities)
 |   └──exampleDomain.go
 ├──infrastructure
 |   └──sqlhandler.go
 ├──lib
 ├──log
 |   ├──api_debug.log
 |   └──log.go
 ├──api
 |   ├──controller
 |   |   ├──context.go
 |   |   └──exampleControlle.go
 |   ├──database
 |   |   ├──sqlhandler.go
 |   |   └──exampleRepository.go
 |   ├──response
 |   |   ├──errorResponse.go
 |   |   └──exampleResponse.go
 |   └──router
 |       └──outer.go
 ├──usecase
 |   ├──exampleInteractor.go
 |   └──exampleRepository.go
 ├──.air.toml
 ├──.env.example
 ├──.gitignore
 ├──docker-compose.yaml
 ├──go.mod
 ├──go.sum
 └──README.md
```
## 環境構築
1.コンテナを起動
```
docker compose up -d --build
```
2.コンテナに入る
```
docker container exec -it observation-backend-app-1 bash
```
3.マイグレーションの作成(例：usersテーブル)
```
migrate create -ext sql -dir docs/migration -seq create_users
```
4.マイグレーションの実行
```
migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/observation-backend?multiStatements=true" -path=docs/migration up
```
5.Swaggerのビルド
```
swag init --output docs/swagger
```
