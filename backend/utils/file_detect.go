package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// FileInfo holds the metadata for a file.
type FileInfo struct {
	FilePath string `json:"filePath"`
	MTime    int64  `json:"mtime"`
	Size     int64  `json:"size"`
	Hash     string `json:"hash"`
}

// FileIndex is a map that holds FileInfo for each file path.
type FileIndex map[string]FileInfo

// GetFileInfo retrieves file metadata.
func GetFileInfo(filePath string) (*FileInfo, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		FilePath: filePath,
		MTime:    stat.ModTime().Unix(),
		Size:     stat.Size(),
	}, nil
}

// GetFileHash computes the SHA-1 hash of a file.
func GetFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha1.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// LoadIndex loads the file index from a JSON file.
func LoadIndex(filePath string) (FileIndex, error) {
	index := make(FileIndex)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return index, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &index)
	return index, err
}

// SaveIndex saves the file index to a JSON file.
func SaveIndex(filePath string, index FileIndex) error {
	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0644)
}

// DetectChanges detects changes in the files and updates the index.
func DetectChanges(rootDir string, indexFile string) (map[string][]string, error) {
	oldIndex, err := LoadIndex(indexFile)
	if err != nil {
		return nil, err
	}
	newIndex := make(FileIndex)
	changes := map[string][]string{
		"added":    {},
		"removed":  {},
		"modified": {},
	}

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		newFileInfo, err := GetFileInfo(path)
		if err != nil {
			return err
		}

		if oldFileInfo, exists := oldIndex[path]; !exists {
			// New file
			newFileInfo.Hash, err = GetFileHash(path)
			if err != nil {
				return err
			}
			newIndex[path] = *newFileInfo
			changes["added"] = append(changes["added"], path)
		} else {
			// Existing file
			if newFileInfo.MTime != oldFileInfo.MTime || newFileInfo.Size != oldFileInfo.Size {
				newFileInfo.Hash, err = GetFileHash(path)
				if err != nil {
					return err
				}
				if newFileInfo.Hash != oldFileInfo.Hash {
					changes["modified"] = append(changes["modified"], path)
				}
			} else {
				newFileInfo.Hash = oldFileInfo.Hash
			}
			newIndex[path] = *newFileInfo
		}

		return nil
	})
	if err != nil {
		return changes, err
	}

	// Detect removed files
	for oldFile := range oldIndex {
		if _, exists := newIndex[oldFile]; !exists {
			changes["removed"] = append(changes["removed"], oldFile)
		}
	}

	// Save the new index
	err = SaveIndex(indexFile, newIndex)
	if err != nil {
		return changes, err
	}

	return changes, nil
}

func main() {
	rootDir := "E:\\Scripting\\docs-v3"
	indexFile := "file_index.json"

	startTime := time.Now()
	changes, err := DetectChanges(rootDir, indexFile)
	if err != nil {
		log.Fatalf("Failed to detect changes: %v", err)
	}
	endTime := time.Now()

	fmt.Println("Added files:", changes["added"])
	fmt.Println("Removed files:", changes["removed"])
	fmt.Println("Modified files:", changes["modified"])
	fmt.Println("Time taken:", endTime.Sub(startTime).Seconds(), "seconds")
}
