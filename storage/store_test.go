package storage

import "testing"

func TestSaveAndGetURL(t *testing.T) {
	original := "https://www.google.com"

	// Test Saving
	id := SaveURL(original)
	if id == "" {
		t.Fatal("Expected an ID, got an empty string")
	}

	// Test Retrieving
	retrieved, exists := GetURL(id)
	if !exists {
		t.Errorf("Expected URL to exist for ID %s", id)
	}
	if retrieved != original {
		t.Errorf("Expected %s, got %s", original, retrieved)
	}
}

func TestGetNonExistentURL(t *testing.T) {
	_, exists := GetURL("invalid_id")
	if exists {
		t.Error("Expected exists to be false for invalid ID")
	}
}
