package main
import (
    "fmt"
    "sync"
  )

var wg sync.WaitGroup
func fun(ch chan int, a int){
  wg.Done()
	fmt.Println(a + <-ch)
}

func main(){
  ch := make(chan int)
  ele := []int{22,33,11,55}
  for i := range(ele){
    wg.Add(1)
    go fun(ch, 33)
    ch <- ele[i]
  }
  wg.Wait()
  close(ch)
  _, ok:= <-ch
  if !ok {
    fmt.Println("Channel Closed!!")
  }
}

