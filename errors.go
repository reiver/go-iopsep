package iopsep

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilPipeReader = erorr.Error("iop: nil pipe-reader")
	errNilPipeWriter = erorr.Error("iop: nil pipe-writer")
	errNilReader     = erorr.Error("iop: nil reader")
	errNilReceiver   = erorr.Error("iop: nil receiver")
	errNilWriter     = erorr.Error("iop: nil writer")
)

const (
	errNilReturnEOFFunction   = erorr.Error("iop: nil return-eof function")
	errNilReturnErrorFunction = erorr.Error("iop: nil return-error function")
	errNilWriteRuneFunction   = erorr.Error("iop: nil write-rune function")
)
