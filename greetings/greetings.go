package greetings

import "fmt"

// Hello This is basically a function that is to be called from somewhere else
// Hello function name (parameterName parameterType) returnType
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
