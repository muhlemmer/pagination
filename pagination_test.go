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
	r *Pagination
}

var newTests = []testParam{
	{
		a: Args{5, 3, 1, 10, 100, 10},
		r: &Pagination{
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
		r: &Pagination{
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
		r: &Pagination{
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
		r: &Pagination{
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
		r: &Pagination{
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
		r: &Pagination{
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
		r: &Pagination{
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

func TestNew(t *testing.T) {
	for _, tp := range newTests {
		p, err := New(tp.a)
		if err != nil {
			t.Error(err.Error())
			return
		}

		fail := false
		if tp.r.Prev != p.Prev || tp.r.Page != p.Page || tp.r.Next != p.Next || tp.r.Results != p.Results || tp.r.Records != p.Records || tp.r.Size != p.Size || tp.r.Pages != p.Pages {
			fail = true
		}
		if len(tp.r.Entries) == len(p.Entries) {
			for k, e := range tp.r.Entries {
				if e != p.Entries[k] {
					fail = true
				}
			}
		} else {
			fail = true
		}
		if fail {
			t.Error(
				"\nTestNew for: ", tp.a,
				"\nExpected:    ", tp.r,
				"\nGot:         ", p,
			)
		}
	}
}
