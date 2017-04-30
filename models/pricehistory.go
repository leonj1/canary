package models

import (
	"fmt"
	"github.com/kataras/go-errors"
	"time"
)

const PriceHistoryTable = "price_history"

type PriceHistory struct {
	Id int64
	ProductId int64
	Price string
	AlternatePrice string
	Seller string
	CreateDate time.Time
}

func (p PriceHistory) Save() (*PriceHistory, error) {
	var sql string
	if p.Id == 0 {
		p.CreateDate = time.Now()
		p.CreateDate.Format(time.RFC3339)
		sql = fmt.Sprintf("INSERT INTO %s (`product_id`, `price`, `alternate_price`, `seller`, `create_date`) VALUES (?,?,?,?,?)", PriceHistoryTable)
	} else {
		sql = fmt.Sprintf("UPDATE %s SET `product_id`=?, `price`=?, `alternate_price`=?, `seller`=?, `create_date`=? WHERE `id`=%d", PriceHistoryTable, p.Id)
	}

	res, err := db.Exec(sql, p.ProductId, p.Price, p.AlternatePrice, p.Seller, p.CreateDate)
	if err != nil {
		return nil, err
	}

	if p.Id == 0 {
		p.Id, err = res.LastInsertId()
		if err != nil {
			return nil, err
		}
	}

	return &p, nil
 }

func (p PriceHistory) FindByProductId(id int64) (*[]PriceHistory, error) {
	if id == 0 {
		return nil, errors.New("Please provide an id")
	}

	sql := fmt.Sprintf("select `id`, `product_id`, `price`, `alternate_price`, `seller`, `create_date` from %s where `product_id`=?", PriceHistoryTable)
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priceHistories []PriceHistory
	for rows.Next() {
		p := new(PriceHistory)
		err := rows.Scan(&p.Id, &p.ProductId, &p.Price, &p.AlternatePrice, &p.Seller, &p.CreateDate)
		if err != nil {
			return nil, err
		}
		priceHistories = append(priceHistories, *p)
	}

	return &priceHistories, nil
}
