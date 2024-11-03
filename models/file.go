package models

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"
	"storage/config"
	"encoding/base64"
    "mime/multipart"
	"crypto/sha256"
	"encoding/hex"
)

func StoreFileFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return storeFile(content, filepath.Base(url))
}

func CopyFileFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return storeFile(content, filepath.Base(url))
}

func StoreBase64File(data string) (string, error) {
	content, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return storeFile(content, fmt.Sprintf("%d", time.Now().Unix()))
}

func StoreDirectFile(content string) (string, error) {
	return storeFile([]byte(content), fmt.Sprintf("%d", time.Now().Unix()))
}

func StoreUploadedFile(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return storeFile(content, file.Filename)
}

func GetFileHash(file []byte) (string, error) {
	hash := sha256.New()
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func storeFile(content []byte, filename string) (string, error) {
	path := "/storage/" + filename
	hash, err := GetFileHash(content)
	if err!= nil { return "", err }
	
	sqlresult, err := config.DB.Exec("UPDATE vault_files SET lastchange= ? WHERE filepath = ? order by id desc limit 1", time.Now().Unix(), path)
	rowsAffected, err := sqlresult.RowsAffected()
	if(int(rowsAffected) < 1) {
		sqlresult, err = config.DB.Exec("INSERT INTO vault_files (filepath, hashcheck, registered) VALUES (?, ?, ?)", path, hash, time.Now().Unix())
	}
	if err != nil { return "", err }


	sqlresult, err = config.DB.Exec("UPDATE vault_hashes SET lastcheck=? WHERE hashcheck=? order by id desc limit 1", time.Now().Unix(), hash)
	rowsAffected, err = sqlresult.RowsAffected()
	if(int(rowsAffected) < 1) {
		sqlresult, err = config.DB.Exec("INSERT INTO vault_hashes (hashcheck,registered,lastcheck,content) VALUES (?, ?, '0', ?)", hash, time.Now().Unix(), content)
	}

	if err != nil { return "", err }	

	return path, nil
}
