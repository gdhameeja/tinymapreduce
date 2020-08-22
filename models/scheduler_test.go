package models

import (
	"testing"
	"log"
)

func TestSplitFile(t *testing.T) {
	master := Master{1, make(chan bool)}
	master.SplitFile("test.txt")
	log.Println("Everything seems to work")
}
