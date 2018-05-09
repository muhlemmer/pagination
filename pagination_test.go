// Copyright 2018 Tim Möhlmann. All rights reserved.
// This project is licensed under the BSD 3-Clause
// See the LICENSE file for details.

package pagination

import (
	"testing"
)

var errInvSizeTests = []Args{
	{5, 3, 1, 10, 100, 0},
	{5, 3, 1, 10, 100, -1},
}

func TestErrInvSize(t *testing.T) {
	for _, a := range errInvSizeTests {
		err := a.check()
		if err == nil {
			t.Error(
				"For: ", a,
				"; Expected: ", errInvSize,
				"; Got: ", "nil",
			)
			continue
		}
		if err.Error() != errInvSize {
			t.Error(
				"For: ", a,
				"; Expected: ", errInvSize,
				"; Got: ", err.Error(),
			)
		}
	}
}

var errResultSizeTests = []Args{
	{5, 3, 1, 11, 100, 10},
	{5, 3, 1, 77, 100, 50},
}

func TestErrResultSize(t *testing.T) {
	for _, a := range errResultSizeTests {
		err := a.check()
		if err == nil {
			t.Error(
				"For: ", a,
				"; Expected: ", errResultSize,
				"; Got: ", "nil",
			)
			continue
		}
		if err.Error() != errResultSize {
			t.Error(
				"For: ", a,
				"; Expected: ", errResultSize,
				"; Got: ", err.Error(),
			)
		}
	}
}

var errRecordSizeTests = []Args{
	{5, 3, 1, 399, 200, 400},
	{5, 3, 1, 101, 5, 300},
}

func TestErrRecordSize(t *testing.T) {
	for _, a := range errRecordSizeTests {
		err := a.check()
		if err == nil {
			t.Error(
				"For: ", a,
				"; Expected: ", errRecordSize,
				"; Got: ", "nil",
			)
			continue
		}
		if err.Error() != errRecordSize {
			t.Error(
				"For: ", a,
				"; Expected: ", errRecordSize,
				"; Got: ", err.Error(),
			)
		}
	}
}

var pagesErrorTests = []Args{
	{5, 3, 11, 10, 100, 10},
	{5, 3, 2, 10, 100, 101},
}

func TestPagesError(t *testing.T) {
	for _, a := range pagesErrorTests {
		_, err := a.pages()
		if err == nil {
			t.Error(
				"For: ", a,
				"; Expected: ", ErrPageNo,
				"; Got: ", "nil",
			)
			continue
		}
		if err.Error() != ErrPageNo {
			t.Error(
				"For: ", a,
				"; Expected: ", ErrPageNo,
				"; Got: ", err.Error(),
			)
		}
	}
}

type testParam struct {
	a Args
	r testResult
}

type testResult struct {
	Prev    int     // Previous page
	Page    int     //Current page
	Next    int     // Next page
	Records int     //Current results
	Total   int     //Total amount of records
	Size    int     //Records per page
	Pages   int     //Total amount of pages.
	Entries []Entry //Entry range for template
}

var tests = []testParam{
	{
		a: Args{5, 3, 1, 10, 100, 10},
		r: testResult{
			0, 1, 2, 10, 100, 10, 10,
			[]Entry{
				{true, 1},
				{false, 2},
				{false, 3},
				{false, 4},
				{false, 5},
			},
		},
	},
	{
		a: Args{5, 3, 10, 10, 100, 10},
		r: testResult{
			9, 10, 0, 10, 100, 10, 10,
			[]Entry{
				{false, 6},
				{false, 7},
				{false, 8},
				{false, 9},
				{true, 10},
			},
		},
	},
	{
		a: Args{5, 3, 5, 10, 100, 10},
		r: testResult{
			4, 5, 6, 10, 100, 10, 10,
			[]Entry{
				{false, 3},
				{false, 4},
				{true, 5},
				{false, 6},
				{false, 7},
			},
		},
	},
	{
		a: Args{5, 3, 0, 10, 100, 10},
		r: testResult{
			0, 0, 1, 10, 100, 10, 10,
			[]Entry{
				{false, 1},
				{false, 2},
				{false, 3},
				{false, 4},
				{false, 5},
			},
		},
	},
	{
		a: Args{5, 3, 22, 30, 5000, 30},
		r: testResult{
			21, 22, 23, 30, 5000, 30, 167,
			[]Entry{
				{false, 20},
				{false, 21},
				{true, 22},
				{false, 23},
				{false, 24},
			},
		},
	},
	{
		a: Args{5, 3, 9, 10, 100, 10},
		r: testResult{
			8, 9, 10, 10, 100, 10, 10,
			[]Entry{
				{false, 6},
				{false, 7},
				{false, 8},
				{true, 9},
				{false, 10},
			},
		},
	},
	{
		a: Args{9, 3, 13, 27, 550, 27},
		r: testResult{
			12, 13, 14, 27, 550, 27, 21,
			[]Entry{
				{false, 11},
				{false, 12},
				{true, 13},
				{false, 14},
				{false, 15},
				{false, 16},
				{false, 17},
				{false, 18},
				{false, 19},
			},
		},
	},
}

func Test(t *testing.T) {
	for _, tp := range tests {
		p, err := New(tp.a)
		if err != nil {
			t.Error("New() for: ", tp.a, " Error: ", err.Error())
			continue
		}

		if p.Prev() != tp.r.Prev {
			t.Error("Prev() for: ", tp.a, " Expected: ", tp.r.Prev, " Got: ", p.Prev())
		}

		if p.Page() != tp.r.Page {
			t.Error("Page() for: ", tp.a, " Expected: ", tp.r.Page, " Got: ", p.Page())
		}

		if p.Next() != tp.r.Next {
			t.Error("Next() for: ", tp.a, " Expected: ", tp.r.Next, " Got: ", p.Next())
		}

		if p.Records() != tp.r.Records {
			t.Error("Records() for: ", tp.a, " Expected: ", tp.r.Records, " Got: ", p.Records())
		}

		if p.Total() != tp.r.Total {
			t.Error("Total() for: ", tp.a, " Expected: ", tp.r.Total, " Got: ", p.Total())
		}

		if p.Size() != tp.r.Size {
			t.Error("Size() for: ", tp.a, " Expected: ", tp.r.Size, " Got: ", p.Size())
		}

		if p.Pages() != tp.r.Pages {
			t.Error("Pages() for: ", tp.a, " Expected: ", tp.r.Pages, " Got: ", p.Pages())
		}

		entries := p.Entries()
		if len(entries) != len(tp.r.Entries) {
			t.Error("len(Entries()) for: ", tp.a, " Expected: ", len(tp.r.Entries), " Got: ", len(entries))
			continue
		}

		for k, e := range entries {
			if e != tp.r.Entries[k] {
				t.Error("Entries() for ", tp.a, " Expected: ", tp.r.Entries, " Got: ", entries)
				break
			}
		}
	}
}
