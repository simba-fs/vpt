build:
	# mac
	GOOS=darwin GOARCH=amd64 go build -o dist/stl-macOS-amd64 .
	# linux
	GOOS=linux GOARCH=arm64 go build -o dist/stl-linux-arm64 .
	GOOS=linux GOARCH=386 go build -o dist/stl-linux-386 .
	GOOS=linux GOARCH=amd64 go build -o dist/stl-linux-amd64 .
	# windows
	GOOS=windows GOARCH=386 go build -o dist/stl-windows-386 .
	GOOS=windows GOARCH=amd64 go build -o dist/stl-windows-amd64 .

clean:
	rm -f ./dist/*

run:
	go run .
