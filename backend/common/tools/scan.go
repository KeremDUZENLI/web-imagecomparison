package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "../backend"
	ignore := []string{"scan.go"}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		for _, ig := range ignore {
			if strings.Contains(path, ig) {
				return nil
			}
		}
		if strings.HasSuffix(path, ".go") {
			data, err := os.ReadFile(path)
			if err == nil {
				fmt.Printf("\n%s=\n%s\n", filepath.Base(path), data)
				fmt.Println("--------------------------------------------------")
			}
		}
		return nil
	})
}
