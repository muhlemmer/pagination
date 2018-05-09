# Pagination
Pagination is a small Go library for setting up a range with data for use in a html/template. Features:
* Whenever possible, current page at a fixed offset. So that there is consistent placement among all pages. The offset is shifted at the beginning or end of the page range.
* Max amount of pages to show, to be sure the pagination fits you template.
* Convenience of all applicable numbers from the pagination object: no need to access all kind of data structures and arithmetic from inside the template
  * Next, Prev, Current methods.
  * List statistics 

The exported methods of the *Pagination* type can be called directly from the template and aim to be guaranteed to run from the template. Sanity checks are done early, when creating the new *Pagination* object. Hence, no error should occur at template execution time. (No excessive testing has been done yet, this just reflects the project aim)

## ToDo
1. Package summary documentation for godoc
2. Travis-ci integration
3. Goreport card integration
4. More real life testing

## Getting Started

````Shell
go get github.com/muhlemmer/pagination
````

### Prerequisites

This package has been developed with Go version 1.10.1 However, no extremely new features have been used, by best knowledge of the author. So it might work on older version as well.

### Usage

A working example can be found in the `example/` folder. Godoc is still in progress and the link will be included here when done. In principle there are 4 simple steps:
1. Import the library
2. Create the `pagination.Args` object after you collect the necessary data from queries etc.
3. Call `pagination.New(Args)` which will do a sanity check on the supplied data and return a pointer to a new pagination object.
4. Pass the pointer to your template and call the methods from there.

## Running the tests

The test suite is provided in the `pagination_test.go` file.  The last range of tests are confirming stable output.

Test file will be build and run by:
````Shell
go test
````

### Error testing

The first range of tests are using bogus values to trigger all the errors in sanity checking. All the error test function start with `func TestErr*`

### Output testing

The final test function tests against various data sets, to ensure predictable output. These test are done in the final `Test(T)` function.

## Built With

* [Go](https://golang.org) – Version 1.10.1
* [Bootstrap](http://getbootstrap.com/) - Used in the example template
* [Lorem Ipsum](https://www.lipsum.com/) - Used as data in the example

## Contributing

Feel free to send any *pull requests*. Please format your code with `gofmt` before doing so.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Tim Möhlmann** - *Initial work* - [Muhlemmer](https://github.com/muhlemmer)

See also the list of [contributors](https://github.com/muhlemmer/pagination/contributors) who participated in this project.

## License

This project is licensed under the BSD 3-Clause - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This is the authors first open sourced library ever written and was originally part of a small dedicated web application. Writing this software and its documentation was fun ;). Hopefully this will be useful for others out there and any improvements are more then welcome!
