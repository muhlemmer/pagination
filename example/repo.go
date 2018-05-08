package main

import (
	"errors"
)

type row string

func (r row) Title() string {
	return string(r[:15]) + "..."
}

func (r row) Desc() string {
	return string(r)
}

type query struct {
	Page  int
	Limit int
}

func (q query) Offset() int {
	return (q.Page - 1) * q.Limit // (38 - 1) * 4 = 37 * 4 = 148
}

type db []row

type result struct {
	Title   string
	Content string
}

func (d db) list(offset int, limit int) (res []result, err error) {
	if offset >= len(d) || offset < 0 {
		err = errors.New("Db list: limit outside range")
		return
	}
	if offset+limit > len(d) { // 6 + 4 = 10		7 + 4 = 11		0 + 4 = 0
		limit = len(d) - offset //10 - 6  = 4		10 - 7 = 3
	}

	res = make([]result, limit)
	for i := 0; i < limit; i++ {
		r := d[offset+i]
		res[i] = result{
			Title:   r.Title(),
			Content: r.Desc(),
		}
	}
	return
}

func (d db) Query(q query) (res []result, err error) {
	return d.list(q.Offset(), q.Limit)
}

func (d db) Count() int {
	return len(d)
}
