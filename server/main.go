package main
import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)
type Listener int
type Reply struct {
	Data string
}
type Details struct {
	Name string
	Password string
}
type InvalidData struct {}

func (i * InvalidData) Error() string{
	return "error in something"
}
func (l *Listener) GetLine(line []byte, reply *Reply) error {
	rv := string(line)
	fmt.Printf("Receive: %v\n", rv)
	*reply = Reply{rv}
	return nil
}
func Validation(c Details) error{
	if c.Name != "llchh" || c.Password != "sad"{
		return &InvalidData{}
	}
	return nil
}
func (l* Listener) Auth(c Details, reply *Reply) error{
	err := Validation(c)
	if err != nil{
		return err
	}
	return nil
}
func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}