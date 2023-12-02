package main

import "fmt"

type Posilka interface {
	SaveGiver() string
	SaveReciver() string
	Send()
}

type Box struct {
	Reciver string
	Giver   string
}

type Convert struct {
	Reciver string
	Giver   string
}

func (b *Box) Send() {
	fmt.Println("Відправлення коробки:", b)
}

func (b *Box) SaveGiver() string {
	fmt.Println("Введите адрес отправителя:")
	fmt.Scanln(&b.Giver)
	return b.Giver
}

func (b *Box) SaveReciver() string {
	fmt.Println("Введите адрес получателя:")
	fmt.Scanln(&b.Reciver)
	return b.Reciver
}

func (c *Convert) Send() {
	fmt.Println("Відправлення конверту:", c)
}

func (c *Convert) SaveGiver() string {
	fmt.Println("Введіть адресу відправника:")
	fmt.Scanln(&c.Giver)
	return c.Giver
}

func (c *Convert) SaveReciver() string {
	fmt.Println("Введіть адресу отримувача:")
	fmt.Scanln(&c.Reciver)
	return c.Reciver
}

type Sortirovanie struct{}

func (s *Sortirovanie) Sortiv(p Posilka) {
	switch p.(type) {
	case *Box:
		fmt.Println("Сортувальний відділ: відправити коробку")
	case *Convert:
		fmt.Println("Сортувальний відділ: відправити конверт")
	default:
		fmt.Println("Сортувальний відділ: тип посилки не визначено")
	}
}

type Shipment struct {
	Type        string
	SenderAdr   string
	ReceiverAdr string
}

func ChoosePosillka() {
	var shipments []Shipment
	for {
		var senderAdr, receiverAdr, posilkaType string
		fmt.Println("Виберіть будь ласка тип відправки. Коробка чи конверт: ")
		fmt.Scanln(&posilkaType)

		if posilkaType != "конверт" && posilkaType != "коробка" {
			fmt.Println("Будь ласка, Оберіть 'конверт' чи 'коробка'.")
			continue
		}

		fmt.Print("Введіть будь ласка адресу відправника: ")
		fmt.Scanln(&senderAdr)

		fmt.Print("Введіть будь ласка адресу отримувача: ")
		fmt.Scanln(&receiverAdr)

		shipment := Shipment{
			Type:        posilkaType,
			SenderAdr:   senderAdr,
			ReceiverAdr: receiverAdr,
		}
		shipments = append(shipments, shipment)

		fmt.Println("Інформація об відправленнях:")
		for _, shipment := range shipments {
			fmt.Printf("Тип: %s, Відправник: %s, Отримувач: %s\n", shipment.Type, shipment.SenderAdr, shipment.ReceiverAdr)

		}
		break
	}
	fmt.Println("Програма завершилась.")
}

func main() {
	box := Box{Giver: "Відправник 1", Reciver: "Отримувач 1"}
	convert := Convert{Giver: "Відправник 2", Reciver: "Отримувач 2"}
	var s Sortirovanie
	s.Sortiv(&box)
	s.Sortiv(&convert)
	ChoosePosillka()

}
