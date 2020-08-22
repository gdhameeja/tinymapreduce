package models

import (
	"os"
	"log"
	"fmt"
	"time"
	"io/ioutil"
)


func RegisterWorker(w Worker) {
	workers[w] = false
}

func ScheduleJob(filname string) {
	chunkPaths := SplitFile(filename)
	for worker, status := range workers {
		if !status {
			worker.MakeMaster(chunkPaths)
			return
		}
	}
	time.Sleep(5 * time.Second)
	ScheduleJob(filename)
}

func SplitFile(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Print("Something went wrong while opening file ", filename)
		m.IsDone <- true
		return
	}

	fileInfo, _ := file.Stat()
	chunkSize := uint64(fileInfo.Size() / M_factor)
	var chunkPaths []string

	for i := 0; i < M_factor; i++ {
		buffer := make([]byte, chunkSize)
		file.Read(buffer)
	
		// write chunk to disk
		tempDir := fmt.Sprintf("/tmp/master_%d", m.Id)
		chunkFileName := fmt.Sprintf("%s/chunk_%d", tempDir, i)
		chunkPaths = append(chunkPaths, chunkFileName)
		os.MkdirAll(tempDir, 0755) 
		_, err := os.Create(chunkFileName)
		if err != nil {
			log.Print("Error creating file ", chunkFileName)
			m.IsDone <- true
			return
		}

		ioutil.WriteFile(chunkFileName, buffer, os.ModeAppend)
	}
	return chunkPaths
}
