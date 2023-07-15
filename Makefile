DB_URL=postgresql://root:secret@localhost:5432/blog_post?sslmode=disable

imagerm:
	docker rm -vf $(docker ps -aq)

network:
	docker network create blog-network

postgres:
	docker run --name postgres-blog-post --network blog-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine


createdb:
	docker exec -it postgres-blog-post createdb --username=root --owner=root blog_post

dropdb:
	docker exec -it postgres dropdb blog_post

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ruffiano/blog-post/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/ruffiano/blog-post/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=blog_post \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: imagerm network postgres createdb dropdb migrateup migratedown db_docs db_schema sqlc test server mock proto evans redis