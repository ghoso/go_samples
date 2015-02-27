package main
import (
	"fmt"
	"obj_lib"
)

func main(){
	base_obj := obj_lib.NewBase("Base Test")
	fmt.Println("Base name: ",base_obj.GetName())
	child_obj := obj_lib.NewChild(base_obj,"Child Test")
	fmt.Println("Child name: ",child_obj.GetFullName())
}
