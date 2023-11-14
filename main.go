package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("ディレクトリの読み込みに失敗しました: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		ext := filepath.Ext(filename)
		if isImageFile(ext) {
			hashedName, err := hashFileData(filename)
			if err != nil {
				log.Printf("ファイルデータのハッシュ化に失敗しました: %v", err)
				continue
			}

			newName := "image-" + hashedName + ext
			if err := os.Rename(filename, newName); err != nil {
				log.Printf("ファイルのリネームに失敗しました: %v", err)
			}
		}
	}
}

// 画像ファイルかどうかを判定する
func isImageFile(ext string) bool {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return true
	default:
		return false
	}
}

// ファイルのデータからSHA256ハッシュを計算する
func hashFileData(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}
