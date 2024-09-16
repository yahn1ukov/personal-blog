package service

import "errors"

var (
	ErrTitleRequired   = errors.New("title is required")
	ErrContentRequired = errors.New("content is required")
	ErrNoFieldsUpdate  = errors.New("no fields to update")
)
