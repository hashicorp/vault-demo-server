// Copyright IBM Corp. 2015, 2026

package reflectwalk

//go:generate stringer -type=Location location.go

type Location uint

const (
	None Location = iota
	Map
	MapKey
	MapValue
	Slice
	SliceElem
	Array
	ArrayElem
	Struct
	StructField
	WalkLoc
)
