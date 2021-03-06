// Copyright 2018 Tim Möhlmann. All rights reserved.
// This project is licensed under the BSD 3-Clause
// See the LICENSE file for details.

package pagination

import (
	"errors"
)

//Args contains the arguments for constructing a New pagination object.
type Args struct {
	Max     int //Maximum amount of pagination entries
	Pos     int //Position of active page
	Page    int //Current page
	Records int //Current results (optional)
	Total   int //Total amount of records
	Size    int //Records per page
}

const (
	//ErrPageNo is a string constant used in the error returned by pagination.New() whenever the Page number is invalid.
	ErrPageNo     = errBase + "Page > Pages"
	errBase       = "Error in pagination: "
	errInvSize    = errBase + "Size <= 0"
	errResultSize = errBase + "Results > Size"
	errRecordSize = errBase + "Results > Records"
)

// pages calculates the amount of pages, based on the total amount of records and pages size.
// Its result is always rounded up.
// An error is returned if the current page is larger then pages.
func (a *Args) pages() (p int, err error) {
	p = (a.Total-1)/a.Size + 1
	if a.Page > p {
		err = errors.New(ErrPageNo)
	}
	return
}

// check does a sanity check on provided values
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

// Pagination holds the pagination methods. This object can be passed directly to a html/template.
// The methods defiend on the object can be called directly from the template.
type Pagination struct {
	pages int
	args  *Args
}

// New creates a new pagination object and return a pointer to it. This method performs some sanity checks on the Args data and returns a nil pointer and an error if a bogus value is supplied.
//
// See ErrPageNo for invalid page number.
// For any other error message, something probably went wrong in the calling code, DB query etc.
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

// Prev returns the number of the previous page. It returns 0 if there is no previous page.
func (p *Pagination) Prev() int {
	if p.args.Page <= 1 {
		return 0
	}
	return p.args.Page - 1
}

// Page is a getter wrapper for Args.Page. It returns the current page.
func (p *Pagination) Page() int {
	return p.args.Page
}

// Next returns the number of the next page. It returns 0 if there is no next page.
func (p *Pagination) Next() int {
	if p.args.Page == p.pages {
		return 0
	}
	return p.args.Page + 1
}

// Records is a getter wrapper for Args.Records. It returns the current amount of records.
func (p *Pagination) Records() int {
	return p.args.Records
}

// Total is a getter wrapper for Args.Total. It returns the total amount of records.
func (p *Pagination) Total() int {
	return p.args.Total
}

// Size is a getter wrapper for Args.Size. It returns the page size.
func (p *Pagination) Size() int {
	return p.args.Size
}

// Pages returns the total number of calculated pages, based on Args.Total and Args.Size.
// The number of pages is always rounded up.
func (p *Pagination) Pages() int {
	return p.pages
}

// Entry represents a page number in the pagination range. The active page has the "Active" field set to "true".
type Entry struct {
	Active bool // true for the current page, false for any other
	Number int  // The page number this entry is representing.
}

// Entries returns a slice of Entry, over which can be ranged inside the template.
func (p *Pagination) Entries() (r []Entry) {
	sn := p.args.Page - p.args.Pos //sn is the start page number of the entries range
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
