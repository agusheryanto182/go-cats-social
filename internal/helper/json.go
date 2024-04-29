package helper

import (
	"encoding/json"
	"net/http"

	"github.com/agusheryanto182/go-social-media/internal/model/web"
)

func ReadFromRequestBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(result); err != nil {
		return err
	}
	return nil
}

func WriteResponse(w http.ResponseWriter, response *web.WebResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)

	encoder := json.NewEncoder(w)
	return encoder.Encode(response)
}
