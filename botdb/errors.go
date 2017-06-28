package botdb

import "errors"

var ErrIsExist error = errors.New("Is exist")

var ErrFieldsEmpty error = errors.New("one or all fields is empty")

var ErrNotCorrectId error = errors.New("Not Correct Id")
