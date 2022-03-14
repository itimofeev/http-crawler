
build:
	go build -race -o ./bin/http-crawler github.com/itimofeev/http-crawler

test:
	go test -race ./...

run:
	go run -race github.com/itimofeev/http-crawler -parallel 10 google.com yandex.ru https://adjust.com http://www.reddit.com/r/funny/