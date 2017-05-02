package models

import (
	"fmt"
	"time"
)

const ExecutionTable = "executions"

type Execution struct {
	Id 		int64		`json:"id,string,omitempty"`
	CreateDate 	time.Time	`json:"create_date,string,omitempty"`
}

func (p *Execution) Save() (*Execution, error) {
	var sql string
	if p.Id == 0 {
		p.CreateDate = time.Now()
		p.CreateDate.Format(time.RFC3339)
		sql = fmt.Sprintf("INSERT INTO %s (`create_date`) VALUES (?)", ExecutionTable)
	} else {
		sql = fmt.Sprintf("UPDATE %s SET `create_date`=? WHERE `id`=%d", ExecutionTable, p.Id)
	}

	res, err := db.Exec(sql, p.CreateDate)
	if err != nil {
		return nil, err
	}

	if p.Id == 0 {
		p.Id, err = res.LastInsertId()
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p Execution) FindAll() ([]*Execution, error) {
	sql := fmt.Sprintf("select `id`, `create_date` from %s", ExecutionTable)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var executions []*Execution
	for rows.Next() {
		p := new(Execution)
		err := rows.Scan(&p.Id, &p.CreateDate)
		if err != nil {
			return nil, err
		}
		executions = append(executions, p)
	}

	return executions, nil
}

