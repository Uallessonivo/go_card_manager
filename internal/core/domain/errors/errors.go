package errors

import "errors"

var (
	NotFound         = errors.New("record not found")
	OwnerNotFound    = errors.New("to create a new card the owner has to be registered")
	NoDataFound      = errors.New("we didn't find any record")
	AlreadyExists    = errors.New("this record already exists")
	InvalidEmail     = errors.New("email is invalid")
	InvalidPassword  = errors.New("password must be alphanumeric and must consists of at least 6 characters and not more than 15 characters")
	InvalidParams    = errors.New("please make sure all required fields are filled out and try again")
	InvalidFields    = errors.New("one or more fields in the data you sent are invalids")
	MaxNumberOfCards = errors.New("each employee must have a maximum of 2 cards")
	FileExtension    = errors.New("We only support '.xlsx' files")
	InvalidLogin     = errors.New("Email or Password invalid")
)
