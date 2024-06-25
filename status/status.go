package status

import (
	"errors"
	"strings"
)

var _ error = &Status{}

var codeStatusMap = map[int32]*Status{}

type Status struct {
	code int32
	name string
	msg  string
}

func NewStatus(
	code int32,
	name string,
	msg string,
) *Status {
	name = strings.TrimSpace(name)
	status := &Status{
		code: code,
		msg:  msg,
		name: name,
	}
	codeStatusMap[code] = status
	return status
}

func (s *Status) Code() int32 {
	return s.code
}

func (s *Status) Is(err error) bool {
	st, ok := FromError(err)
	if !ok {
		return false
	}
	return s.Equal(st)
}

func (s *Status) As(target interface{}) bool {
	if _, ok := target.(*Status); ok {
		return true
	}
	return false
}

func (s *Status) Error() string {
	return s.msg
}

func (s *Status) Equal(t *Status) bool {
	if t == nil {
		return s == nil
	}
	if s == nil {
		return t == nil
	}
	return s.Code() == t.Code()
}

func Code2Status(code int32) (*Status, bool) {
	statusCode := codeStatusMap[code]
	if statusCode != nil {
		return statusCode, true
	}
	return nil, false
}

func FromError(err error) (*Status, bool) {
	if err == nil {
		return nil, false
	}
	var st *Status
	ok := errors.As(err, &st)
	return st, ok
}
