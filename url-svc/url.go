package url_svc

import (
	"context"
)

type Params struct {
	Url string `json:"url"`
}

type Response struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

// Create service and repository instances
var urlRepo URLRepository
var urlService *URLService
var idGenerator IDGenerator

// init initializes the repository and service
func init() {
	urlRepo = NewPostgresURLRepository(db_postgres)
	idGenerator = NewRandomIDGenerator()
	urlService = NewURLService(urlRepo, idGenerator)
}

// encore:api public method=POST path=/url
func SaveUrlToDB(ctx context.Context, p *Params) (*Response, error) {
	// Just call the service - all business logic is there
	return urlService.CreateShortURL(ctx, p.Url)
}

// encore:api public method=GET path=/url/:id
func GetUrlFromDB(ctx context.Context, id string) (*Response, error) {
	// Just call the service - all business logic is there
	return urlService.GetOriginalURL(ctx, id)
}
