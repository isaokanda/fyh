FLAG="-s -w -X main.version=$(git tag | sort -V | tail -1)-$(date +%Y%m%d)"

#go build "$FLAG"
go build -ldflags "$FLAG" 
