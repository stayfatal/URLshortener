package errors

import "errors"

var UnexpectedSignMethod = errors.New("Unexpected signing method")

var InvalidToken = errors.New("Invalid token")

var NotUniqueLink = errors.New("This link was already taken")
