package pagination

import (
	"errors"
)

type Entry struct {
	Active bool
	Number int
}

type Args struct {
	Max     int //Maximum amount of pagination entries
	Pos     int //Position of active page
	Page    int //Current page
	Records int //Current results
	Total   int //Total amount of records
	Size    int //Records per page
}

const (
	errBase       = "Pagination: "
	errInvSize    = errBase + "Size <= 0"
	ErrPageNo     = errBase + "Page > Pages"
	errResultSize = errBase + "Results > Size"
	errRecordSize = errBase + "Results > Records"
)

func (a *Args) pages() (p int, err error) {
	p = (a.Total-1)/a.Size + 1
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
	if a.Records > a.Size {
		err = errors.New(errResultSize)
		return
	}
	if a.Records > a.Total {
		err = errors.New(errRecordSize)
	}
	return
}

type Pagination struct {
	pages int
	args  *Args
}

//Current page, No. results, Total no. records, Page size
func New(a Args) (pag *Pagination, err error) {
	if err = a.check(); err != nil {
		return
	}
	p, err := a.pages()
	if err != nil {
		return
	}

	//Determine the needed amount of entries
	if p < a.Max {
		a.Max = p
	}

	pag = &Pagination{
		pages: p,
		args:  &a,
	}
	return
}

func (p *Pagination) Prev() int {
	if p.args.Page <= 1 {
		return 0
	}
	return p.args.Page - 1
}

func (p *Pagination) Page() int {
	return p.args.Page
}

func (p *Pagination) Next() int {
	if p.args.Page == p.pages {
		return 0
	}
	return p.args.Page + 1
}

func (p *Pagination) Records() int {
	return p.args.Records
}

func (p *Pagination) Total() int {
	return p.args.Total
}

func (p *Pagination) Size() int {
	return p.args.Size
}

func (p *Pagination) Pages() int {
	return p.pages
}

func (p *Pagination) Entries() (r []Entry) {
	//sn is the start page number of the entries range
	sn := p.args.Page - p.args.Pos
	switch {
	case sn < 0: //Don't show negative page numbers.
		sn = 0
		break
	case p.args.Max-p.args.Pos+p.args.Page > p.pages:
		sn = p.pages - p.args.Max
	}
	sn++ //start with number 1

	r = make([]Entry, p.args.Max)
	for i := 0; i < p.args.Max; i++ {
		var e Entry
		e.Number = sn + i
		e.Active = e.Number == p.args.Page
		r[i] = e
	}
	return
}
