package types

import (
	"errors"
	"fmt"
	"strings"

	"github.com/creachadair/jrpc2"
)

type (
	Error        string
	ComplexError struct {
		err error
		msg string
	}
	WrappedError []error
)

const (
	ErrClient Error = "API client error"
	ErrLogin  Error = "cannot login to LANBilling API"
)

func (e Error) Error() string {
	return string(e)
}

func NewComplexError(msg string, err error) error {
	return ComplexError{
		err: err,
		msg: msg,
	}
}

func NewComplexErrorf(format string, args ...any) error {
	return ComplexError{
		msg: fmt.Sprintf(format, args...),
	}
}

func NewComplexWithErrorf(err error, format string, args ...any) error {
	return ComplexError{
		err: err,
		msg: fmt.Sprintf(format, args...),
	}
}

func (e ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.err)
}

func (e ComplexError) Unwrap() error {
	return e.err
}

func (e WrappedError) Unwrap() error {
	if len(e) > 1 {
		return e[1:]
	}
	return nil
}

func NewWrappedError(errors ...error) error {
	return WrappedError(errors)
}

func (e WrappedError) Error() string {
	output := make([]string, 0, len(e))
	for _, err := range e {
		output = append(output, err.Error())
	}
	return strings.Join(output, ": ")
}

func ParseJSONRPCError(err error) error {
	var e *jrpc2.Error
	if errors.As(err, &e) {
		return NewComplexErrorf("JSON-RPC error: %d, message: %s", e.Code, e.Message)
	}

	return err
}
