package typicode

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type TypicodeService struct {
	client *http.Client
}

func NewTypicodeService(client *http.Client) *TypicodeService {
	return &TypicodeService{
		client: client,
	}
}

func (ts *TypicodeService) GetUsers() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/users", nil)
	if err != nil {
		return nil, err
	}
	resp, err := ts.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}
