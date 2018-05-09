/*Package pagination provides a convenient way to set up a pagination range with data for use in a "html/template".
It places the current page on a defined position (while preventing overflows). There is a maximum size of the pagination, to fit your template design when the result set grows. For convenience it provides all neccesary data, like next, prev and current page easily called from the pagination methods inside the template.

Usage:

	a := pagination.Args{
			Max:     9,    //Maximum amount of pagination entries
			Pos:     3,    //Position of active page
			Page:    7,    //Current page
			Records: 5,    //Current results
			Total:   499,  //Total amount of records
			Size:    5,    //Records per page
	}

Naturally, "Pos" should always be smaller than "Max".
"Page" should always be smaller than (Total / Size = Pages), rounded up.
"Records" can never be bigger than "Size" and "Total".
"Records" is usually equal to Size, but can be smaller on the last page in practice. "Records" value is not used in any calculation and can be ommited if not needed in the template for example statistics.

All the constraints are checked during creation of the new pagination object:

	pag, err := pagination.New(a)
	if err != nil {
		if err.Error() == pagination.ErrPageNo {
			//Invalid request from client, requested page number to high
			return
		}
		//One of the remaining constraints are not met
		return
	}

Now the Pagination object pointer can be incorperated in the data structure passed to a template
	view := Page{
		Articles:   results,
		Pagination: pag,
	}
	if err = templates.ExecuteTemplate(w, "layout", view); err != nil {
		//Do something
		return
	}
Let's say you use a seperate template for pagination, you can call it from a layout template like:
	{{template "pagination" .Pagination}}

Finally, inside the pagination template simply call the methods where you need them. ".Entries" gives the actual pagination range. Every "Entry" has two fields, "Active" and "Number":

"Active" is set to boolean true if that Entry represents the current page. False for all others. So a direct test can be done unsing {{if .Active}}.

"Number" is set to the number the pagination entry is referring to.

All the other methods just print a number for statistical puposes.

	{{define "pagination"}}
		<!--This example uses bootstrap pagination classes-->
		<ul class="pagination">
		{{- if eq .Prev 0}}
			<li class="page-item disabled"><a class="page-link" >&lt;&lt;&nbsp;First</a></li>
		{{- else}}
			<li class="page-item"><a class="page-link" href=".">&lt;&lt;&nbsp;First</a></li>
			<li class="page-item"><a class="page-link" rel="prev" href="?page={{.Prev}}">&lt;&nbsp;Previous</a></li>
		{{end}}
		{{- range .Entries}}
			<li class="page-item{{if .Active}} active{{end}}"><a class="page-link" href="?page={{.Number}}">{{.Number}}</a></li>
		{{- end}}
		{{- if eq .Next 0}}
			<li class="page-item disabled"><a class="page-link" >last&nbsp;&gt;&gt;</a></li>
		{{- else}}
			<li class="page-item"><a class="page-link" rel="next" href="?page={{.Next}}">Next&nbsp;&gt;</a></li>
			<li class="page-item"><a class="page-link" href="/?page={{.Pages}}">Last&nbsp;&gt;&gt;</a></li>
		{{- end}}
		</ul>
		<p>Page {{.Page}} of {{.Pages}}, showing {{.Records}} record(s) out of {{.Total}} total</p>
	{{end}}
*/
package pagination
