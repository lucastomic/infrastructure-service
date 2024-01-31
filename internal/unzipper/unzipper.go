package unzipper

import (
	"archive/zip"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func Unzip(src multipart.File, size int64, dest string) error {
	zipReader, err := zip.NewReader(src, size)
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		if err := extractFile(file, dest); err != nil {
			return err
		}
	}
	return nil
}

func extractFile(file *zip.File, dest string) error {
	fpath := filepath.Join(dest, file.Name)

	if file.FileInfo().IsDir() {
		return os.MkdirAll(fpath, os.ModePerm)
	}

	if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	fileInZip, err := file.Open()
	if err != nil {
		return err
	}
	defer fileInZip.Close()

	_, err = io.Copy(outFile, fileInZip)
	return err
}
