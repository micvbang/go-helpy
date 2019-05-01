package main

func main() {}

// int
//go:generate go run gen/pointer/pointer.go -type int -package-name inty
//go:generate go run gen/unique/unique.go -type int -package-name inty
//go:generate go run gen/minmax/minmax.go -type int -package-name inty

// int64
//go:generate go run gen/pointer/pointer.go -type int64 -package-name int64y
//go:generate go run gen/unique/unique.go -type int64 -package-name int64y
//go:generate go run gen/minmax/minmax.go -type int64 -package-name int64y

// uint64
//go:generate go run gen/pointer/pointer.go -type uint64 -package-name uint64y
//go:generate go run gen/unique/unique.go -type uint64 -package-name uint64y
//go:generate go run gen/minmax/minmax.go -type uint64 -package-name uint64y

// bool
//go:generate go run gen/pointer/pointer.go -type bool -package-name booly

// string
//go:generate go run gen/pointer/pointer.go -type string -package-name stringy
//go:generate go run gen/unique/unique.go -type string -package-name stringy

// time.Time
//go:generate go run gen/pointer/pointer.go -type time.Time -type-name Time -package-name timey -import time
