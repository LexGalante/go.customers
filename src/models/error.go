package models

import "fmt"

//Error -> Represent an error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//MakeNotFoundError -> when entity not found
func MakeNotFoundError(resourceName string) Error {
	e := Error{
		Code:    "NOT_FOUND_ERROR",
		Message: fmt.Sprintf("The resource %s not found", resourceName),
	}

	return e
}

//MakeUnauthorizedError -> when entity not found
func MakeUnauthorizedError() Error {
	e := Error{
		Code:    "UNAUTHORIZED",
		Message: "unrecognized user/password",
	}

	return e
}

//MakeForbiddenError -> when entity not found
func MakeForbiddenError() Error {
	e := Error{
		Code:    "FORBIDDEN",
		Message: "you cannot perform this action",
	}

	return e
}

//MakeInvalidParameterError -> when url parameter is invalid
func MakeInvalidParameterError(parameterName string) Error {
	e := Error{
		Code:    "INVALID_PARAMETER_ERROR",
		Message: fmt.Sprintf("The parameter %s is invalid", parameterName),
	}

	return e
}

//MakeInvalidJSONBodyError -> when body cannot parse to struct
func MakeInvalidJSONBodyError() Error {
	e := Error{
		Code:    "INVALID_JSON_BODY",
		Message: "The body cannot be deserialize",
	}

	return e
}

//MakeUnexpectedError -> when unexpected error ocurred
func MakeUnexpectedError() Error {
	e := Error{
		Code:    "UNEXPECTED_ERROR",
		Message: "An unexpected error ocurred",
	}

	return e
}

//MakeUnexpectedWithBodyError -> when unexpected error ocurred
func MakeUnexpectedWithBodyError(details string) Error {
	e := Error{
		Code:    "UNEXPECTED_ERROR",
		Message: fmt.Sprintf("An unexpected error ocurred: %s", details),
	}

	return e
}
