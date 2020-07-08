package request

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/types/responsetype"
)

func Get(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&target)
}

func Post(url string, body []byte) (responsetype.Base, error) {
	result := responsetype.Base{}
	r, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return result, err
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&result)
	return result, nil
}
