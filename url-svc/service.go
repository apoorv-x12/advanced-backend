package url_svc

import (
	"context"
)

// URLService handles URL shortening business logic
type URLService struct {
	repo URLRepository
	id IDGenerator
}

// NewURLService creates a new URL service
func NewURLService(repo URLRepository, id IDGenerator) *URLService {
	return &URLService{
		repo: repo,
		id: id,
	}
}

// CreateShortURL creates a shortened URL
func (s *URLService) CreateShortURL(ctx context.Context, originalURL string) (*Response, error) {
	// Generate unique ID
	id, err := s.id.Generate()
	if err != nil {
		return nil, err
	}

	// Save to repository
	err = s.repo.Save(ctx, id, originalURL)
	if err != nil {
		return nil, err
	}

	return &Response{Id: id, Url: originalURL}, nil
}

// GetOriginalURL retrieves the original URL by short ID
func (s *URLService) GetOriginalURL(ctx context.Context, id string) (*Response, error) {
	url, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Response{Id: id, Url: url}, nil
}
