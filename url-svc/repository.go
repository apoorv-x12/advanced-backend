package url_svc

import (
	"context"

	"encore.dev/storage/sqldb"
)

// URLRepository defines the interface for URL data operations
type URLRepository interface {
	Save(ctx context.Context, id, url string) error
	GetByID(ctx context.Context, id string) (string, error)
}

// PostgresURLRepository implements URLRepository using PostgreSQL
type PostgresURLRepository struct {
	db *sqldb.Database
}

// NewPostgresURLRepository creates a new PostgreSQL repository
func NewPostgresURLRepository(db *sqldb.Database) *PostgresURLRepository {
	return &PostgresURLRepository{db: db}
}

// Save stores a URL with the given ID
func (r *PostgresURLRepository) Save(ctx context.Context, id, url string) error {
	_, err := r.db.Exec(ctx, "INSERT INTO url (id, url) VALUES ($1, $2)", id, url)
	return err
}

// GetByID retrieves a URL by its ID
func (r *PostgresURLRepository) GetByID(ctx context.Context, id string) (string, error) {
	var url string
	err := r.db.QueryRow(ctx, "SELECT url FROM url WHERE id = $1", id).Scan(&url)
	return url, err
}
