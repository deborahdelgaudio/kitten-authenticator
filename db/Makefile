LOCAL_DATABASE := "mysql://admin_kitten:kitten_pwd@tcp(localhost:3306)/kitten_house"
DATABASE := ${DATABASE}

hello: 
	echo "Kitten Authenticator 1.0"

migrations: 
	echo "Migrating database"
	migrate -path ./migrations/ -database $(DATABASE) up

migrations_revert:
	echo "Reverting migrations"
	migrate -path ./migrations/ -database $(DATABASE) down

.PHONY: migrations migrations_revert