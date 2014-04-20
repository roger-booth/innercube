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

}
