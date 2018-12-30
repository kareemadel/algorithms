package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
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

func (g graph) print() {
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
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	tsvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var vLabel, uLabel string
	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}
	for _, line := range tsvData {
		vLabel = line[0]
		g.createVertex(vLabel)
		for j := 1; j < len(line)-1; j++ {
			s := strings.FieldsFunc(line[j], f)
			uLabel = s[0]
			weight, _ := strconv.ParseUint(s[1], 10, 64)
			g.createVertex(uLabel)
			g.addEdge(vLabel, uLabel, uint(weight))
			g.addReverseEdge(uLabel, vLabel, uint(weight))
		}
	}
}

func (g *graph) init() {
	g.adList = make(map[string]*vertex)
}

func (g *graph) addEdge(vLabel, uLabel string, count uint) bool {
	n := g.getVertex(vLabel)
	if n == nil {
		return false
	}
	n.addEdge(uLabel, count)
	return true
}

func (g *graph) addReverseEdge(vLabel, uLabel string, count uint) bool {
	n := g.getVertex(vLabel)
	if n == nil {
		return false
	}
	n.addReverseEdge(uLabel, count)
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
