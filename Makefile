build:
	go build -o target/azstorageping main.go

clean:
	rm -rf target/

run:
	go run main.go $(account) $(key) $(container)