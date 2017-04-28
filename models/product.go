package models

import (
	"fmt"
	"github.com/kataras/go-errors"
	"time"
)

const ProductTable = "products"

type Product struct {
	Id int64
	Name string
	Url string
	TargetPrice string
	CreateDate time.Time
	status string
}

func (p Product) FindByStatus(status string) (*[]Product, error) {
	if status == "" {
		return nil, errors.New("Please provide a status")
	}

	sql := fmt.Sprintf("select `id`, `name`, `url`, `target_price`, `create_date`, `status` from %s where `status`=?", ProductTable)
	rows, err := db.Query(sql, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.Id, &p.Name, &p.Url, &p.TargetPrice, &p.CreateDate, &p.status)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return &products, nil
}

func (p Product) Save() (*Product, error) {
	var sql string
	if p.Id == 0 {
		p.CreateDate = time.Now()
		p.CreateDate.Format(time.RFC3339)
		sql = fmt.Sprintf("INSERT INTO %s (`name`, `url`, `target_price`, `create_date`, `status` VALUES (?,?,?,?,?)", ProductTable)
	} else {
		sql = fmt.Sprintf("UPDATE %s SET `name`=?, `url`=?, `target_price`=?, `create_date`=?, `status`=? WHERE `id`=%d", ProductTable, p.Id)
	}

	res, err := db.Exec(sql, p.Name, p.Url, p.TargetPrice, p.CreateDate, p.status)
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