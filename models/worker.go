package models

import (
	"log"
)

var workers map[Worker]bool

const (
	M_factor = 5  // factor by which total number of inputs will be divided to obtain M splits
	R_factor = 2  // factory by which total intermediate pairs will be divided by to obtain R splits
)

type Map func(string, string) (string, string)

type Reduce func(string, []string)

type Worker struct {
	Id     int
	IsDone chan bool
}

func (w *Worker) PerformMap(chunkPath string, mapFunc Map) {
	// read the chunkPath and call mapFunc for every word

}

func (w *Worker) PerformReduce(key string, intermediateValues []string, redFunc Reduce) {
	// iterate over intermediate values for key and call reduce
}

func (w *Worker) MakeMaster(chunkPaths... string, mapFunc Map, redFunc Reduce) {
	master := Master{chunkPaths: chunkPaths, mapFunc: mapFunc, redFunc: redFunc}
}

type Task struct {
}

type Master struct {
	chunkPaths []string
	intermediateKeyValuePairs [[]string]string

	mapFunc Map
	redFunc Reduce
}


func (m *Master) AssignMappers() {
	for i := 0; i < len(m.chunkPaths); i++ {
		for worker, status := range workers {
			if !status {
				worker.PerformMap(chunkPaths[i])
			}
		}
		// TODO: if none of the workers are free this function will just return and letover map jobs will
		// never be assigned. Figure out a way of polling the workers to assign a new mapper when any worker
		// is available.
	}
}
