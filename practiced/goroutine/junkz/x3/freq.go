package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func countLetters(url string, freqTable []int, mu *sync.Mutex) {
	body, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer body.Body.Close()
	wrd, err := io.ReadAll(body.Body)
	if err != nil {
		log.Println(err)
		return
	}
	mu.Lock()

	for _, letter := range string(wrd) {
		c := strings.ToLower(string(letter))
		cindex := strings.Index(letters, c)
		if cindex >= 0 {
			freqTable[cindex]++
		}
	}
	mu.Unlock()
	fmt.Println("Completed: ", url, time.Now().Format("15:04:04"))
}
