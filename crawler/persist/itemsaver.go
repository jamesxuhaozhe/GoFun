package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			log.Printf("Got item %v\n", item)
		}
	}()
	return out
}
