package requests

import (
	"log"
	"net/http"
	"sync"
)

// Prepare 訪問所需
type Prepare struct {
	Amount int
	Dest   string
}

// Travel 開始訪問
func (info Prepare) Travel(wg *sync.WaitGroup) {
	_, err := http.Get(info.Dest)
	if err != nil {
		log.Fatal(err)
	}
	defer wg.Done()
}
