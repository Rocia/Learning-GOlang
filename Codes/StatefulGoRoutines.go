package main
//One go routine owns the state in this method to ensure that the data is never corrupted by concurrent access.
//To acquire and release access to the state a message is sent and received respectively by the owning go routine containing the request and a way of owning the routine
import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    var readOps uint64 = 0				//count read and write Ops
    var writeOps uint64 = 0

    reads := make(chan *readOp)			//These read and write channels are used to access the hold the read and write structs as declared above
    writes := make(chan *writeOp)

    go func() {
        var state = make(map[int]int)			//Go routine that owns the state which is a map of the state that selects on the readOps and writeOps counter
        for {
            select {
            case read := <-reads:				//Case where there is a message in the read channel
            	fmt.Println("Read Operation")//Some Ops
                read.resp <- state[read.key]	//sends a value indicating success and the desired value to be read

            case write := <-writes:				//Write case where there is a message in the write channel
            	fmt.Println("Write Operation")//Some Ops
                state[write.key] = write.val	//Write Value to state
                write.resp <- true				//Indicate action completed by sync
            }
        }
    }()

    // This starts 100 goroutines to issue reads to the
    // state-owning goroutine via the `reads` channel.
    // Each read requires constructing a `readOp`, sending
    // it over the `reads` channel, and the receiving the
    // result over the provided `resp` channel.
    for r := 0; r < 5; r++ {				//starts 50 GoRoutines that request reads to the state owning GoRoutine
        go func() {
            for {
                read := &readOp{			//Create a ReadOp of the specified structure
                    key:  rand.Intn(5),	//set the key as specified
                    resp: make(chan int)}	//Create response channel as requested
                reads <- read				//Send this ReadOp structure over the reads channel
                <-read.resp					//Wait for Read Response
                atomic.AddUint64(&readOps, 1)	//Increment the ReadOps atomic Counter
                time.Sleep(time.Millisecond)//Sleep for 1ms before exiting
            }
        }()
    }

    for w := 0; w < 2; w++ {				//Start 10 go routines requesting writes
        go func() {
            for {
                write := &writeOp{			//Create Write structure as declared above
                    key:  rand.Intn(5),	//Set key
                    val:  rand.Intn(100),//Set Value
                    resp: make(chan bool)}	//Create boolean response channel
                writes <- write				//send writeOps structure over the write channel
                <-write.resp				//Wait for write response
                atomic.AddUint64(&writeOps, 1)//Increment the write counter
                time.Sleep(time.Millisecond)//Sleep for 1ms before exiting
            }
        }()
    }

    // Let the goroutines work for a second.
    time.Sleep(time.Millisecond*1)					//Let all the goroutines work for 1 second

    // Finally, capture and report the op counts.
    readOpsFinal := atomic.LoadUint64(&readOps)	//Then capture and report the readOps counter and WriteOps data structure values
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}


/*
Read Operation
Read Operation
Write Operation
Read Operation
Read Operation
Read Operation
Write Operation
Read Operation
readOps: 6
writeOps: 2
 */