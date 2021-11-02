package environment

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment_IsValid(t *testing.T) {

	var tt = []struct {
		Name    string
		Value   string
		Valid   bool
		Message string
	}{
		{
			Name:    "Validating 'test' environment",
			Value:   "test",
			Valid:   true,
			Message: "'test' environment should be valid!",
		},
		{
			Name:    "Validating 'developer' environment",
			Value:   "developer",
			Valid:   true,
			Message: "'developer' environment should be valid!",
		},
		{
			Name:    "Validating 'sandbox",
			Value:   "sandbox",
			Valid:   true,
			Message: "'sandbox' environment should be valid!",
		},
		{
			Name:    "Validating 'production",
			Value:   "production",
			Valid:   true,
			Message: "'production' environment should be valid!",
		},
		{
			Name:    "Validating empty environment",
			Value:   "",
			Valid:   false,
			Message: "empty environment should not be valid!",
		},
		{
			Name:    "Validating incorrect environment",
			Value:   "somethingelse",
			Valid:   false,
			Message: "incorrect environment should not be valid!",
		},
	}

	for _, tc := range tt {
		testName := tc.Name
		t.Run(testName, func(t *testing.T) {

			if tc.Valid == true {
				assert.True(t, IsValid(tc.Value), tc.Message)
			} else {
				assert.False(t, IsValid(tc.Value), tc.Message)
			}

		})
	}

}

func TestEnvironment_AsStringAsString(t *testing.T) {

	var tt = []struct {
		Name        string
		Environment Environment
		Expected    string
	}{

		{
			Name:        "Comparing 'test' environment AsString",
			Environment: Test,
			Expected:    "test",
		},
		{
			Name:        "Comparing 'developer' environment AsString",
			Environment: Developer,
			Expected:    "developer",
		},
		{
			Name:        "Comparing 'prod' environment AsString",
			Environment: Production,
			Expected:    "production",
		},
		{
			Name:        "Comparing 'sandbox' environment AsString",
			Environment: Sandbox,
			Expected:    "sandbox",
		},
	}

	for _, tc := range tt {
		testName := tc.Name
		t.Run(testName, func(t *testing.T) {
			env := tc.Environment.String()
			assert.Equal(t, env, tc.Expected, fmt.Sprintf("Expecting receive %s. Received: %s", tc.Expected, env))
		})
	}

}

func TestEnvironment_FromString(t *testing.T) {

	var tt = []struct {
		Name     string
		String   string
		Expected Environment
	}{

		{
			Name:     "Converting 'test' to Environment.Test",
			String:   "test",
			Expected: Test,
		},
		{
			Name:     "Converting 'developer' to Environment.Developer",
			String:   "developer",
			Expected: Developer,
		},
		{
			Name:     "Converting 'prod' to Environment.Production",
			String:   "production",
			Expected: Production,
		},
		{
			Name:     "Converting 'sandbox' to Environment.Sandbox",
			String:   "sandbox",
			Expected: Sandbox,
		},
		{
			Name:     "Converting 'somethingelse' to Environment.Developer",
			String:   "developer",
			Expected: Developer,
		},
	}

	for _, tc := range tt {
		testName := tc.Name
		t.Run(testName, func(t *testing.T) {
			env := FromString(tc.String)
			assert.Equal(t, env, tc.Expected)
		})
	}
}
