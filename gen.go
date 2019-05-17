package main

func main() {}

// int8
//go:generate go run gen/pointer/pointer.go -type int8 -package-name int8y
//go:generate go run gen/unique/unique.go -type int8 -package-name int8y
//go:generate go run gen/minmax/minmax.go -type int8 -package-name int8y
//go:generate go run gen/contains/contains.go -type int8 -package-name int8y
//go:generate go run gen/set/set.go -type int8 -package-name int8y

// uint8
//go:generate go run gen/pointer/pointer.go -type uint8 -package-name uint8y
//go:generate go run gen/unique/unique.go -type uint8 -package-name uint8y
//go:generate go run gen/minmax/minmax.go -type uint8 -package-name uint8y
//go:generate go run gen/contains/contains.go -type uint8 -package-name uint8y
//go:generate go run gen/set/set.go -type uint8 -package-name uint8y

// int16
//go:generate go run gen/pointer/pointer.go -type int16 -package-name int16y
//go:generate go run gen/unique/unique.go -type int16 -package-name int16y
//go:generate go run gen/minmax/minmax.go -type int16 -package-name int16y
//go:generate go run gen/contains/contains.go -type int16 -package-name int16y
//go:generate go run gen/set/set.go -type int16 -package-name int16y

// uint16
//go:generate go run gen/pointer/pointer.go -type uint16 -package-name uint16y
//go:generate go run gen/unique/unique.go -type uint16 -package-name uint16y
//go:generate go run gen/minmax/minmax.go -type uint16 -package-name uint16y
//go:generate go run gen/contains/contains.go -type uint16 -package-name uint16y
//go:generate go run gen/set/set.go -type uint16 -package-name uint16y

// int32
//go:generate go run gen/pointer/pointer.go -type int32 -package-name int32y
//go:generate go run gen/unique/unique.go -type int32 -package-name int32y
//go:generate go run gen/minmax/minmax.go -type int32 -package-name int32y
//go:generate go run gen/contains/contains.go -type int32 -package-name int32y
//go:generate go run gen/set/set.go -type int32 -package-name int32y

// uint32
//go:generate go run gen/pointer/pointer.go -type uint32 -package-name uint32y
//go:generate go run gen/unique/unique.go -type uint32 -package-name uint32y
//go:generate go run gen/minmax/minmax.go -type uint32 -package-name uint32y
//go:generate go run gen/contains/contains.go -type uint32 -package-name uint32y
//go:generate go run gen/set/set.go -type uint32 -package-name uint32y

// int
//go:generate go run gen/pointer/pointer.go -type int -package-name inty
//go:generate go run gen/unique/unique.go -type int -package-name inty
//go:generate go run gen/minmax/minmax.go -type int -package-name inty
//go:generate go run gen/contains/contains.go -type int -package-name inty
//go:generate go run gen/set/set.go -type int -package-name inty

// uint
//go:generate go run gen/pointer/pointer.go -type uint -package-name uinty
//go:generate go run gen/unique/unique.go -type uint -package-name uinty
//go:generate go run gen/minmax/minmax.go -type uint -package-name uinty
//go:generate go run gen/contains/contains.go -type uint -package-name uinty
//go:generate go run gen/set/set.go -type uint -package-name uinty

// int64
//go:generate go run gen/pointer/pointer.go -type int64 -package-name int64y
//go:generate go run gen/unique/unique.go -type int64 -package-name int64y
//go:generate go run gen/minmax/minmax.go -type int64 -package-name int64y
//go:generate go run gen/contains/contains.go -type int64 -package-name int64y
//go:generate go run gen/set/set.go -type int64 -package-name int64y

// uint64
//go:generate go run gen/pointer/pointer.go -type uint64 -package-name uint64y
//go:generate go run gen/unique/unique.go -type uint64 -package-name uint64y
//go:generate go run gen/minmax/minmax.go -type uint64 -package-name uint64y
//go:generate go run gen/contains/contains.go -type uint64 -package-name uint64y
//go:generate go run gen/set/set.go -type uint64 -package-name uint64y

// float32
//go:generate go run gen/pointer/pointer.go -type float32 -package-name float32y
//go:generate go run gen/unique/unique.go -type float32 -package-name float32y
//go:generate go run gen/minmax/minmax.go -type float32 -package-name float32y
//go:generate go run gen/contains/contains.go -type float32 -package-name float32y
//go:generate go run gen/set/set.go -type float32 -package-name float32y

// float64
//go:generate go run gen/pointer/pointer.go -type float64 -package-name float64y
//go:generate go run gen/unique/unique.go -type float64 -package-name float64y
//go:generate go run gen/minmax/minmax.go -type float64 -package-name float64y
//go:generate go run gen/contains/contains.go -type float64 -package-name float64y
//go:generate go run gen/set/set.go -type float64 -package-name float64y

// bool
//go:generate go run gen/pointer/pointer.go -type bool -package-name booly

// string
//go:generate go run gen/pointer/pointer.go -type string -package-name stringy
//go:generate go run gen/unique/unique.go -type string -package-name stringy
//go:generate go run gen/set/set.go -type string -package-name stringy

// time.Time
//go:generate go run gen/pointer/pointer.go -type time.Time -type-name Time -package-name timey -import time
