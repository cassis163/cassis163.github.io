cd ./src
go install github.com/a-h/templ/cmd/templ@latest
$GOPATH/bin/templ generate .
GOBIN=$(pwd)/../functions go install .
