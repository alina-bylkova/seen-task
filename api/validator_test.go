package api

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateId(t *testing.T) {
	testCases := []struct {
		id             int64
		expectedResult error
		caseName       string
	}{
		{
			id:             -1,
			expectedResult: errors.New("Invalid id is provided"),
			caseName:       "Negative id number",
		},
		{
			id:             0,
			expectedResult: errors.New("Invalid id is provided"),
			caseName:       "Id number is zero",
		},
		{
			id:             1,
			expectedResult: nil,
			caseName:       "Valid id",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing %v", tc.caseName), func(t *testing.T) {
			result := validateId(tc.id)
			if result == nil {
				if tc.expectedResult != nil {
					t.Errorf("Expected %v, got %v", tc.expectedResult, result)
				}
			} else if result.Error() != tc.expectedResult.Error() {
				t.Errorf("Expected %v, got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult error
		caseName       string
	}{
		{
			name:           "",
			expectedResult: nil,
			caseName:       "Name is not provided",
		},
		{
			name:           "a",
			expectedResult: errors.New("Invalid name is provided"),
			caseName:       "Name with one letter",
		},
		{
			name:           "Jane",
			expectedResult: nil,
			caseName:       "Valid name",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing %v", tc.caseName), func(t *testing.T) {
			result := validateName(tc.name)
			if result == nil {
				if tc.expectedResult != nil {
					t.Errorf("Expected %v, got %v", tc.expectedResult, result)
				}
			} else if result.Error() != tc.expectedResult.Error() {
				t.Errorf("Expected %v, got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		email          string
		expectedResult error
		caseName       string
	}{
		{
			email:          "",
			expectedResult: nil,
			caseName:       "Email is not provided",
		},
		{
			email:          "bad-email",
			expectedResult: errors.New("Invalid email is provided"),
			caseName:       "Invalid email example 1",
		},
		{
			email:          "test-email.com",
			expectedResult: errors.New("Invalid email is provided"),
			caseName:       "Invalid email example 2",
		},
		{
			email:          "test44@gmail.com",
			expectedResult: nil,
			caseName:       "Valid id",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing %v", tc.caseName), func(t *testing.T) {
			result := validateEmail(tc.email)
			if result == nil {
				if tc.expectedResult != nil {
					t.Errorf("Expected %v, got %v", tc.expectedResult, result)
				}
			} else if result.Error() != tc.expectedResult.Error() {
				t.Errorf("Expected %v, got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	testCases := []struct {
		phone          string
		expectedResult error
		caseName       string
	}{
		{
			phone:          "",
			expectedResult: nil,
			caseName:       "Phone number is not provided",
		},
		{
			phone:          "123",
			expectedResult: errors.New("Invalid phone number is provided"),
			caseName:       "Phone number with less than 8 digits",
		},
		{
			phone:          "1234567890123456",
			expectedResult: errors.New("Invalid phone number is provided"),
			caseName:       "Phone number with more than 15 digits",
		},
		{
			phone:          "abc",
			expectedResult: errors.New("Invalid phone number is provided"),
			caseName:       "Phone number non-digit characters",
		},
		{
			phone:          "4799999999",
			expectedResult: nil,
			caseName:       "Valid phone number",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing %v", tc.caseName), func(t *testing.T) {
			result := validatePhone(tc.phone)
			if result == nil {
				if tc.expectedResult != nil {
					t.Errorf("Expected %v, got %v", tc.expectedResult, result)
				}
			} else if result.Error() != tc.expectedResult.Error() {
				t.Errorf("Expected %v, got %v", tc.expectedResult, result)
			}
		})
	}
}
