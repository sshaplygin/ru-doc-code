package bik

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	ru_doc_code "github.com/mrfoe7/go-codes-validator"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("invalid bik length", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "1234567888776",
				Error:   ru_doc_code.ErrInvalidBIKLength,
				IsValid: false,
			},
			{
				Code:    "044525",
				Error:   ru_doc_code.ErrInvalidBIKLength,
				IsValid: false,
			},
			{
				Code:    "044525225",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "044525012",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})

	t.Run("invalid bik value", func(t *testing.T) {
		testCases := []ru_doc_code.TestCodeCase{
			{
				Code:    "0445?5226",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "054525225",
				Error:   ru_doc_code.ErrInvalidBIKCountryCode,
				IsValid: false,
			},
			{
				Code:    "104525225",
				Error:   ru_doc_code.ErrInvalidBIKCountryCode,
				IsValid: false,
			},
			{
				Code:    "044#55#25",
				Error:   ru_doc_code.ErrInvalidValue,
				IsValid: false,
			},
			{
				Code:    "044525225",
				Error:   nil,
				IsValid: true,
			},
			{
				Code:    "044525012",
				Error:   nil,
				IsValid: true,
			},
		}
		for _, test := range testCases {
			isValid, err := Validate(test.Code)
			assert.Equal(t, test.IsValid, isValid, test.Code, test.IsValid)
			assert.Equal(t, true, errors.Is(test.Error, err), test.Code)
		}
	})
}
