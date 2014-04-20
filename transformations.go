package main
 
import (
	"container/ring"
)

type Color string

var colors = [...]Color {"white","blue","red","yellow","orange","green"}

type Face []Color

type Cube map[Color]Face

type ThreeDTransformer struct {
    faceRing ring.Ring
    edgeRing ring.Ring
}

func main() {

}
