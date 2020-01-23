package msg6131

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	req := NewRequest(1, 0, 0, 0,
		999, 7, 1792, 150, []byte("008240"))
	//req := NewRequest(116, []byte("00700"), 10, 0)
	// fmt.Printf("%#+v", req)
	// require.NoError(t, err)
	//require.Equal(t, byte(150), req.Body.Market)
	assert.Equal(t, []byte("008240"), req.Code)
	//assert.Equal(t, byte(0), req.Body.FuQuan)
	// assert.Equal(t, hash, m.HashID())
	// assert.Equal(t, hash, m.HashID())
	buf := new(bytes.Buffer)
	err := req.Marshal(buf)
	require.NoError(t, err)

	res := new(RequestMessage)
	err = res.Unmarshal(buf)
	require.NoError(t, err)
	assert.Equal(t, res, req)
}