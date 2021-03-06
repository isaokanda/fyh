FLAG="-s -w -X main.version=$(git tag | sort -V | tail -1)"

#go build "$FLAG"
go build -ldflags "$FLAG" 
