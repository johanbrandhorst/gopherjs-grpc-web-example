package main

//go:generate gopherjs build -m client.go -o html/index.js
//go:generate go-bindata -pkg compiled -nometadata -o compiled/client.go -prefix html ./html
//go:generate bash -c "rm html/*.js*"

func main() {}
