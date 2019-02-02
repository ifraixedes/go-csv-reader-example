# Go Example CSV Reader

This implementation was created to fulfill the requirements that you can find in the [_requirements.md file_](requirements.md), in around 8 hours and with the purpose of containing good practices like testing, decent git commit history and a decent source documentation.

## Implementation

The implementation is split in two packages. A `main` package (in the root) and the `stats` package which is in subfolder named _stats_.

The `main` package is the command line tool, the binary; it doesn't have any logic rather than the one related to parse the input parameters, open the CSV file to be analyzed and to pretty print the result to the `stdout`, a part of printing any error which could happen to the `stderr`. For knowing which command line arguments the tool accepts, run the binary with the `-h` argument.

The `stats` package is the package which has all the types and functions to perform the required operations/computations. All the exported members are documented using the Go doc conventions, so I invite you to read them if you want/need more thorough information of each one.

Below there are some points which give a general and brief description on what you will find in the `stats` package:

* An interface which represents the methods of the `csv.Reader` type; it's used for having an abstraction of it and having specific implementations of such type.
* A type which satisfies the `csv.Reader` interface whose methods behave like the `csv.Reader` but only acts on records which are inside of specified time window.
* A type which represents the required stats of the _cloud remote builder service_ and a function which compute them, from an input CSV file reader in a specified time window.
* Other types, which are helpful for parsing each record of the determined _cloud remote builder service_ CSV output file.

### Tests

All the implemented tests are _"black box tests"_, which means that there is not test of any unexported function, type, type field, etc.; this is achieved using the package name with the `_test` suffix. Hence some of them are integrations tests, others could be considered unit tests if the item under test doesn't involve others, but even with that, if the are not stateless, its internal state isn't verified unless that it be exposed through an exported method, nonetheless, it doesn't mean that the tests aren't thorough enough.

The advantage of having _"black box tests"_ is that they require less maintenance on future changes, obviously if the aren't breaking changes on the exported interface, types, etc. This doesn't mean that _"white box tests"_ (unit tests) are useless, there are several cases where they are needed and are very helpful, but for this implementation and considering the time which I had, I chose to only have the integrations ones.

I believe that the tests are quite clear by itself, hence I invite you to read them for having more insights on what's test it and what may not.

## Improvements

This section include some improvements which could be done having a bit more time:

1. Implement the tests which are marked with `testing.T.Skip` function.
2. Have test fixtures with corner cases and have tests which use them to ensure that the implementation is resilient to already known corner cases.
3. The time window reader constructor (`NewTimeWindowReader`) could accept one more parameter for specifying which field index contains the time value for finding out if the record is in the input time windows. This will make this component more flexible and will allow to be used for filtering records based in other time fields, like the request time, etc.
4. Although I could assume that the records of the CSV file are sorted by date from older to newer, I opted for not doing such assumption and providing a more robust solution, because it works with CSV which are sorted and unsorted; however, if we could assume so, the reader returned by the `NewTimeWindowReader` constructor function could be more efficient, just stopping on the first record whose date is more recent than the upper limit date of the time window, without having to iterate all the records until the last one.
5. Add a proper help message of the command line tool (`main`) to inform to the user what this tool does.
5. Command line tool could accepts time windows in more human format, like "1 day ago", "last week", etc., for easing its usage to the user.


On the other hand, many other improvements could be done having an exhaustive information of the stakeholders' requirements and more knowledge about the business domain, not only in terms of features (e.g. more stats calculations), but in terms of optimizing the calculations for the different stats calculations for having less iterations and with so better performance; nonetheless, the mentioned performance optimizations should be thought and deeply evaluated, because they will probably require a more complex implementation with the trade-offs of having a more difficulty to understand and maintain  it.

## License

MIT, read [the license file](LICENSE) for more information.
