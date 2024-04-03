package iopsep

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilPipeReader = erorr.Error("iopsep: nil pipe-reader")
	errNilPipeWriter = erorr.Error("iopsep: nil pipe-writer")
	errNilReader     = erorr.Error("iopsep: nil reader")
	errNilReceiver   = erorr.Error("iopsep: nil receiver")
	errNilWriter     = erorr.Error("iopsep: nil writer")
)

const (
	errNilReturnEOFFunction   = erorr.Error("iopsep: nil return-eof function")
	errNilReturnErrorFunction = erorr.Error("iopsep: nil return-error function")
	errNilWriteRuneFunction   = erorr.Error("iopsep: nil write-rune function")
)
