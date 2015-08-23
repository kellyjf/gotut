package main

import (
"fmt" 
"errors" 
"os" 
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack *Stack) Append(item interface{}) {
	*stack = append(*stack, item)	
}

func (stack *Stack) Pop() (item interface{}, err error) {
	astack := *stack
	if len(astack)==0 {
		return nil, errors.New("empty stack")
	}
	*stack = astack[0:len(astack)-1]
	return astack[len(astack)-1],nil 
}

func main() {
	var stack Stack
	var val interface{}
	var err error

	for x, y:= range os.Args {
		stack.Append(os.Args[x])
		fmt.Println("Hello, World " + string(x))
		fmt.Println(x)
		fmt.Println(y)
	}
	fmt.Println(stack.Len())
	for val,err = stack.Pop(); err==nil; val,err=stack.Pop() {
		fmt.Println(val)
	}
	fmt.Println(stack.Len())

}


