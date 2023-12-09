package domain

import "errors"

var ErrNotEnoughSpace = errors.New("not enough space")
var ErrNotEnoughStorages = errors.New("not enough storages")
var ErrFileNotFound = errors.New("file not found")
var ErrFileNotCompletelyUploaded = errors.New("file is not completely uploaded yet")
