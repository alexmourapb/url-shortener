package domain

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// nolint
func ExampleErrTypeIdentification() {
	var MyDomainErr = errors.New("my sentinel domain error")
	var errFn = func() error {
		operation := "errFn"
		// some processing, then an error occurs and we return it
		return Error(operation, MyDomainErr)
	}
	err := errFn()
	if err != nil {
		if errors.Is(err, MyDomainErr) {
			fmt.Println("we caught a domain err!")
		} else {
			fmt.Println("unmapped domain err... probably something edgy happened")
		}
	}
	// Output: we caught a domain err!
}

type testErrorMessage struct {
	fn          func() error
	expectedErr string
}

func TestApplicationError_ErrorMessage(t *testing.T) {
	type fields struct {
		errFunc func() error
	}
	tests := []struct {
		name          string
		fields        fields
		wantErrString string
	}{
		{
			name: "empty error return nil",
			fields: fields{
				errFunc: func() error {
					return Error("errNil", nil)
				},
			},
			wantErrString: "",
		},
		{
			name: "ordinary error prints err message without operation breadcrumb",
			fields: fields{
				errFunc: func() error {
					return errors.New("something bad happened")
				},
			},
			wantErrString: "something bad happened",
		},
		{
			name: "application error prints err message with full operation breadcrumb",
			fields: fields{
				errFunc: errBreadcrumbFunc().fn,
			},
			wantErrString: errBreadcrumbFunc().expectedErr,
		},
		{
			name: "application error prints err message with full breadcrumb and extra params",
			fields: fields{
				errFunc: extraParamsErrFunc().fn,
			},
			wantErrString: extraParamsErrFunc().expectedErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := tt.fields.errFunc()
			if tt.wantErrString != "" {
				assert.Equal(t, tt.wantErrString, errMsg.Error())
			} else {
				assert.Empty(t, errMsg)
			}
		})
	}
}

func TestApplicationError_Unwrap(t *testing.T) {
	var MyDummyErr = errors.New("oh noes")
	tests := []struct {
		name            string
		errFunc         func() error
		ExpectedErrType error
	}{
		{
			name: "succesfully unwraps domain error",
			errFunc: func() error {
				return MyDummyErr
			},
			ExpectedErrType: MyDummyErr,
		},
		{
			name: "succesfully unwraps fmt.Errorf wrapped domain error",
			errFunc: func() error {
				return fmt.Errorf("shit happened and we must know the type: %w", MyDummyErr)
			},
			ExpectedErrType: MyDummyErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, errors.Is(tt.errFunc(), tt.ExpectedErrType))
		})
	}
}

func errBreadcrumbFunc() testErrorMessage {
	return testErrorMessage{
		fn: func() error {
			op := "errBreadcrumbFunc"
			a := func() error {
				op := "a"
				return Error(op, errors.New("something bad happened"))
			}
			b := func() error {
				op := "b"
				if err := a(); err != nil {
					return Error(op, err)
				}
				return nil
			}
			err := b()
			return Error(op, err)
		},
		expectedErr: "errBreadcrumbFunc()->b()->a(): something bad happened",
	}
}

func extraParamsErrFunc() testErrorMessage {
	return testErrorMessage{
		fn: func() error {
			op := "errExtraParams"
			a := func(int) error {
				op := "a"
				err := errors.New("something bad happened")
				return Error(op, err)
			}
			b := func(string) error {
				op := "b"
				value := 100405
				err := a(value)
				return Error(op, err, Params{"value": value})
			}
			c := func(bool) error {
				op := "c"
				err := b("b23-3ef-6da")
				if err != nil {
					return Error(op, err, Params{"id": "b23-3ef-6da"})
				}
				return nil
			}
			err := c(true)
			return Error(op, err, Params{"bool_param": true})
		},
		expectedErr: `errExtraParams()->c("bool_param":true)->b("id":"b23-3ef-6da")->a("value":100405): something bad happened`,
	}
}
