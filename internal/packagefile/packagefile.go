package packagefile

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type PackageFile struct {
	Package      Package  `yaml:"package"`
	Dependencies []string `yaml:"dependencies"` // URLs with versioning - IE "https://github.com/user/physics-lib@1.2.3"
	Exclude      []string `yaml:"exclude"`
}

type Package struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Url     string `yaml:"url"`
}

// ParseDependency extracts URL and version from a dependency string
func ParseDependency(dep string) (string, string, error) {
	parts := strings.Split(dep, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid dependency format: %s", dep)
	}
	return parts[0], parts[1], nil
}

// CreatePackageFile creates a blank packagefile.gpm if it doesn't exist
func CreatePackageFile(name string, url string) error {
	filename := "package_file.gpm"
	defaultPackage := PackageFile{
		Package: Package{
			Name:    name,
			Version: "0.1.0",
			Url:     url,
		},
		Dependencies: []string{},
		Exclude:      []string{".git", "tests/", "docs/"},
	}

	// Convert struct to YAML
	data, err := yaml.Marshal(&defaultPackage)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Check if the file already exists
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("file already exists: %s", filename)
	}

	// Write YAML to file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// EnsurePackagesDir checks and creates the packages/ directory if needed
func EnsurePackagesDir() error {
	if _, err := os.Stat("packages"); os.IsNotExist(err) {
		err := os.Mkdir("packages", 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		fmt.Println("✔ Created packages/ directory")
	}
	return nil
}
