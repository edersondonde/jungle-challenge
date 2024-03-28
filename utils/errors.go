package utils

import "errors"

var (
	//404 errors
	ErrClientNotFound = errors.New("client not found")

	//400 errors
	ErrStartDateWrongFormat   = errors.New("wrong format of startBirthday, expecting yyyy-mm-dd")
	ErrEndDateWrongFormat     = errors.New("wrong format of endBirthday, expecting yyyy-mm-dd")
	ErrEndDateBeforeStartDate = errors.New("endBirthday should be after startBirthday")
)
