package requests

import (
	"log"
	"net/http"
	"sync"
)

// Prepare _
type Prepare struct {
	Amount int
	Dest   string
}

// Travel _
func (info Prepare) Travel(wg *sync.WaitGroup) {
	_, err := http.Get(info.Dest)
	if err != nil {
		log.Fatal(err)
	}
	defer wg.Done()
}
