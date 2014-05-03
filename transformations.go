package main

import (
	"container/ring"
	"fmt"
	"time"
)

type Color string

var colors = [...]Color{"white", "blue", "red", "yellow", "orange", "green"}
var intForColor = map[Color]int{
	"white":  0,
	"blue":   1,
	"red":    2,
	"yellow": 3,
	"orange": 4,
	"green":  5,
}
var edgesForFace = map[Color][]Color{
	"white":  {"red", "green", "orange", "blue"},
	"blue":   {"white", "orange", "yellow", "red"},
	"red":    {"blue", "yellow", "green", "white"},
	"yellow": {"green", "red", "blue", "orange"},
	"orange": {"yellow", "blue", "white", "green"},
	"green":  {"orange", "white", "red", "yellow"},
}

// Based on the coordinate system I discovered in 1984
var edgePos = [...]int{0, 7, 6, 4, 3, 2, 6, 5, 4, 2, 1, 0}

// Pair cube faces by front-to-front projection
// Eight cubes X six faces
var straightProjection = [8][6]int{
	{1, 3, 4, 1, 4, 3},
	{0, 2, 5, 0, 5, 2},
	{3, 1, 6, 3, 6, 1},
	{2, 0, 7, 2, 7, 0},
	{5, 7, 0, 5, 0, 7},
	{4, 6, 1, 4, 1, 6},
	{7, 5, 2, 7, 2, 5},
	{6, 4, 3, 6, 3, 4},
}

type Face [8]Color

type Edge [12]*Color

type Cube struct {
	faceMap map[Color]*Face
	edgeMap map[Color]Edge
}

type Entanglement [8]*Cube

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

func NewEntanglement() (*Entanglement, error) {
	var newEntanglement Entanglement
	for i := 0; i < 8; i++ {
		newEntanglement[i], _ = NewCube()
	}
	return &newEntanglement, nil
}

type ThreeDTransformer struct {
	faceRing *ring.Ring
	edgeRing *ring.Ring
}

func ThreeDRotate(op ThreeDOperation) error {
	newFaceRing := ring.New(8)
	newEdgeRing := ring.New(12)
	trx := ThreeDTransformer{
		newFaceRing, newEdgeRing}
	for _, faceColor := range op.ent[op.cubeId].faceMap[op.face] {
		trx.faceRing.Value = faceColor
		trx.faceRing = trx.faceRing.Next()
	}
	for _, edgeColorPtr := range op.ent[op.cubeId].edgeMap[op.face] {
		trx.edgeRing.Value = *edgeColorPtr
		trx.edgeRing = trx.edgeRing.Next()
	}

	trx.faceRing = trx.faceRing.Move(2 * op.direction)
	trx.edgeRing = trx.edgeRing.Move(3 * op.direction)

	for i, _ := range op.ent[op.cubeId].faceMap[op.face] {
		if v, ok := trx.faceRing.Value.(Color); ok {
			op.ent[op.cubeId].faceMap[op.face][i] = v
		}
		trx.faceRing = trx.faceRing.Next()
	}
	for i, _ := range op.ent[op.cubeId].edgeMap[op.face] {
		if v, ok := trx.edgeRing.Value.(Color); ok {
			*op.ent[op.cubeId].edgeMap[op.face][i] = v
		}
		trx.edgeRing = trx.edgeRing.Next()
	}

	return nil
}

func Sister(cubeId int, face Color) (sisterCubeId int, sisterFace Color) {
	sisterCubeId = straightProjection[cubeId][intForColor[face]]
	sisterFace = face
	return
}

type ThreeDOperation struct {
	ent       *Entanglement
	cubeId    int
	face      Color
	direction int
}

func SplitMessage(op ThreeDOperation, opchan chan ThreeDOperation) {
	sisterCubeId, sisterFace := Sister(op.cubeId, op.face)
	fmt.Printf("sister cubeID: %d, sister face: %s\n", sisterCubeId, sisterFace)
	opchan <- op
	var sisterOp = ThreeDOperation{op.ent, sisterCubeId, sisterFace, op.direction}
	opchan <- sisterOp
	return
}

func countDown(count chan int) {
	for i := 20; i >= 0; i-- {
		count <- i
		time.Sleep(1000000000)
	}
}

func takeSample(ent *Entanglement) {
	fmt.Println(ent[0].faceMap["red"][1])
	fmt.Println(ent[0].faceMap["red"][2])
	fmt.Println(*ent[0].edgeMap["red"][2])
	fmt.Println(*ent[0].edgeMap["red"][3])
	fmt.Println(*ent[0].edgeMap["red"][8])
	fmt.Println(*ent[0].edgeMap["red"][11])
}

func player1(blab chan ThreeDOperation, ent *Entanglement) {
	var sim1 = [...]ThreeDOperation{
		{ent, 0, "red", 1},
	}
	for i := range sim1 {
		blab <- sim1[i]
		time.Sleep(1000000)
	}
}

func player2(blab chan ThreeDOperation, ent *Entanglement) {
	var sim2 = [...]ThreeDOperation{
		{ent, 2, "blue", 1},
	}
	for i := range sim2 {
		blab <- sim2[i]
		time.Sleep(1000000)
	}
}

func main() {
	entanglement1, _ := NewEntanglement()
	operations_p := make(chan ThreeDOperation)
	operations_s := make(chan ThreeDOperation)
	count := make(chan int)
	takeSample(entanglement1)
	go countDown(count)
	go player1(operations_p, entanglement1)
	go player2(operations_p, entanglement1)

	for {
		select {
		case o_p := <-operations_p:
			fmt.Println("Received primary operation")
			go SplitMessage(o_p, operations_s)
		case o_s := <-operations_s:
			fmt.Println("Received secondary operation")
			ThreeDRotate(o_s)
		case i := <-count:
			if 0 == i {
				takeSample(entanglement1)
				return
			}
			fmt.Printf("%d seconds remaining\n", i)
		}
	}
}
