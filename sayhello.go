package helloworld

import "fmt"

func SayHelloWorld() {
	fmt.Println("Gitlab Hello world!")
}

func Version() string {
	return "v2.0.0"
}

func SayOther(str string) {
	fmt.Printf("Gitlab Saying: %s\n", str)
}
