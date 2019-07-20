package errors

import "errors"

type AppleErrorDef struct {
	ApnsKeyNotValid  error
	ApplePushDisable error
}

var AppleError = AppleErrorDef{
	ApnsKeyNotValid:  errors.New("ApnsKeyNotValid"),
	ApplePushDisable: errors.New("ApplePushDisable"),
}
