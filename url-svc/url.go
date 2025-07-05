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

// encore:api public method=POST path=/url
func SaveUrlToDB(ctx context.Context, p *Params) (*Response, error) {

	// generate id
	id, err := GenerateRandomBase64String()
	if err != nil {
		return nil, err
	}

	// save to db
	_, err = db.Exec(ctx, "INSERT INTO url (id, url) VALUES ($1, $2)", id, p.Url)
	if err != nil {
		return nil, err
	}

	// return response
	return &Response{Id: id, Url: p.Url}, nil
}

// encore:api public method=GET path=/url/:id
func GetUrlFromDB(ctx context.Context, id string) (*Response, error) {

	var url string
	// get from db
	err := db.QueryRow(ctx, "SELECT url FROM url WHERE id = $1", id).Scan(&url)
	if err != nil {
		return nil, err
	}

	// return response
	return &Response{Id: id, Url: url}, nil
}
