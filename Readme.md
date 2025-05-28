Мы делаем сервис, который собирает сниппет товара. Каждый сниппет состоит из отформатированного описания и стоимости товара в рублях.

Чтобы собрать сниппет нужно:

Получить описание из одного сервиса, а затем отформатированное его через prettify
Получить цену (в долларах) из другого сервиса, а затем перевести ее в рубли через priceToRub
Вернуть готовый сниппет


package main

// itemDescription симулирует получение описания из сервиса
func itemDescription(itemID int) string {
    time.Sleep(time.Second * 1)
    return "☆☆☆Lorem ♡♡♡ ipsum  dolor sit amet..."
}

// itemPrice симулирует получение цены в долларах из сервиса
func itemPrice(itemId int) float64 {
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

func BuildSnippet(itemId int) Snippet {
    // write your code here
}
