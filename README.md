# http-crawler

# Instructions
You should have golang installed of version that supports go modules.
- to run tests use `go test -race ./...` (or `make test` if you have make installed)
- to build binary use `go build -race -o ./bin/http-crawler github.com/itimofeev/http-crawler` (or `make build`)
- to run with some example arguments use `go run -race github.com/itimofeev/http-crawler -parallel 10 google.com yandex.ru https://adjust.com http://www.reddit.com/r/funny/ www.reddit.com/r/notfunny https://baroquemusiclibrary.com/` (or `make run`)

This is quite simple implementation, that I tried to keep as simple as possible.
For production code there may be done some improvements, more correct error handling, adding linters, and so on.

## Task
You must build a tool which makes http requests and prints the address of the request along with the MD5 hash of the response.
- The tool must be written in the Go programming language
- The tool must be able to perform the requests in parallel so that the tool can complete sooner. The order in which addresses are printed is not important.
- The tool must be able to limit the number of parallel requests, to prevent exhausting local resources. The tool must accept a flag to indicate this limit, and it should default to 10 if the flag is not provided.
- The tool must have unit tests.
- A README.md must be included describing the usage of this tool.
- Use the simplest structure possible
- Do not use external dependencies. Use only the Go standard library. If you are not very familiar with Go’s standard library, a good quality solution should be achievable using only the `flag`, `net/http` and `sync` packages.
- Do not add any extra features to the tool.
- Building the tool with `-race` flag, or running the tests with `-race`, should show no data races. We strongly recommend running both the tool and the tests with the `-race` flag before submitting a solution.
- Submit the results as a link to a github repository.
- It's perfectly acceptable to write this tool in a single source file. Successful submissions have had as few as 60 lines of code. You should also feel free to split the code up into separate packages, but please don't make it unnecessarily complex.
- Please ensure that the examples written below work with your solution. In particular please note that some URLs start with `http://` and some don’t.