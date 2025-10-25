package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// renameFile renames a file from the old name to the new name.
func RenameFile(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	if err != nil {
		return fmt.Errorf("error renaming file: %v", err)
	}
	return nil
}

func CreateRequestAndGetResponse[T any](path string, queryParams map[string]string) (T, error) {
	u, err := url.Parse(API_URL + path)
	if err != nil {
		return *new(T), fmt.Errorf("error parsing URL: %v", err)
	}

	// Add query parameters to the URL
	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	q.Set("language", language_param)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return *new(T), fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+getAPIKey())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *new(T), fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *new(T), fmt.Errorf("error reading response body: %v", err)
	}

	var response T
	if err := json.Unmarshal(body, &response); err != nil {
		return *new(T), fmt.Errorf("error unmarshaling response: %v", err)
	}
	return response, nil
}

// ListFilesInDirectory lists the names of files in the specified directory.
func ListFilesInDirectory() ([]string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	var mkvFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".mkv" {
			mkvFiles = append(mkvFiles, file.Name())
		}
	}

	return mkvFiles, nil
}

func getAPIKey() string {
	if apiKey != "" {
		return apiKey
	}
	keyFile, err := os.ReadFile("key.txt")
	if err != nil {
		fmt.Println("Error reading API key:", err)
		return ""
	}
	return string(keyFile)
}
