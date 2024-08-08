package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func ReadFromQueryParams(key string, request *http.Request) (string, error) {
	queryValues := request.URL.Query()
	if queryValues.Has(key) {
		return queryValues.Get(key), nil
	}
	return "", errors.New("not found")
}

func JsonEncode(v any)(string){
	jsonBytes, err := json.Marshal(v)
	PanicIfError(err)
	return string(jsonBytes)
}

func JsonDecode(v string, result any){
	err := json.Unmarshal([]byte(v), &result)
	PanicIfError(err)
}