package models

import "errors"

var EntityNotFoundError error = errors.New("entityt was not found")
var SetCachedContentFailedError error = errors.New("setting content cache failed")
