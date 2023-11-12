package bluelion

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func backupFile(filePath string) (string, error) {
	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	// Extract the original filename and its extension
	originalFilename := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	extension := filepath.Ext(filePath)

	// Extract the directory of the original file
	originalDir := filepath.Dir(filePath)

	// List existing backups
	backupFiles, err := filepath.Glob(filepath.Join(originalDir, ".*"))
	if err != nil {
		return "", err
	}

	// Sort backup files based on their timestamp in reverse order
	sort.Sort(sort.Reverse(sort.StringSlice(backupFiles)))

	maxBackupCount := 10

	// Ensure we have at most 10 backups
	for len(backupFiles) > maxBackupCount {
		oldestBackup := backupFiles[len(backupFiles)-1]
		if err := os.Remove(oldestBackup); err != nil {
			return "", err
		}
		backupFiles = backupFiles[:len(backupFiles)-1]
	}

	// Create a new backup filename with the current epoch timestamp, original filename, and extension
	epochTimestamp := time.Now().Unix()
	newBackupFile := filepath.Join(originalDir, fmt.Sprintf(".%d.%s%s", epochTimestamp, originalFilename, extension))

	// Copy the file to the new backup file
	err = copyFile(filePath, newBackupFile)
	if err != nil {
		return "", err
	}

	slog.Debug("backup", "max", maxBackupCount, "count", len(backupFiles), "status", "success", "path", newBackupFile)

	return newBackupFile, nil
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
