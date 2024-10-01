package errors

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func New(msg string) error {
	return errors.New(msg)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// Cause 取得錯誤原因
func Cause(err error) error {
	return errors.Cause(err)
}

// CauseStackTrace 取得錯誤原因發生位置
func CauseStackTrace(err error) string {
	stack := fmt.Sprintf("%+v", err)
	split := strings.Split(stack, "\n")

	stackTrace := make([]string, 0)
	for _, row := range split {
		if strings.Contains(row, "\t") {
			stackTrace = append(stackTrace, strings.ReplaceAll(row, "\t", ""))
		}
	}
	if len(stackTrace) > 6 {
		stackTrace = stackTrace[1:6]
	} else {
		stackTrace = stackTrace[1:]
	}

	return strings.Join(stackTrace, `::`)
}
