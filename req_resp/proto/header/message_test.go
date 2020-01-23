package header

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {

	head := New(6, 31)

	assert.Equal(t, int16(6), head.Type)
	assert.Equal(t, byte('{'), head.Prefix)
	assert.Equal(t, int8(-126), head.Attribute[0])
	assert.Equal(t, int8(0), head.Attribute[1])
	assert.Equal(t, uint16(31), head.PackageSize)

	buf := new(bytes.Buffer)
	err := head.Marshal(buf)
	require.NoError(t, err)

	res := new(Header)
	err = res.Unmarshal(buf)
	require.NoError(t, err)
	assert.Equal(t, res, head)
}
