package url_svc

import (
	"context"
	"testing"
)

func TestSaveAndGetUrlToDB(t *testing.T) {
	ctx := context.Background()
	p := &Params{Url: "https://www.google.com"}
	res, err := SaveUrlToDB(ctx, p)
	if err != nil {
		t.Fatalf("SaveUrlToDB failed: %v", err)
	} else if res.Url != p.Url {
		t.Fatalf("SaveUrlToDB failed: %v", err)
	} else{
		t.Logf("SaveUrlToDB passed: %v", res)
	}

	res, err = GetUrlFromDB(ctx, res.Id)
	if err != nil {
		t.Fatalf("GetUrlFromDB failed: %v", err)
	} else if res.Url != p.Url {
		t.Fatalf("GetUrlFromDB failed: %v", err)
	} else{
		t.Logf("GetUrlFromDB passed: %v", res)
	}

}