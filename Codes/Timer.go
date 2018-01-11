package main


import (
	"fmt"
	"time"
)

func main(){
	timer := time.NewTimer(time.Second * 3)			//Timer of time period 3 seconds

	<- timer.C			//Ensures timer expires before continuing execution.
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second*2)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()


	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}


