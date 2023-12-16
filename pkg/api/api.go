package api

import "fmt"

type API struct {
	// fields
}

func (a *API) SendRequest() error {
	fmt.Println("Sending API request")
	return nil
}
