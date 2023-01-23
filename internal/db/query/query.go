package query

import (
	"database/sql"
	"log"
	"strings"
)

type Orders struct {
	List []Item
}

type Item struct {
	Name        string
	OrderNumber string
	FileUrl     string
	Amount      string
	MainRack    string
	OnRack      string
	AnyRack     []string
}

func SelectOrders(db *sql.DB, count []string) (*Orders, bool) {

	param := "{" + strings.Join(count, ",") + "}"

	rows, err := db.Query(`SELECT order_number,items.name,items.file_url,amount,storage.rack_name,main_rack FROM client_order, items, storage WHERE client_order.item_id = items.id AND storage.item_id = items.id AND client_order.order_number = ANY($1::int[])`, param)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	defer rows.Close()

	orders := Orders{List: nil}

	for rows.Next() {
		item := Item{}

		err := rows.Scan(&item.OrderNumber, &item.Name, &item.FileUrl, &item.Amount, &item.OnRack, &item.MainRack)
		if err != nil {
			log.Println(err)
			continue
		}
		orders.List = append(orders.List, item)

	}

	if len(orders.List) == 0 {
		return nil, false
	}

	return &orders, true
}
