// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mytype

import (
	"github.com/openark/vitess/go/bytes2"

	"bytes"

	"github.com/openark/vitess/go/bson"
)

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

// MarshalBson bson-encodes MyType.
func (myType *MyType) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	// *int64
	if myType.Ptr == nil {
		bson.EncodePrefix(buf, bson.Null, "Ptr")
	} else {
		bson.EncodeInt64(buf, "Ptr", (*myType.Ptr))
	}
	// **int64
	if myType.PtrPtr == nil {
		bson.EncodePrefix(buf, bson.Null, "PtrPtr")
	} else {
		// *int64
		if (*myType.PtrPtr) == nil {
			bson.EncodePrefix(buf, bson.Null, "PtrPtr")
		} else {
			bson.EncodeInt64(buf, "PtrPtr", (*(*myType.PtrPtr)))
		}
	}
	// *[]byte
	if myType.PtrBytes == nil {
		bson.EncodePrefix(buf, bson.Null, "PtrBytes")
	} else {
		bson.EncodeBinary(buf, "PtrBytes", (*myType.PtrBytes))
	}
	// *[]int64
	if myType.PtrSlice == nil {
		bson.EncodePrefix(buf, bson.Null, "PtrSlice")
	} else {
		// []int64
		{
			bson.EncodePrefix(buf, bson.Array, "PtrSlice")
			lenWriter := bson.NewLenWriter(buf)
			for _i, _v1 := range *myType.PtrSlice {
				bson.EncodeInt64(buf, bson.Itoa(_i), _v1)
			}
			lenWriter.Close()
		}
	}
	// *map[string]int64
	if myType.PtrMap == nil {
		bson.EncodePrefix(buf, bson.Null, "PtrMap")
	} else {
		// map[string]int64
		{
			bson.EncodePrefix(buf, bson.Object, "PtrMap")
			lenWriter := bson.NewLenWriter(buf)
			for _k, _v2 := range *myType.PtrMap {
				bson.EncodeInt64(buf, _k, _v2)
			}
			lenWriter.Close()
		}
	}
	// *Custom
	if myType.PtrCustom == nil {
		bson.EncodePrefix(buf, bson.Null, "PtrCustom")
	} else {
		(*myType.PtrCustom).MarshalBson(buf, "PtrCustom")
	}

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into MyType.
func (myType *MyType) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for MyType", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Ptr":
			// *int64
			if kind != bson.Null {
				myType.Ptr = new(int64)
				(*myType.Ptr) = bson.DecodeInt64(buf, kind)
			}
		case "PtrPtr":
			// **int64
			if kind != bson.Null {
				myType.PtrPtr = new(*int64)
				// *int64
				if kind != bson.Null {
					(*myType.PtrPtr) = new(int64)
					(*(*myType.PtrPtr)) = bson.DecodeInt64(buf, kind)
				}
			}
		case "PtrBytes":
			// *[]byte
			if kind != bson.Null {
				myType.PtrBytes = new([]byte)
				(*myType.PtrBytes) = bson.DecodeBinary(buf, kind)
			}
		case "PtrSlice":
			// *[]int64
			if kind != bson.Null {
				myType.PtrSlice = new([]int64)
				// []int64
				if kind != bson.Null {
					if kind != bson.Array {
						panic(bson.NewBsonError("unexpected kind %v for (*myType.PtrSlice)", kind))
					}
					bson.Next(buf, 4)
					(*myType.PtrSlice) = make([]int64, 0, 8)
					for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
						bson.SkipIndex(buf)
						var _v1 int64
						_v1 = bson.DecodeInt64(buf, kind)
						(*myType.PtrSlice) = append((*myType.PtrSlice), _v1)
					}
				}
			}
		case "PtrMap":
			// *map[string]int64
			if kind != bson.Null {
				myType.PtrMap = new(map[string]int64)
				// map[string]int64
				if kind != bson.Null {
					if kind != bson.Object {
						panic(bson.NewBsonError("unexpected kind %v for (*myType.PtrMap)", kind))
					}
					bson.Next(buf, 4)
					(*myType.PtrMap) = make(map[string]int64)
					for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
						_k := bson.ReadCString(buf)
						var _v2 int64
						_v2 = bson.DecodeInt64(buf, kind)
						(*myType.PtrMap)[_k] = _v2
					}
				}
			}
		case "PtrCustom":
			// *Custom
			if kind != bson.Null {
				myType.PtrCustom = new(Custom)
				(*myType.PtrCustom).UnmarshalBson(buf, kind)
			}
		default:
			bson.Skip(buf, kind)
		}
	}
}
