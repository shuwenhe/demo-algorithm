package stackarray

// StackArray Interface
type StackArray interface {
	Clear()                // 清空
	Size() int             // 大小
	Pop() interface{}      // 弹出
	Push(data interface{}) // 压入
	IsFull() bool          // 判断是否已满
	IsEmpty() bool         // 判断是否已空
}

// Stack Struct
type Stack struct {
	dataSource  []interface{} // Slice
	capsize     int           // Caption
	currentsize int           // Current size
}

// NewStack New Stack
func NewStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = make([]interface{}, 10) // Slice
	myStack.capsize = 10                         // Caption
	myStack.currentsize = 0
	return myStack
}

// Clear 清空
func (myStack *Stack) Clear() {
	myStack.dataSource = make([]interface{}, 10) // Slice
	myStack.capsize = 10                         // Cation size
	myStack.currentsize = 0
}

// Size Stack size
func (myStack *Stack) Size() int {
	return myStack.currentsize
}

// Pop Pop
func (myStack *Stack) Pop() interface{} {
	if !myStack.IsEmpty() {
		last := myStack.dataSource[myStack.currentsize-1]             // Last data
		myStack.dataSource = myStack.dataSource[:myStack.currentsize] // Delete last data
		myStack.currentsize--                                         // Delete
		return last
	}
	return nil
}

// Push Push
func (myStack *Stack) Push(data interface{}) {
	if !myStack.IsFull() {
		myStack.dataSource = append(myStack.dataSource, data) // Append push data into stack
		myStack.currentsize++
	}
}

// IsFull 判断是否已满
func (myStack *Stack) IsFull() bool {
	if myStack.currentsize >= myStack.capsize {
		return true
	}
	return false
}

// IsEmpty 判断是否已空
func (myStack *Stack) IsEmpty() bool {
	if myStack.currentsize == 0 {
		return true
	}
	return false
}
