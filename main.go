package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"test/internal/config"
	"test/internal/db"
	"test/internal/db/query"
)

func main() {

	Config := config.GetConfig()
	dbConnection := db.ConnectDb(Config)
	defer dbConnection.Close()

	var count []string
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите номера заказов по одному")
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		count = append(count, sc.Text())
	}

	orders, ok := query.SelectOrders(dbConnection, count)
	if !ok {
		log.Fatal("Ошибка запроса")
		return
	}
	orderMap := map[string][]query.Item{}
	var orderLeft []query.Item
	for _, v := range orders.List {
		if v.OnRack == v.MainRack {
			orderMap[v.OnRack] = append(orderMap[v.OnRack], v)
		} else {
			orderLeft = append(orderLeft, v)
		}
	}

	for _, v := range orderLeft {
		if _, ok := orderMap[v.OnRack]; !ok {
			if _, ok := orderMap[v.MainRack]; ok {
				for i, v2 := range orderMap[v.MainRack] {
					if v.Name == v2.Name {
						orderMap[v.MainRack][i].AnyRack = append(orderMap[v.MainRack][i].AnyRack, v.OnRack)
					} else {
						orderMap[v.OnRack] = append(orderMap[v.OnRack], v)
					}
				}
			}
		} else {
			if _, ok := orderMap[v.MainRack]; ok {
				for i, v2 := range orderMap[v.MainRack] {
					if v.Name == v2.Name {
						orderMap[v.MainRack][i].AnyRack = append(orderMap[v.MainRack][i].AnyRack, v.OnRack)
					}
				}
			}

		}
	}
	for i, v := range orderMap {
		fmt.Println("\nСтеллаж: " + i)
		for _, v2 := range v {
			fmt.Println("Заказ: " + v2.OrderNumber + " Предмет: " + v2.Name + " " + v2.Amount + " шт")
			if v2.AnyRack != nil {
				fmt.Print("доп стеллаж: ")
				for i3 := range v2.AnyRack {
					fmt.Print(v2.AnyRack[i3], " ")
				}
				fmt.Print("\n")
			}
		}
	}
}
