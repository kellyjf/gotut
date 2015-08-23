package main

import (
"fmt" 
"errors" 
"strings" 
//"io/ioutil" 
"bufio" 
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
	var passwd []byte
	var err error
	var line string
//	passwd, err = ioutil.ReadFile("./passwd")
	file, err := os.Open("./passwd")
	if err != nil  {
		return
	}
	reader := bufio.NewReader(file)
	for line,err = reader.ReadString('\n'); err==nil; line,err=reader.ReadString('\n')  {
		fmt.Print(line)
		fields := strings.Split(line,":")
		for _,f := range fields {
			fmt.Println(f)
		}
	}
		
	if err == nil {
		line:=string(passwd)
		fmt.Println(line)
	}
}

func stack_main() {
	var stack Stack
	var val interface{}
	var err error

	for x := range os.Args {
		stack.Append(os.Args[x])
	}
	fmt.Println(stack.Len())
	for val,err = stack.Pop(); err==nil; val,err=stack.Pop() {
		fmt.Println(val)
	}
	fmt.Println(stack.Len())

}


