package helpers

import (
	"errors"
	"fmt"
)

type ValidationError error

var (
	ErrorEmptyUsername = ValidationError(errors.New("username field mustn't be empty"))
	ErrorLastname = ValidationError(errors.New("lastname error"))
	ErrorPassword = ValidationError(errors.New("password error"))
	ErrorEmail = ValidationError(errors.New("email format invalid"))
	ErrorPasswordEncryption = ValidationError(errors.New("we couldnt cypher your password"))
	ErrorNotValidData = ValidationError(errors.New("data not valid"))
	ErrorUserRegistred = ValidationError(errors.New("this email has been registred already"))
	NotFoun = ValidationError(errors.New("this user couldn't be found"))
	ErrorPostByUserId = ValidationError(errors.New("this user couldn't be found"))
	ErrorPostData = ValidationError(errors.New("post data not valid, check out."))
	ErrorInvalidPost = ValidationError(errors.New("post  not valid"))
	InterServerError = ValidationError(errors.New("unexpected error"))
	InvalidToken = ValidationError(errors.New("Invalid token"))
)

func UnableToParseDataToStruct(args string) error  {
	return ValidationError(errors.New(fmt.Sprintf("Unable to parse data. Details: %v", args)))
}

func NotFound(context string) error  {
	return ValidationError(errors.New(fmt.Sprintf("Not found: %v", context)))

}

