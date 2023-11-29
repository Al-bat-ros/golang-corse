package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	inFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	fileInf, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if fileInf.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if fileInf.Size() == 0 {
		return ErrUnsupportedFile
	}

	sizeFile := fileInf.Size() - offset
	if fileInf.Size()-offset > limit && limit != 0 {
		sizeFile = limit
	}
	// create file

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer toFile.Close()
	bar := pb.Full.Start64(sizeFile)
	barRender := bar.NewProxyReader(inFile)
	_, err = io.CopyN(toFile, barRender, sizeFile)
	bar.Finish()
	return err
}
