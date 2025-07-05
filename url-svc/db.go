package url_svc

import (
	"encore.dev/storage/sqldb"
)

var db_postgres = sqldb.NewDatabase("url", sqldb.DatabaseConfig{
    Migrations: "./migrations_postgres",
})




