package main
 
import (
	"container/ring"
)

type Color string

var colors = [...]Color {"white","blue","red","yellow","orange","green"}

type Face [8]Color

type Edge [12]*Color

type Cube struct {
    faceMap map[Color]Face
    edgeMap map[Color]Edge
}    

type ThreeDTransformer struct {
    faceRing ring.Ring
    edgeRing ring.Ring
}

func main() {
    cube1 := new(Cube)
    face1 := Face {"red","red","red","red","red","red","red","red"}
    faceMap1 := make(map[Color]Face)
    faceMap1["red"] = face1
    cube1.faceMap = faceMap1
    edge1 := Edge {&face1[0], &face1[0], &face1[0], &face1[0], &face1[0], &face1[0],
                   &face1[0], &face1[0], &face1[0], &face1[0], &face1[0], &face1[0]}
    edgeMap1 := make(map[Color]Edge)
    edgeMap1["red"] = edge1
    cube1.edgeMap = edgeMap1
}
