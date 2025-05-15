package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var suffixes = []string{"go"}
var ignore = []string{"helpers"}

func main() {
	if err := walkTree(".", suffixes, ignore); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := runPythonScript("helpers/project_structure.py"); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run Python script: %v\n", err)
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
			if shouldIgnore(path, ignore) {
				continue
			}
			if err := walkTree(path, suffixes, ignore); err != nil {
				return err
			}
			continue
		}

		if !hasSuffix(path, suffixes) {
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

func shouldIgnore(path string, ignore []string) bool {
	for _, dir := range ignore {
		if strings.HasPrefix(filepath.ToSlash(path), dir+"/") || filepath.ToSlash(path) == dir {
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

func runPythonScript(script string) error {
	cmd := exec.Command("python", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
