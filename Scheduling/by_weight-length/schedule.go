package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func (s *schedule) init() {
	s.sort(0, len(s.jobs)-1)
	s.calCompletionTimes()
}

func (s *schedule) getCompletionTime() int {
	return s.completionTimes[len(s.completionTimes)-1]
}

func (s *schedule) calCompletionTimes() {
	sumLength := 0
	sumTime := 0
	for i, v := range s.jobs {
		sumLength += v.length
		sumTime += sumLength * v.weight
		s.completionTimes[i] = sumTime
	}
}

func (s *schedule) sort(start, end int) {
	if start >= end {
		return
	}
	pivotIndex := s.partion(start, end)
	s.sort(start, pivotIndex-1)
	s.sort(pivotIndex+1, end)
}

func (s *schedule) partion(start, end int) int {
	jobs := s.jobs
	n := rand.Intn(end-start) + start
	jobs[start], jobs[n] = jobs[n], jobs[start]
	rightPartionEnd := start
	for i := start + 1; i <= end; i++ {
		if jobs[i].less(jobs[start]) {
			rightPartionEnd++
			if i > rightPartionEnd {
				jobs[rightPartionEnd], jobs[i] = jobs[i], jobs[rightPartionEnd]
			}
		}
	}
	jobs[rightPartionEnd], jobs[start] = jobs[start], jobs[rightPartionEnd]
	return rightPartionEnd
}

func readScheduleFromFile(path string) schedule {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	sizeText, _ := reader.Read()
	size, _ := strconv.Atoi(sizeText[0])
	s := schedule{
		jobs:            make([]*job, size),
		completionTimes: make([]int, size),
	}
	ssvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var weight, length int
	var cost float64
	for i, line := range ssvData {
		weight, _ = strconv.Atoi(line[0])
		length, _ = strconv.Atoi(line[1])
		// cost = weight - length
		cost = float64(weight) / float64(length)
		s.jobs[i] = &job{
			weight: weight,
			length: length,
			cost:   cost,
		}
	}
	return s
}

type schedule struct {
	jobs            []*job
	completionTimes []int
}
