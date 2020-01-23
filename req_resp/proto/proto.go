package proto

import (
	"encoding/binary"
	"errors"
	"io"
)

type Marshaler interface {
	Marshal(w io.Writer) error
}

type Unmarshaler interface {
	Reset()

	Unmarshal(w io.Reader) error
}

func marshal(w io.Writer, v interface{}) error {

	if pm, ok := v.(Marshaler); ok {
		return pm.Marshal(w)
	} else {
		return errors.New("not Marshaler")
	}

}

func unmarshal(r io.Reader, v interface{}) error {

	if pm, ok := v.(Unmarshaler); ok {
		pm.Reset()
		return pm.Unmarshal(r)
	} else {
		return errors.New("not Unmarshaler")
	}

}

//协议尾部
type Tail struct {
	Postfix byte
}

func NewTail() *Tail {
	return &Tail{
		Postfix: byte('}'),
	}
}

func (t *Tail) Marshal(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, t.Postfix)
	if err != nil {
		return err
	}
	return nil
}

//默认小端模式
type Encoder struct {
	writer io.Writer
	err    error
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
	}
}

func (e *Encoder) Error() error {
	return e.err
}

func (e *Encoder) Push(v interface{}) *Encoder {
	return e.errFilter(func() {
		e.err = marshal(e.writer, v)
	})
}

func (e *Encoder) errFilter(f func()) *Encoder {
	if e.err == nil {
		f()
	}
	return e
}

//默认小端模式
type Decoder struct {
	reader io.Reader
	err    error
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: r,
	}
}

func (e *Decoder) Error() error {
	return e.err
}

func (e *Decoder) Fetch(v interface{}) *Decoder {
	return e.errFilter(func() {
		e.err = unmarshal(e.reader, v)
	})
}

func (e *Decoder) errFilter(f func()) *Decoder {
	if e.err == nil {
		f()
	}
	return e
}
