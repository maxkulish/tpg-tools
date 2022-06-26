package findgo

import (
	"os"
	"path"
)

func CountGoFiles(folder string, count int) int {
	files, err := os.ReadDir(folder)
	if err != nil {
		return count
	}

	for _, f := range files {
		if f.IsDir() {
			count = CountGoFiles(folder+"/"+f.Name(), count)
		}
		if path.Ext(f.Name()) == ".go" {
			count++
		}
	}
	return count
}
