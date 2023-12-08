package domain

import "errors"

var ErrFileAlreadyExists = errors.New("file already exists")
var ErrFileNotFound = errors.New("file not found")
