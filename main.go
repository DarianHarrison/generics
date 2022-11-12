package main

import (
    "log"
)

// constraints are pretty useless for generic use
// type I interface {
// 	A | B
// }

type A struct {
	a uint8
}

type B struct {
	b uint16
}


func genericA(i any) {
    switch i.(type) {
        case A:
            log.Printf("Type %T\n", i)
        case B:
            log.Printf("Type %T\n", i)
        default:
            log.Printf("Unknown Type %T\n", i)
    }
}

// func genericB[T any](i T) { // will not compile it thinks T any is a constraint that cant be used in switch statement (if you use a switch statement as in eample above "genericA")
// func genericB[T I](i T) { // pretty useless for generic use, not recommended // cant use case switch inside

func main() {

	// The language specifications explicitly disallow using interfaces with type elements as anything other than type parameter constraints
	messages := make([]any, 2)
    messages[0] = A{1}
    messages[1] = B{2}

    genericA(A{1}) // concrete type inference
    // genericA[B](B{2}) // concrete type with inference // this will not work for switch statements as constrains only work as argument parameters
    // genericA(&A{1}) // will not work // will give unknown type

    genericA(messages[1]) // type inference // will NOT compile
    // genericA[B](messages[1]) // concrete type with inference // this will not work for switch statements as constrains only work as argument parameters

}