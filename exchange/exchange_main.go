package main

import(
	"fmt"
	"exchanger"
)

func main(){
	str_p := exchanger.NewStringPair("red","blue")
	int_p := exchanger.NewIntPair(100,200)
	float_p := exchanger.NewFloatPair(12.55,3021.12)
	
	for _,x := range []exchanger.Exchanger{str_p,int_p,float_p} {
		fmt.Println(x.String())
		x.Exchange()
		fmt.Println(x.String())
	}
}
