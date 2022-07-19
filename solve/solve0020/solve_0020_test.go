package solve0020

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func NewUser() *User {
	return &User{}
}

func TestStack(t *testing.T) {
	stack := NewStack()
	user := NewUser()
	user.Name = "zou"
	user.Age = 22
	stack.Push(user)
	// stack.Push(user)
	// stack.Push(user)
	// stack.Push(user)
	stackByte, _ := json.Marshal(stack)
	fmt.Println("stack", string(stackByte))
	fmt.Println("stack", stack)
	for _, element := range stack.elements {
		fmt.Println("element", element)
	}
}

func TestIsValid(t *testing.T) {
	str := "({[})"
	assert.Equal(t, false, isValid(str))
	str1 := "({[]})"
	assert.Equal(t, true, isValid(str1))
	str2 := "()[]"
	assert.Equal(t, true, isValid(str2))
}
