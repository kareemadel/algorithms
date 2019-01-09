package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	edges, nodes := readEdges("/home/kareemadel/algorithms/Graphs/clustering/clustering1.txt")
	radixSort(edges)
	var m, n *node
	l := len(nodes)
	for i := 0; l > 4; i++ {
		m = nodes[edges[i].v]
		n = nodes[edges[i].u]
		if m.find() != n.find() {
			m.union(n)
			l--
		}
	}
	for i := 0; i < len(edges); i++ {
		m = nodes[edges[i].v]
		n = nodes[edges[i].u]
		if m.find() != n.find() {
			fmt.Println(edges[i].weight)
			break
		}
	}
}

func radixSort(e []*edge) {
	max, isNegative := getMaxDigit(e)
	for i := 1; i <= max; i++ {
		sortByDigit(e, i)
	}
	if isNegative {
		sortBySign(e)
	}
}

func getMaxDigit(e []*edge) (int, bool) {
	var max, buf int
	var isNegative bool
	for _, v := range e {
		buf = len(strconv.Itoa(v.weight))
		if v.weight < 0 {
			buf--
			isNegative = true
		}
		if buf > max {
			max = buf
		}
	}
	return max, isNegative
}

func sortByDigit(e []*edge, d int) {
	if d == 0 {
		return
	}
	n := len(e)
	exp := int(math.Pow10(d - 1))
	output := make([]*edge, n)
	count := make([]int, 10)
	for _, v := range e {
		count[(v.weight/exp)%10]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		output[count[(e[i].weight/exp)%10]-1] = e[i]
		count[(e[i].weight/exp)%10]--
	}
	copy(e, output)
}

func sortBySign(e []*edge) {
	n := len(e)
	output := make([]*edge, n)
	count := make([]int, 2)
	for _, v := range e {
		if v.weight < 0 {
			count[0]++
		} else {
			count[1]++
		}
	}
	count[1] += count[0]
	for i := n - 1; i >= 0; i-- {
		if e[i].weight < 0 {
			output[count[0]-1] = e[i]
			count[0]--
		} else {
			output[count[1]-1] = e[i]
			count[1]--
		}
	}
	copy(e, output)
}

func readEdges(path string) ([]*edge, map[string]*node) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var edges []*edge
	nodes := make(map[string]*node)
	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	reader.Read()
	ssvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var e *edge
	var n *node
	for _, line := range ssvData {
		weight, _ := strconv.Atoi(line[2])
		e = &edge{line[0], line[1], weight}
		edges = append(edges, e)
		_, ok := nodes[line[0]]
		_, ok1 := nodes[line[1]]
		if !ok {
			n = &node{label: line[0]}
			n.parent = n
			nodes[line[0]] = n
		}
		if !ok1 {
			n = &node{label: line[1]}
			n.parent = n
			nodes[line[1]] = n
		}
	}
	return edges, nodes
}
