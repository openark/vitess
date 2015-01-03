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

	bson.EncodeInt(buf, "Local", myType.local)
	bson.EncodeInt(buf, "Local1", myType.Local2)

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
		case "Local":
			myType.local = bson.DecodeInt(buf, kind)
		case "Local1":
			myType.Local2 = bson.DecodeInt(buf, kind)
		default:
			bson.Skip(buf, kind)
		}
	}
}
