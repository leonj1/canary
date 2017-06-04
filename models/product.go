package models

import (
	"fmt"
	"github.com/kataras/go-errors"
	"time"
)

const ProductTable = "products"

type Product struct {
	Id 		int64		`json:"id,string,omitempty"`
	Name 		string		`json:"name,omitempty"`
	Url 		string		`json:"url,omitempty"`
	TargetPrice 	string		`json:"target_price,omitempty"`
	CurrentPrice 	string		`json:"current_price,omitempty"`
	CreateDate 	time.Time	`json:"create_date,string,omitempty"`
	Status 		string		`json:"status,omitempty"`
	Website 	string		`json:"website,omitempty"`
}

func (p Product) FindByStatus(status string) ([]*Product, error) {
	if status == "" {
		return nil, errors.New("Please provide a status")
	}

	sql := fmt.Sprintf("select `id`, `name`, `url`, `target_price`, `create_date`, `status`, `website` from %s where `status`=?", ProductTable)
	rows, err := db.Query(sql, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.Id, &p.Name, &p.Url, &p.TargetPrice, &p.CreateDate, &p.Status, &p.Website)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (p Product) Save() (*Product, error) {
	var sql string
	if p.Id == 0 {
		p.CreateDate = time.Now()
		p.CreateDate.Format(time.RFC3339)
		p.Status = "ACTIVE"
		sql = fmt.Sprintf("INSERT INTO %s (`name`, `url`, `target_price`, `create_date`, `status`, `website`) VALUES (?,?,?,?,?,?)", ProductTable)
	} else {
		sql = fmt.Sprintf("UPDATE %s SET `name`=?, `url`=?, `target_price`=?, `create_date`=?, `status`=?, `website`=? WHERE `id`=%d", ProductTable, p.Id)
	}

	res, err := db.Exec(sql, p.Name, p.Url, p.TargetPrice, p.CreateDate, p.Status, p.Website)
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

func (p Product) FindByProductId(id int64) (*[]Product, error) {
	if id == 0 {
		return nil, errors.New("Please provide an id")
	}

	sql := fmt.Sprintf("select `id`, `name`, `url`, `target_price`, `create_date`, `status`, `website` from %s where `id`=?", ProductTable)
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.Id, &p.Name, &p.Url, &p.TargetPrice, &p.CreateDate, &p.Status, &p.Website)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return &products, nil
}

