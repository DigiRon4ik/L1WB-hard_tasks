package main

import (
	"errors"
	"testing"
	"time"
)

// MockTimeFetcher - mock for the TimeFetcher interface.
type MockTimeFetcher struct {
	MockTime time.Time
	MockErr  error
}

// FetchTime - implementation of a mock for the TimeFetcher interface.
func (m MockTimeFetcher) FetchTime(server string) (time.Time, error) {
	return m.MockTime, m.MockErr
}

// TestGetExactTimeSuccess - successful test using mock.
func TestGetExactTimeSuccess(t *testing.T) {
	// Setting up a mock: successful time return
	mockFetcher := MockTimeFetcher{
		MockTime: time.Date(2025, 1, 11, 12, 34, 56, 0, time.UTC),
		MockErr:  nil,
	}

	// Testing the function
	timeStr, err := GetExactTime(mockFetcher, "mock.server")
	if err != nil {
		t.Fatalf("Success was expected, but an error occurred: %v", err)
	}

	// Checking that the time is returned in the correct format
	expected := "2025-01-11 12:34:56 UTC"
	if timeStr != expected {
		t.Errorf("Expected %q, but received %q", expected, timeStr)
	}
	t.Logf("Successfully obtained exact time: %s", timeStr)
}

// TestGetExactTimeError - error handling test.
func TestGetExactTimeError(t *testing.T) {
	// Setting up the mock: returning an error
	mockFetcher := MockTimeFetcher{
		MockErr: errors.New("test error"),
	}

	// Testing the function
	_, err := GetExactTime(mockFetcher, "mock.server")
	if err == nil {
		t.Fatalf("An error was expected, but it did not happen")
	}

	// Checking that the error is processed correctly
	expectedErr := "test error"
	if err.Error() != expectedErr {
		t.Errorf("Expected error %q, but received %q", expectedErr, err.Error())
	}
	t.Logf("Error processed successfully: %v", err)
}
