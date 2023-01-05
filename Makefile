createmigrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3307)/game_store?multiStatements=true" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3307)/game_store?multiStatements=true" -verbose down

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9000 -r repl

.PHONY: createmigrate migrateup migratedown proto evans