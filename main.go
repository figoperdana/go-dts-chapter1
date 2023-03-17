package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var w sync.WaitGroup
var m sync.Mutex

func main() {

	coba := []interface{}{"coba1", "coba2", "coba3"}
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}

	RandomData(coba, bisa)
	LockedData(coba, bisa)

	w.Wait()
}

func RandomData(coba []interface{}, bisa []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 8; i++ {
		w.Add(1)

		go func(id int) {
			defer w.Done()
			if rand.Intn(2) == 0 {
				fmt.Printf("%v %d\n", bisa, id)
			} else {
				fmt.Printf("%v %d\n", coba, id)
			}
		}(i)
	}

}

func LockedData(coba1 []interface{}, bisa1 []interface{}) {

	for i := 1; i <= 8; i++ {
		w.Add(1)

		go func(id int) {
			m.Lock()
			defer m.Unlock()
			if id%2 == 1 {
				fmt.Printf("%v %d\n", coba1, id)
			} else {
				fmt.Printf("%v %d\n", bisa1, id)
			}

			w.Done()
		}(i)
	}
}