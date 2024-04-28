
db := "deities.sqlite"

# Run wp24-deities.
run:
    @go run main.go

create-database:
    rm -f {{db}}.sqlite
    sqlite3 {{db}} ".read extras/deities.sql"
