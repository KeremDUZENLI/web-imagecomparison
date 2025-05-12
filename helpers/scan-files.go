package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	suffixes := []string{"html", "js", ".go"}
	ignore := []string{"image-json.js", "project-structure.py", "scan-files.go"}

	if err := walkTree(".", suffixes, ignore); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func walkTree(dir string, suffixes, ignore []string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	sort.Slice(entries, func(i, j int) bool {
		iIsDir := entries[i].IsDir()
		jIsDir := entries[j].IsDir()
		if iIsDir != jIsDir {
			return iIsDir
		}
		return entries[i].Name() < entries[j].Name()
	})

	for _, e := range entries {
		path := filepath.Join(dir, e.Name())

		if e.IsDir() {
			if err := walkTree(path, suffixes, ignore); err != nil {
				return err
			}
			continue
		}

		if shouldIgnore(e.Name(), ignore) || !hasSuffix(path, suffixes) {
			continue
		}

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read %s: %v\n", path, err)
			continue
		}

		relPath, err := filepath.Rel(".", path)
		if err != nil {
			relPath = path
		}

		fmt.Printf("\n%s=\n%s\n", relPath, content)
		fmt.Println(strings.Repeat("-", 50))
	}

	return nil
}

func shouldIgnore(name string, ignore []string) bool {
	for _, ig := range ignore {
		if name == ig {
			return true
		}
	}
	return false
}

func hasSuffix(path string, suffixes []string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(path, s) {
			return true
		}
	}
	return false
}
