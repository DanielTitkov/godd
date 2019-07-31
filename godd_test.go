package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestFileLengthCheck(t *testing.T) {
	testContent := []byte("Foobar123 Володя Какого хрена ?!")
	tempFileName := "testLen" + strconv.FormatInt(time.Now().UnixNano(), 10)
	err := ioutil.WriteFile(tempFileName, testContent, 0755)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Open(tempFileName)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	defer os.Remove(tempFileName)

	if fl := getFileLength(file); int(fl) != len(testContent) {
		t.Errorf("Expected %d, got %d", len(testContent), fl)
	}
}

func TestCopyFile(t *testing.T) {
	testContent := []byte("Foobar123 Володя Какого хрена ?!")
	tempFileName := "testCopy" + strconv.FormatInt(time.Now().UnixNano(), 10)
	testOutName := "testOut" + strconv.FormatInt(time.Now().UnixNano(), 10)
	err := ioutil.WriteFile(tempFileName, testContent, 0755)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tempFileName)
	defer os.Remove(testOutName)

	err = copyFile(tempFileName, testOutName, 1024)
	if err != nil {
		t.Error(err)
	}

	if res, err := ioutil.ReadFile(testOutName); !(bytes.Equal(res, testContent)) {
		t.Errorf("Expected %v, got %v", testContent, res)
		if err != nil {
			t.Error(err)
		}
	}
}
