package chess

type item struct {
	value interface{}
	next  *item
}

// Stack is a simple stack for stuff
type Stack struct {
	size int
	top  *item
}

// Push adds a value onto the stack
func (stack *Stack) Push(value interface{}) {
	newItem := &item{next: stack.top, value: value}
	stack.top = newItem
	stack.size++
}

// Pop removes the top most item from the stack
func (stack *Stack) Pop() interface{} {
	var retVal interface{}
	if stack.Len() > 0 {
		retVal = stack.top.value
		stack.top = stack.top.next
		stack.size--
	}
	return retVal
}

// Len returns how many items are in the stack
func (stack *Stack) Len() int {
	return stack.size
}
