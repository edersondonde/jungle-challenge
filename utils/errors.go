package utils

import "errors"

var (
	//404 errors
	ErrClientNotFound = errors.New("client not found")

	//400 errors
	ErrStartDateWrongFormat = errors.New("wrong format of startBirthDate, expecting yyyy-mm-dd")
	ErrEndDateWrongFormat = errors.New("wrong format of endBirthDate, expecting yyyy-mm-dd")
	ErrEndDateBeforeStartDate = errors.New("endBirthDate should be after startBirthDate")

)
