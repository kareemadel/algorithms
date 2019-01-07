package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func (g *graph) deepCopy() *graph {
	g.lock.Lock()
	defer g.lock.Unlock()

	newG := graph{}
	newG.init()
	for _, v := range g.adList {
		n, _ := newG.createVertex(v.label)
		n.importEdges(v)
	}
	return &newG
}

func (g *graph) print() {
	var strBuilder strings.Builder
	for _, v := range g.adList {
		strBuilder.WriteString(fmt.Sprintf("%s\t", v.label))
		for l, w := range v.edges {
			strBuilder.WriteString(fmt.Sprintf("%s(%d)\t", l, w))
		}
		s := strBuilder.String()
		fmt.Println(s[:len(s)-1])
		strBuilder.Reset()
	}
}

func (g *graph) readFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	reader.Read()
	tsvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var vLabel, uLabel string
	var weight int
	for _, line := range tsvData {
		vLabel = line[0]
		uLabel = line[1]
		weight, _ = strconv.Atoi(line[2])
		g.createVertex(vLabel)
		g.createVertex(uLabel)
		g.addEdge(vLabel, uLabel, weight)
	}
}

func (g *graph) init() {
	g.adList = make(map[string]*vertex)
}

func (g *graph) addEdge(vLabel, uLabel string, weight int) bool {
	v := g.getVertex(vLabel)
	u := g.getVertex(uLabel)
	if v == nil || u == nil {
		return false
	}
	v.addEdge(uLabel, weight)
	u.addEdge(vLabel, weight)
	return true
}

func (g *graph) getVertex(label string) *vertex {
	if g.isVertex(label) {
		return g.adList[label]
	}
	return nil
}

func (g *graph) createVertex(label string) (*vertex, bool) {
	isNew := false
	if !g.isVertex(label) {
		newVertex := &vertex{
			label: label,
		}
		g.addVertex(newVertex)
		isNew = true
	}
	return g.getVertex(label), isNew
}

func (g *graph) addVertex(v *vertex) bool {
	if g.isVertex(v.label) {
		return false
	}
	g.adList[v.label] = v
	return true
}

func (g *graph) isVertex(label string) bool {
	_, ok := g.adList[label]
	return ok
}

type graph struct {
	adList map[string]*vertex
	lock   sync.Mutex
}
