package main

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	snippet := BuildSnippet(2)

	fmt.Printf("Descriptin: %s\nPrice: %0.2f рублей\nВремя работы сервиса: %v", snippet.Description, snippet.Price, time.Since(start))

}

// itemDescription симулирует получение описания из сервиса
func itemDescription(itemID int) string {
	time.Sleep(time.Second * 1)
	return "☆☆☆Lorem ♡♡♡ ipsum  dolor sit amet..."
}

// itemPrice симулирует получение цены в долларах из сервиса
func itemPrice(itemID int) float64 {
	time.Sleep(2 * time.Second)
	return 100
}

func prettify(description string) string {
	time.Sleep(time.Second * 1)
	re := regexp.MustCompile("\\W+")
	return re.ReplaceAllString(description, "")
}

func priceToRub(price float64) float64 {
	time.Sleep(1 * time.Second)
	return price * 70
}

type Snippet struct {
	Price       float64
	Description string
}

// write your code here
// при использовании горутин время работы сервиса снизилось на 40%
// без использования мьютекса  и с его использованием гонки горутин не замечено
func BuildSnippet(itemId int) Snippet {
	snippet := Snippet{}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		description := itemDescription(itemId)
		prettyDescription := prettify(description)
		mu.Lock()
		snippet.Description = prettyDescription
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		dollarPrice := itemPrice(itemId)
		rubPrice := priceToRub(dollarPrice)
		mu.Lock()
		snippet.Price = rubPrice
		mu.Unlock()
	}()
	wg.Wait()

	return snippet

}
