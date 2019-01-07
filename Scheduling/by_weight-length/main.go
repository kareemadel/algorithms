package main

import "fmt"

func main() {
	s := readScheduleFromFile("/home/kareemadel/algorithms/Scheduling/by_weight-length/jobs.txt")
	s.init()
	fmt.Println(s.getCompletionTime())
}
