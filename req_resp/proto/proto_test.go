package proto

import (
	"bytes"
	"testing"

	"req_resp/proto/header"
	"req_resp/proto/msg6131"

	"github.com/stretchr/testify/require"
)

func TestProtoMarshal(t *testing.T) {

	head := header.New(6, 31)
	buf := new(bytes.Buffer)
	err := marshal(buf, head)
	require.NoError(t, err)

	req := msg6131.NewRequest(1, 0, 0, 0, 999, 7, 1792, 150, []byte("008240"))
	buf1 := new(bytes.Buffer)
	err = marshal(buf1, req)
	require.NoError(t, err)
}

func TestProtoUnmarshal(t *testing.T) {
	head := header.New(6, 31)
	buf := new(bytes.Buffer)
	err := marshal(buf, head)
	require.NoError(t, err)
	head.Reset()
	err = unmarshal(buf, head)

}

func TestProtoEncode(t *testing.T) {

	head := header.New(6, 31)
	buf := new(bytes.Buffer)
	encoder := NewEncoder(buf)

	encoder.Push(head)
	require.NoError(t, encoder.Error())

	req := msg6131.NewRequest(1, 0, 0, 0, 999, 7, 1792, 150, []byte("008240"))
	encoder.Push(req)
	require.NoError(t, encoder.Error())

	tail := Tail{}
	encoder.Push(head).Push(req).Push(tail)
	//require.NoError(t, encoder.Error())

}

func TestProtoDecode(t *testing.T) {
	head := header.New(6, 31)
	buf := new(bytes.Buffer)
	encoder := NewEncoder(buf)

	encoder.Push(head)
	require.NoError(t, encoder.Error())

	decoder := NewDecoder(buf)
	decoder.Fetch(head)
	require.NoError(t, decoder.Error())

}
