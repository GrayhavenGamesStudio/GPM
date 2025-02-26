package packagefile

import (
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

// TestCreatePackageFile checks if packagefile.gpm is created correctly
func TestCreatePackageFile(t *testing.T) {
	// Use a temporary file for testing
	testFile := "test_packagefile.gpm"

	// Ensure the file doesn't exist before testing
	os.Remove(testFile)

	// Run the function to create the package file
	err := CreatePackageFile(testFile)
	if err != nil {
		t.Fatalf("CreatePackageFile failed: %v", err)
	}

	// Verify the file was created
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to be created, but it was not", testFile)
	}

	// Read the file to check contents
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read created file: %v", err)
	}

	// Unmarshal YAML into struct
	var pkg PackageFile
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		t.Fatalf("Invalid YAML format in packagefile.gpm: %v", err)
	}

	// Validate expected values
	expectedName := "example-package"
	if pkg.Package.Name != expectedName {
		t.Errorf("Expected package name %q, got %q", expectedName, pkg.Package.Name)
	}

	expectedVersion := "0.1.0"
	if pkg.Package.Version != expectedVersion {
		t.Errorf("Expected package version %q, got %q", expectedVersion, pkg.Package.Version)
	}

	// Ensure function does not overwrite an existing file
	err = CreatePackageFile(testFile)
	if err == nil {
		t.Errorf("Expected an error when creating an existing file, but got nil")
	}

	// Cleanup test file
	os.Remove(testFile)
}

// TestEnsurePackagesDir verifies that EnsurePackagesDir creates the packages/ directory
func TestEnsurePackagesDir(t *testing.T) {
	// Ensure the directory is removed before testing
	_ = os.RemoveAll("packages")

	// Step 1: First run should create the directory
	err := EnsurePackagesDir()
	if err != nil {
		t.Fatalf("EnsurePackagesDir failed: %v", err)
	}

	// Step 2: Verify the directory now exists
	if _, err := os.Stat("packages"); os.IsNotExist(err) {
		t.Fatalf("Expected packages/ directory to exist, but it does not")
	}

	// Step 3: Second run should not return an error (idempotency test)
	err = EnsurePackagesDir()
	if err != nil {
		t.Errorf("EnsurePackagesDir should not return an error if the directory already exists, but got: %v", err)
	}

	// Cleanup after test
	_ = os.RemoveAll("packages")
}
