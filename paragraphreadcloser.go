package iopsep

import (
	"errors"
	"io"

	"github.com/reiver/go-unicode"
	"sourcecode.social/reiver/go-utf8"
)

type internalParagraphReadCloser struct {
	reader io.Reader
	pipereader *io.PipeReader
	pipewriter *io.PipeWriter
}

func NewParagraphReadCloser(reader io.Reader) io.ReadCloser {

	if nil == reader {
		return nil
	}

	pipereader, pipewriter := io.Pipe()
	if nil == pipereader {
		return nil
	}
	if nil == pipewriter {
		return nil
	}

	paragraphreader := internalParagraphReadCloser{
		reader:reader,
		pipereader:pipereader,
		pipewriter:pipewriter,
	}

	go paragraphreader.pipewrite()

	return &paragraphreader
}

func (receiver *internalParagraphReadCloser) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	var pipewritererror error
	{
		var pipewriter *io.PipeWriter = receiver.pipewriter
		if nil != pipewriter {
			err := pipewriter.Close()
			if nil != err {
				pipewritererror = err
			}
		}
	}

	{
		var pipereader *io.PipeReader = receiver.pipereader
		if nil != pipereader {
			err := pipereader.Close()
			if nil != err {
				return err
			}
		}
	}

	if nil != pipewritererror {
		return pipewritererror
	}

	return nil
}

func (receiver *internalParagraphReadCloser) Read(p []byte) (n int, err error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var reader io.Reader = receiver.reader
	if nil == reader {
		return 0, errNilReader
	}

	var pipereader io.Reader = receiver.pipereader
	if nil == pipereader {
		return 0, errNilPipeReader
	}

	return pipereader.Read(p)
}

func (receiver *internalParagraphReadCloser) pipewrite() {
	if nil == receiver {
		panic(errNilReceiver)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	var reader io.Reader = receiver.reader
	if nil == reader {
		err := pipewriter.CloseWithError(errNilReader)
		if nil != err {
			panic(err)
		}
	}

	pipewrite(receiver.writerune, receiver.returneof, receiver.returnerror, reader)
}

func (receiver *internalParagraphReadCloser) returneof() {
	if nil == receiver {
		panic(errNilWriter)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	{
		err := pipewriter.CloseWithError(io.EOF)
		if nil != err {
			panic(err)
		}
	}
}

func (receiver *internalParagraphReadCloser) returnerror(err error) {
	if nil == receiver {
		panic(errNilWriter)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	{
		err := pipewriter.CloseWithError(err)
		if nil != err {
			panic(err)
		}
	}
}

func (receiver *internalParagraphReadCloser) writerune(r rune) (exit bool) {
	if nil == receiver {
		panic(errNilWriter)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	{
		_, err := utf8.WriteRune(pipewriter, r)
		if errors.Is(err, io.ErrClosedPipe) {
			pipewriter.Close()
			return true
		}
		if nil != err {
			e := pipewriter.CloseWithError(err)
			if nil != e {
				panic(e)
			}
		}
	}

	return false
}

func pipewrite(writerune func(rune)bool, returneof func(), returnerror func(error), reader io.Reader) {

	if nil == writerune {
		panic(errNilWriteRuneFunction)
	}

	if nil == returneof {
		panic(errNilReturnEOFFunction)
	}

	if nil == returnerror {
		panic(errNilReturnErrorFunction)
	}

	var prevcr bool
	var preveol bool
	var prevlf bool

	for {
		r, size, err := utf8.ReadRune(reader)
		if 0 < size {
			exit := writerune(r)
			if exit {
				return
			}

			switch r {
			case unicode.LF:
				switch {
				case prevlf:
					returneof()
					return
				case prevcr && preveol:
					returneof()
					return
				case prevcr:
					prevcr  = false
					preveol = true
					prevlf  = false
				default:
					prevcr  = false
					prevlf  = true
				}
			case unicode.CR:
				switch {
				case prevcr:
					returneof()
					return
				case prevlf && preveol:
					returneof()
					return
				case prevlf:
					prevcr  = false
					preveol = true
					prevlf  = false
				default:
					prevcr  = true
					prevlf  = false
				}
			case unicode.NEL:
				if preveol {
					returneof()
					return
				}
				prevcr  = false
				preveol = true
				prevlf  = false
			case unicode.PS:
				returneof()
				return
			default:
				prevcr  = false
				preveol = false
				prevlf  = false
			}
		}
		if io.EOF == err {
			returneof()
			return
		}
		if nil != err {
			returnerror(err)
			return
		}
	}
}
