createmigrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3307)/game_store?multiStatements=true" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://root:root@tcp(localhost:3307)/game_store?multiStatements=true" -verbose down

.PHONY: createmigrate migrateup migratedown