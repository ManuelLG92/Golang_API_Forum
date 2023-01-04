package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"golang.com/forum/models"
)


func Decode[T any](bytes []byte, message string) (*T, error) {
    out := new(T)
    if err := json.Unmarshal(bytes, out); err != nil {
        return nil, models.UnableToParseDataToStruct(message)
    }
    return out, nil
}

func DecodeBody[T any](body io.ReadCloser, message string) (*T, error) {
    reqBody, err := ioutil.ReadAll(body)
	 if err != nil {
		return nil, err
	}
    return Decode[T](reqBody, message)
}

func Marshal(bytes []byte) ([]byte, error) {
    value,err := json.Marshal(bytes); 
	if err != nil {
        return nil, err
    }
    return value, nil
}