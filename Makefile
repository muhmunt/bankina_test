migration_up: 
	migrate -path database/migration/ -database "mysql://root:@tcp(127.0.0.1:3306)/bankina_test" -verbose up

migration_down: 
	migrate -path database/migration/ -database "mysql://root:@tcp(127.0.0.1:3306)/bankina_test" -verbose migration_down

migration_fix: 
	migrate -path database/migration/ -database "mysql://root:@tcp(127.0.0.1:3306)/bankina_test" force VERSION

run:
	go run main.go