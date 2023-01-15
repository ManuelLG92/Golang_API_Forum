package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
)


func Decode[T any](bytes []byte, message string) (*T, error) {
    out := new(T)
    if err := json.Unmarshal(bytes, out); err != nil {
        return nil, UnableToParseDataToStruct(message)
    }
    validate := validator.New()
    if err := validate.Struct(out); err != nil {
        return nil, err
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