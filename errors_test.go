package ensure

import (
	"github.com/pkg/errors"
	"testing"
)

func TestNoError_GivenNil(t *testing.T) {
	var err error

	NoError(err)
}

func TestNoError_GivenError(t *testing.T) {
	err := errors.New("expected")

	defer func() {
		r := recover()
		expectToRecoverFromPanic(t, r)

		err = r.(error)
		equalError(t, err, "unexpected error: expected")
	}()

	NoError(err)
}

func TestError_GivenNil(t *testing.T) {
	var err error

	defer func() {
		r := recover()
		expectToRecoverFromPanic(t, r)

		err = r.(error)
		equalError(t, err, "expecting error, but none given")
	}()

	Error(err)
}

func TestError_GivenError(t *testing.T) {
	err := errors.New("expected")

	Error(err)
}

func TestErrorWithMessage(t *testing.T) {
	err1 := errors.New("expected")
	err2 := errors.New("expect")
	re := "^expect(?:ed)?$"

	ErrorWithMessage(err1, re)
	ErrorWithMessage(err2, re)
}

func TestErrorWithMessage_DifferentMessage(t *testing.T) {
	err := errors.New("dogs")
	re := "^cats$"

	defer func() {
		r := recover()
		expectToRecoverFromPanic(t, r)

		err = r.(error)
		equalError(t, err, "given error doesn't match given regexp (^cats$): dogs")
	}()

	ErrorWithMessage(err, re)
}

func equalError(t *testing.T, err error, expectedMessage string) {
	actual := err.Error()
	if actual != expectedMessage {
		t.Errorf("expecting error message to be: %s, but was: %s", expectedMessage, actual)
	}
}

func expectToRecoverFromPanic(t *testing.T, r interface{}) {
	if r == nil {
		t.Fatal(errors.New("expected to recover from panic, but didn't"))
	}
}
