package pagination

import (
	"errors"
	//"fmt"
)

type Entry struct {
	Active bool
	Number int
}

type Args struct {
	Max     int //Maximum amount of pagination entries
	Pos     int //Position of active page
	Page    int //Current page
	Results int //Current results
	Records int //Total amount of records
	Size    int //Records per page
}

type Pagination struct {
	Prev    int
	Page    int //Current page
	Next    int
	Results int     //Current results
	Records int     //Total amount of records
	Size    int     //Records per page
	Pages   int     //Total amount of pages.
	Entries []Entry //Entry range for template
}

const (
	errBase       = "Pagination: "
	errInvSize    = errBase + "Size <= 0"
	ErrPageNo     = errBase + "Page > Pages"
	errResultSize = errBase + "Results > Size"
	errRecordSize = errBase + "Results > Records"
)

func (a *Args) pages() (p int, err error) {
	p = (a.Records-1)/a.Size + 1
	if a.Page > p {
		err = errors.New(ErrPageNo)
	}
	return
}

func (a *Args) check() (err error) {
	if a.Size <= 0 {
		err = errors.New(errInvSize)
		return
	}
	if a.Results > a.Size {
		err = errors.New(errResultSize)
		return
	}
	if a.Results > a.Records {
		err = errors.New(errRecordSize)
	}
	return
}

func (p *Pagination) prev() {
	if p.Page <= 1 {
		return //p.Prev will remain 0
	}
	p.Prev = p.Page - 1
	return
}

func (p *Pagination) next() {
	if p.Page == p.Pages {
		return //p.Next will remain 0
	}
	p.Next = p.Page + 1
	return
}

func (p *Pagination) entries(a *Args) {
	//sn is the start page number of the entries range
	sn := p.Page - a.Pos //13 - 3 = 10
	switch {
	case sn < 0: //Don't show negative page numbers.
		sn = 0
		break
	case a.Max-a.Pos+p.Page > p.Pages: //9-3+13 = 19 > 21 = false
		sn = p.Pages - a.Max
	}
	sn++ //start with number 1

	p.Entries = nil
	p.Entries = make([]Entry, a.Max)
	for i := 0; i < a.Max; i++ {
		var e Entry
		e.Number = sn + i
		e.Active = e.Number == p.Page
		p.Entries[i] = e
	}
	//fmt.Println(p.Entries)
	return
}

//Current page, No. results, Total no. records, Page size
func New(a Args) (p *Pagination, err error) {
	if err = a.check(); err != nil {
		return
	}
	pages, err := a.pages()
	if err != nil {
		return
	}

	p = &Pagination{
		Page:    a.Page,
		Results: a.Results,
		Records: a.Records,
		Size:    a.Size,
		Pages:   pages,
	}
	p.prev()
	p.next()

	//Determine the needed amount of entries
	if pages < a.Max {
		a.Max = pages
	}
	p.entries(&a)
	return
}
