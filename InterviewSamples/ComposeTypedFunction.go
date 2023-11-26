package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	fileNameTransformFunc TransformFunc
}

func (s *Server) handleRequest(fileName string) error {
	newFileName := s.fileNameTransformFunc(fileName)
	fmt.Println("new filename: ", newFileName)

	return nil
}

func hashFileName(fileName string) string {
	hash := sha256.Sum256([]byte(fileName))
	newFileName := hex.EncodeToString(hash[:])

	return newFileName
}

func PrefixFileName(prefix string) TransformFunc {
	return func(fileName string) string {
		return prefix + fileName
	}
}

func main() {
	s := &Server{
		//fileNameTransformFunc: hashFileName,
		fileNameTransformFunc: PrefixFileName("Sefik"),
	}
	s.handleRequest("cool_picture.jpg")
}
