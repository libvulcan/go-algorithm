package back_track

import (
    "fmt"
    "testing"
)

func TestGenerateBrackets(t *testing.T) {
    // 3对括号合法组合
    // (,(,(,),),)
    // (,(,),(,),)
    // (,(,),),(,)
    // (,),(,(,),)
    // (,),(,),(,)
    for _, v := range GenerateBrackets(3) {
        fmt.Println(v)
    }
}
