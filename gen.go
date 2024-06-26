package helpy

// int8
//go:generate go run gen/pointer/pointer.go -type int8 -package-name int8y
//go:generate go run gen/unique/unique.go -type int8 -package-name int8y
//go:generate go run gen/minmax_number/minmax.go -type int8 -package-name int8y
//go:generate go run gen/contains/contains.go -type int8 -package-name int8y
//go:generate go run gen/set/set.go -type int8 -package-name int8y
//go:generate go run gen/rand_int/rand.go -type int8 -package-name int8y
//go:generate go run gen/sort_int/sort.go -type int8 -package-name int8y
//go:generate go run gen/abs_signed_number/abs.go -type int8 -package-name int8y
//go:generate go run gen/fromstring_int/fromstring.go -type int8 -package-name int8y -bit-size 8

// uint8
//go:generate go run gen/pointer/pointer.go -type uint8 -package-name uint8y
//go:generate go run gen/unique/unique.go -type uint8 -package-name uint8y
//go:generate go run gen/minmax_number/minmax.go -type uint8 -package-name uint8y
//go:generate go run gen/contains/contains.go -type uint8 -package-name uint8y
//go:generate go run gen/set/set.go -type uint8 -package-name uint8y
//go:generate go run gen/rand_int/rand.go -type uint8 -package-name uint8y
//go:generate go run gen/sort_int/sort.go -type uint8 -package-name uint8y
//go:generate go run gen/fromstring_int/fromstring.go -type uint8 -package-name uint8y -bit-size 8

// int16
//go:generate go run gen/pointer/pointer.go -type int16 -package-name int16y
//go:generate go run gen/unique/unique.go -type int16 -package-name int16y
//go:generate go run gen/minmax_number/minmax.go -type int16 -package-name int16y
//go:generate go run gen/contains/contains.go -type int16 -package-name int16y
//go:generate go run gen/set/set.go -type int16 -package-name int16y
//go:generate go run gen/rand_int/rand.go -type int16 -package-name int16y
//go:generate go run gen/sort_int/sort.go -type int16 -package-name int16y
//go:generate go run gen/abs_signed_number/abs.go -type int16 -package-name int16y
//go:generate go run gen/fromstring_int/fromstring.go -type int16 -package-name int16y -bit-size 16

// uint16
//go:generate go run gen/pointer/pointer.go -type uint16 -package-name uint16y
//go:generate go run gen/unique/unique.go -type uint16 -package-name uint16y
//go:generate go run gen/minmax_number/minmax.go -type uint16 -package-name uint16y
//go:generate go run gen/contains/contains.go -type uint16 -package-name uint16y
//go:generate go run gen/set/set.go -type uint16 -package-name uint16y
//go:generate go run gen/rand_int/rand.go -type uint16 -package-name uint16y
//go:generate go run gen/sort_int/sort.go -type uint16 -package-name uint16y
//go:generate go run gen/fromstring_int/fromstring.go -type uint16 -package-name uint16y -bit-size 16

// int32
//go:generate go run gen/pointer/pointer.go -type int32 -package-name int32y
//go:generate go run gen/unique/unique.go -type int32 -package-name int32y
//go:generate go run gen/minmax_number/minmax.go -type int32 -package-name int32y
//go:generate go run gen/contains/contains.go -type int32 -package-name int32y
//go:generate go run gen/set/set.go -type int32 -package-name int32y
//go:generate go run gen/rand_int/rand.go -type int32 -package-name int32y
//go:generate go run gen/sort_int/sort.go -type int32 -package-name int32y
//go:generate go run gen/abs_signed_number/abs.go -type int32 -package-name int32y
//go:generate go run gen/fromstring_int/fromstring.go -type int32 -package-name int32y -bit-size 32

// uint32
//go:generate go run gen/pointer/pointer.go -type uint32 -package-name uint32y
//go:generate go run gen/unique/unique.go -type uint32 -package-name uint32y
//go:generate go run gen/minmax_number/minmax.go -type uint32 -package-name uint32y
//go:generate go run gen/contains/contains.go -type uint32 -package-name uint32y
//go:generate go run gen/set/set.go -type uint32 -package-name uint32y
//go:generate go run gen/rand_int/rand.go -type uint32 -package-name uint32y
//go:generate go run gen/sort_int/sort.go -type uint32 -package-name uint32y
//go:generate go run gen/fromstring_int/fromstring.go -type uint32 -package-name uint32y -bit-size 32

// int
//go:generate go run gen/pointer/pointer.go -type int -package-name inty
//go:generate go run gen/unique/unique.go -type int -package-name inty
//go:generate go run gen/minmax_number/minmax.go -type int -package-name inty
//go:generate go run gen/contains/contains.go -type int -package-name inty
//go:generate go run gen/set/set.go -type int -package-name inty
//go:generate go run gen/rand_int/rand.go -type int -package-name inty
//go:generate go run gen/sort_int/sort.go -type int -package-name inty
//go:generate go run gen/abs_signed_number/abs.go -type int -package-name inty
//go:generate go run gen/fromstring_int/fromstring.go -type int -package-name inty -bit-size 32

// uint
//go:generate go run gen/pointer/pointer.go -type uint -package-name uinty
//go:generate go run gen/unique/unique.go -type uint -package-name uinty
//go:generate go run gen/minmax_number/minmax.go -type uint -package-name uinty
//go:generate go run gen/contains/contains.go -type uint -package-name uinty
//go:generate go run gen/set/set.go -type uint -package-name uinty
//go:generate go run gen/rand_int/rand.go -type uint -package-name uinty
//go:generate go run gen/sort_int/sort.go -type uint -package-name uinty
//go:generate go run gen/fromstring_int/fromstring.go -type uint -package-name uinty -bit-size 32

// int64
//go:generate go run gen/pointer/pointer.go -type int64 -package-name int64y
//go:generate go run gen/unique/unique.go -type int64 -package-name int64y
//go:generate go run gen/minmax_number/minmax.go -type int64 -package-name int64y
//go:generate go run gen/contains/contains.go -type int64 -package-name int64y
//go:generate go run gen/set/set.go -type int64 -package-name int64y
//go:generate go run gen/rand_int/rand.go -type int64 -package-name int64y
//go:generate go run gen/sort_int/sort.go -type int64 -package-name int64y
//go:generate go run gen/abs_signed_number/abs.go -type int64 -package-name int64y
//go:generate go run gen/fromstring_int/fromstring.go -type int64 -package-name int64y -bit-size 64

// uint64
//go:generate go run gen/pointer/pointer.go -type uint64 -package-name uint64y
//go:generate go run gen/unique/unique.go -type uint64 -package-name uint64y
//go:generate go run gen/minmax_number/minmax.go -type uint64 -package-name uint64y
//go:generate go run gen/contains/contains.go -type uint64 -package-name uint64y
//go:generate go run gen/set/set.go -type uint64 -package-name uint64y
//go:generate go run gen/rand_int/rand.go -type uint64 -package-name uint64y
//go:generate go run gen/sort_int/sort.go -type uint64 -package-name uint64y
//go:generate go run gen/fromstring_int/fromstring.go -type uint64 -package-name uint64y -bit-size 64

// float32
//go:generate go run gen/pointer/pointer.go -type float32 -package-name float32y
//go:generate go run gen/unique/unique.go -type float32 -package-name float32y
//go:generate go run gen/contains/contains.go -type float32 -package-name float32y
//go:generate go run gen/set/set.go -type float32 -package-name float32y
//go:generate go run gen/rand_float/rand.go -type float32 -package-name float32y
//go:generate go run gen/abs_signed_number/abs.go -type float32 -package-name float32y
//go:generate go run gen/fromstring_float/fromstring.go -type float32 -package-name float32y -bit-size 32

// float64
//go:generate go run gen/pointer/pointer.go -type float64 -package-name float64y
//go:generate go run gen/unique/unique.go -type float64 -package-name float64y
//go:generate go run gen/contains/contains.go -type float64 -package-name float64y
//go:generate go run gen/set/set.go -type float64 -package-name float64y
//go:generate go run gen/rand_float/rand.go -type float64 -package-name float64y
//go:generate go run gen/abs_signed_number/abs.go -type float64 -package-name float64y
//go:generate go run gen/fromstring_float/fromstring.go -type float64 -package-name float64y -bit-size 64

// bool
//go:generate go run gen/pointer/pointer.go -type bool -package-name booly

// string
//go:generate go run gen/pointer/pointer.go -type string -package-name stringy
//go:generate go run gen/unique/unique.go -type string -package-name stringy
//go:generate go run gen/set/set.go -type string -package-name stringy

// time.Time
//go:generate go run gen/pointer/pointer.go -type time.Time -type-name Time -package-name timey -import time
