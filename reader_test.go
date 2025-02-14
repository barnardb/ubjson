package ubjson

import (
	"bytes"
	"reflect"
	"testing"
	"testing/iotest"
)

func TestReadBuf(t *testing.T) {
	expected := []byte("abcdefgh")
	r := newBinaryReader(iotest.OneByteReader(bytes.NewReader(expected)))
	bytes, err := r.readBuf(8)
	if err != nil {
		t.Fatalf("failed to readBuf: %+v\n", err)
	}
	if !reflect.DeepEqual(bytes, expected) {
		t.Errorf("\nexpected: %#v\nbut got:  %#v", expected, bytes)
	}
}

func TestReadString(t *testing.T) {
	expected := "abcdefgh"
	readerData := append([]byte{'U', 8}, expected...)
	r := newBinaryReader(iotest.OneByteReader(bytes.NewReader(readerData)))
	str, err := r.readString(8)
	if err != nil {
		t.Fatalf("failed to readBuf: %+v\n", err)
	}
	if str != expected {
		t.Errorf("\nexpected: %#v\nbut got:  %#v", expected, str)
	}
}
