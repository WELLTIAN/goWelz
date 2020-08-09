package main
import (
	"fmt"
	//"time"
	"sync"
)

var waitgroup sync.WaitGroup

func Testproductor(c chan string){
	for i:=1;i<=100000000;i++{
		c<-"你是个猪猪"
		//waitgroup.Done()
	}
}

func TestCustomer(c chan string){
	for i:=1;i<=10000;i++{
		fmt.Println(<-c)
	}
	waitgroup.Done()
}

func main() {
	waitgroup.Add(10000)
	c :=make(chan string,100000000)
	go Testproductor(c)
	for i:=1;i<=10000;i++{
		go TestCustomer(c)
	}	
	//time.Sleep(1 * time.Second)
	waitgroup.Wait()
	fmt.Println("ok")
}