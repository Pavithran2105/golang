package main

import "fmt"

//keyword  identifier  type

type engine struct { //struct
	horsePower int
	torque     float32
}
type engine1 struct { //struct
	horsePower float32
	torque     float32
}

func (bonet engine) car() float32 { //method
	return float32(bonet.horsePower) * bonet.torque
}
func (bonet2 engine1) car() float32 {
	return float32(bonet2.torque) + float32(bonet2.horsePower)
}

type whole interface {
	car() float32
}

func main() {
	e := engine{
		70,
		80.2,
	}
	e1 := engine1{
		90.1,
		2.1,
	}
	wholeVariable := []whole{e, e1}
	h := whole.car(e)
	j := e.car()
	// v := []whole{e, e1}
	// fmt.Println(e.car())
	// fmt.Println(e.horsePower, e.torque)
	// fmt.Println(e1.car())
	fmt.Println(h)
	fmt.Println(j)
	fmt.Println(wholeVariable)

}
