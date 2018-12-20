package testmod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string, arg string, arg2 string) string {
	return fmt.Sprintf("Hi, %s arg:%s", name, arg)
}
