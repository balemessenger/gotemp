package pkg

import "errors"

type CommonsErrorDef struct {
	InvalidArgument error
}

var CommonsError = CommonsErrorDef{
	InvalidArgument: errors.New("InvalidArgument"),
}
