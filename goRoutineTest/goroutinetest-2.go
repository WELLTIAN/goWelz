package main
import (
	"fmt"
	"time"
	"sync"
	"os"
)

var waitgroup sync.WaitGroup
var waitgroup2 sync.WaitGroup
var sumnum int
var mu sync.Mutex  
var mu1 sync.Mutex  

var f, err = os.Create("test.txt")

func alladd() {
    mu.Lock()
    sumnum = sumnum + 1
    mu.Unlock()
}
func delete() {
    mu.Lock()
    sumnum = sumnum - 1
    mu.Unlock()
}

func addlistitem(str string){
    mu1.Lock()
    f.WriteString(str)
    mu1.Unlock()
}


func Testproductor(c chan string){
	for i:=1;i<=10000;i++{
		 c<-"你是个猪猪\n"
		//waitgroup.Done()
		alladd()
	}
	waitgroup2.Done()
}

func TestCustomer(c chan string){
	for{
		if sumnum == 0 && len(c) == 0{
			break
		}else if len(c) == 0{
			time.Sleep(time.Duration(100) * time.Millisecond)
		}	
		delete()	
		addlistitem(<-c)
	}
	waitgroup.Done()
}

func main() {
	sumnum = 0
	//f, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	defer f.Close()
	waitgroup.Add(10000)
	waitgroup2.Add(10000)
	c :=make(chan string,10000)
	for i:=1;i<=10000;i++{
		go Testproductor(c)
		go TestCustomer(c)
	}	
	//time.Sleep(1 * time.Second)
	waitgroup2.Wait()
	waitgroup.Wait()
	fmt.Println("ok")
}