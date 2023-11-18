package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

const (
	article_thumbnail_dir = "/articles/thumbnail/"
)

type DropboxResponse struct {
	Name      string `json:"name"`
	Pathlower string `json:"path_lower"`
}

type DropboxTempURL struct {
	LinkURL string `json:"link"`
}

func DropboxUploadEndpoint(file *multipart.FileHeader, category string) (string, error) {

	var path string

	src, err := file.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()

	if category == "articles" {
		path = article_thumbnail_dir + file.Filename
	}

	req, errReq := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", src)

	if errReq != nil {
		return "", fmt.Errorf("Error request : ", errReq)
	}

	req.Header.Set("Authorization", "Bearer "+viper.GetString("STORAGE.DROPBOX_ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Dropbox-API-Arg", `{"path":"`+path+`"}`)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("Error client do : ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		return "", fmt.Errorf(string(body))
	}

	var dropboxRes DropboxResponse

	errDecode := json.NewDecoder(resp.Body).Decode(&dropboxRes)

	if errDecode != nil {
		return "", fmt.Errorf("Error decode : ", errDecode)
	}

	fileURL, ErrFileURL := DropboxGetTempLink(dropboxRes.Pathlower)

	if ErrFileURL != nil {
		return "", ErrFileURL
	}

	return fileURL, nil
}

func DropboxGetTempLink(filepath string) (string, error) {

	requestJson := strings.NewReader(`{
		"path": "` + filepath + `"
	}`)

	req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/get_temporary_link", requestJson)

	if err != nil {
		return "", nil
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+viper.GetString("STORAGE.DROPBOX_ACCESS_TOKEN"))

	client := &http.Client{}

	resp, errReps := client.Do(req)

	if errReps != nil {
		return "", errReps
	}

	defer resp.Body.Close()

	var TempLink DropboxTempURL

	errTempLink := json.NewDecoder(resp.Body).Decode(&TempLink)

	if errTempLink != nil {
		return "", errTempLink
	}

	return TempLink.LinkURL, nil

}
