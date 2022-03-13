package io

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

type FileReader struct {
	Name string
	IsTest bool
}

func NewFileReader(name string, isTest bool) *FileReader {
	return &FileReader{
		Name: name,
		IsTest: isTest,
	}
}

func (r *FileReader) Read(ch chan string) {
	suffix := ""
	if (r.IsTest) {
		suffix = "test"
	}

	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/input/%s-%s.txt", currentPath, r.Name, suffix)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}

	close(ch)
}