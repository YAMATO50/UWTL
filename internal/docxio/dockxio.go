package docxio

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	sliceutil "github.com/YAMATO50/UWTL/internal/sliceUtil"
)

func Unzip(path string, destination string) error {
	if !strings.HasSuffix(path, ".docx") {
		return fmt.Errorf("unknown file of type %s", sliceutil.Last(strings.Split(path, ".")))
	}

	err := unzipAll(path, destination)
	if err != nil {
		return err
	}

	return nil
}

func unzipAll(path string, destination string) error {
	reader, err := zip.OpenReader(path)

	if err != nil {
		return err
	}
	defer reader.Close()

	dest, err := filepath.Abs(destination)

	if err != nil {
		return err
	}

	for _, f := range reader.File {
		err = unzipFile(f, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File, destination string) error {
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}

	return nil
}
