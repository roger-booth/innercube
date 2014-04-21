package main

import (
	"container/ring"
	"fmt"
)

type Color string

var colors = [...]Color{"white", "blue", "red", "yellow", "orange", "green"}
var edgesForFace = map[Color][]Color{
	"white":  {"red", "green", "orange", "blue"},
	"red":    {"blue", "yellow", "green", "white"},
	"blue":   {"white", "orange", "yellow", "red"},
	"yellow": {"green", "red", "blue", "orange"},
	"orange": {"yellow", "blue", "white", "green"},
	"green":  {"orange", "white", "red", "yellow"},
}

var edgePos = [...]int{0, 1, 2, 4, 3, 2, 4, 5, 6, 6, 7, 0}

type Face [8]Color

type Edge [12]*Color

type Cube struct {
	faceMap map[Color]*Face
	edgeMap map[Color]Edge
}

func NewCube() (*Cube, error) {
	newFaceMap := make(map[Color]*Face)
	newEdgeMap := make(map[Color]Edge)
	for _, color := range colors {
		newFaceMap[color] = &Face{color, color, color, color, color, color, color, color}
	}
	i := 0
	for _, faceColor := range colors {
		var newEdge Edge
		for _, edgeColor := range edgesForFace[faceColor] {
		        //fmt.Println(faceColor)
		        //fmt.Println(i)
			newEdge[i] = &newFaceMap[edgeColor][edgePos[i]]
			newEdge[i+1] = &newFaceMap[edgeColor][edgePos[i+1]]
			newEdge[i+2] = &newFaceMap[edgeColor][edgePos[i+2]]
			i += 3
			if i == 12 {
			    i = 0
			}
		}
		newEdgeMap[faceColor] = newEdge
	}
	return &Cube{newFaceMap, newEdgeMap}, nil
}

type ThreeDTransformer struct {
	faceRing ring.Ring
	edgeRing ring.Ring
}

func main() {
	cube1,_ := NewCube()

	//fmt.Println(cube1)
	fmt.Println(cube1.faceMap["red"][1])
	fmt.Println(cube1.faceMap["red"][2])
	fmt.Println(*cube1.edgeMap["red"][2])
	fmt.Println(*cube1.edgeMap["red"][3])
	fmt.Println(*cube1.edgeMap["red"][8])
	fmt.Println(*cube1.edgeMap["red"][11])
}
