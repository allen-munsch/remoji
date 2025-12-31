source ~/.bashrc
load_go
go mod tidy
go build
mv remoji /usr/local/bin
