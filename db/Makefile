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