createmigrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3307)/game_store?multiStatements=true" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3307)/game_store?multiStatements=true" -verbose down

proto:
	rm -f pb/*.go
	rm -f docs/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb \
	--grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=api_doc \
    proto/*.proto

evans:
	evans --host localhost --port 9000 -r repl

.PHONY: createmigrate migrateup migratedown proto evans