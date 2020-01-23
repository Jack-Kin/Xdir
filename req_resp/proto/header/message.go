package header

import (
	"encoding/binary"
	"io"
)

const Name = "header"

type Header struct {
	Prefix		byte
	Type 		int16
	Attribute  	[2]byte
	PackageSize	uint16
}

func New(t int16, s uint16) *Header {
	return &Header{
		Prefix:      byte('{'),
		Type:        t,
		Attribute:   [2]byte{0x80, 0x00},
		PackageSize: s,
	}
}

func (h *Header) Reset() { *h = Header{} }

func (h *Header) SetType(t int16) {
	h.Type = t
}

func (h *Header) SetPackageSize(s uint16) {
	h.PackageSize = s
}

func (h *Header) Marshal(w io.Writer) error {
	//Write
	err := binary.Write(w, binary.LittleEndian, h)
	if err != nil {
		return err
	}
	return nil
}

func (h *Header) Unmarshal(r io.Reader) error {
	//read body
	err := binary.Read(r, binary.LittleEndian, h)
	if err != nil {
		return err
	}
	return nil
}