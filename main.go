package backend

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("hello %s", name)
}