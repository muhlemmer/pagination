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

Run ````go get github.com/muhlemmer/pagination````

### Prerequisites

This package has been developed with Go version 1.10. However, no extremely new features have been used, by best knowledge of the author. So it might work on older version as well.

### Usage

A working example can be found in the `example/` folder.

----All content below this line has not yet been edited from the template.----

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc
