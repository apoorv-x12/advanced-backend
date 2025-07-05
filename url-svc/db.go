package url_svc

import (
	"encore.dev/storage/sqldb"
)

var db = sqldb.NewDatabase("url", sqldb.DatabaseConfig{
    Migrations: "./migrations_postgres",
})




