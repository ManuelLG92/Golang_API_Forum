package models

import "errors"

type ValidationError error

var (
	ErrorEmptyUsername = ValidationError(errors.New("username field mustn't be empty"))
	//ErrorShortUsername = ValidationError(errors.New("username too short"))
	//ErrorLargeUsername = ValidationError(errors.New("username too long"))

	ErrorLastname = ValidationError(errors.New("lastname error"))

	ErrorPassword = ValidationError(errors.New("pAssword error"))

	ErrorEmail = ValidationError(errors.New("email format invalid"))

	ErrorPasswordEncryption = ValidationError(errors.New("we couldnt cypher your password"))

	//ErrorNotValidUser = ValidationError(errors.New("user not valid"))
	ErrorNotValidData = ValidationError(errors.New("data not valid"))

	//ErrorInvalidLogin = ValidationError(errors.New("email or password not valid"))

	ErrorUserRegistred = ValidationError(errors.New("this email has been registred already"))

	ErrorPostByUserId = ValidationError(errors.New("this user couldn't be found"))
	ErrorPostData = ValidationError(errors.New("post data not valid, check out."))
	//ErrorPostContent = ValidationError(errors.New("post content not valid"))
	ErrorInvalidPost = ValidationError(errors.New("post  not valid"))

	//ErrorSessionAlreadyRegistred= ValidationError(errors.New("post  not valid"))
)

