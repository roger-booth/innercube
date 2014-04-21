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

type Face [8]Color

type Edge [12]*Color

type Cube struct {
	faceMap map[Color]*Face
	edgeMap map[Color]Edge
}

type ThreeDTransformer struct {
	faceRing ring.Ring
	edgeRing ring.Ring
}

func main() {
	cube1 := new(Cube)
	face1 := &Face{"red", "red", "red", "red", "red", "red", "red", "red"}
	faceMap1 := make(map[Color]*Face)
	faceMap1["red"] = face1
	cube1.faceMap = faceMap1
	edge1 := Edge{&face1[0], &face1[1], &face1[2], &face1[3], &face1[4], &face1[5],
		&face1[6], &face1[7], &face1[0], &face1[1], &face1[2], &face1[3]}
	edgeMap1 := make(map[Color]Edge)
	edgeMap1["red"] = edge1
	cube1.edgeMap = edgeMap1
	*cube1.edgeMap["red"][0] = "blue"
	*cube1.edgeMap["red"][1] = "green"
	fmt.Println(cube1.faceMap["red"][0])
	fmt.Println(cube1.faceMap["red"][1])
	fmt.Println(cube1.faceMap["red"][2])
	fmt.Println(*cube1.edgeMap["red"][2])
}

