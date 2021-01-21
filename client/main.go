package main
import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Details struct {
	Name string
	Password string
}
type Reply struct {
	Data string
}
func main() {
	client, err := rpc.Dial("tcp", "localhost:8082")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(os.Stdin)
	fmt.Println("login: ")
	line, _, err := in.ReadLine()
	c := Details{}
	c.Name = string(line)
	fmt.Println("pass: ")
	line, _, err = in.ReadLine()
	c.Password = string(line)
	if err != nil {
		log.Fatal(err)
	}
	var reply Reply
	err = client.Call("Listener.Auth", c, &reply)
	if err != nil {
		log.Fatal(err)
	}
	for {
		line, _, err = in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		err = client.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v, Data: %v", reply, reply.Data)
	}
}