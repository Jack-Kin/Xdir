package msg6131

import (
	"encoding/binary"
	"io"
)

//proto
const Name = "6131"


///////////////request message
type RequestBody struct {
	Version 		uint16
	ClientType 		uint8
	ClientID 		uint32
	FreeType 		uint8
	MagicID 		uint32
	KlineType		uint8
	Pixels 			uint16
	Market 			uint32
}

type RequestMessage struct {
	Body 		RequestBody
	CodeSize 	uint16
	Code 		[]byte
}

func NewRequest(version uint16, clientType uint8,
	clientID uint32, freeType uint8, magicID uint32,
	klineType uint8, pixels uint16, market uint32,
	code []byte) *RequestMessage {
	return &RequestMessage{
		Body:     RequestBody{
			Version:  		version,
			ClientType: 	clientType,
			ClientID: 		clientID,
			FreeType: 		freeType,
			MagicID: 		magicID,
			KlineType: 		klineType,
			Pixels: 		pixels,
			Market: 		market,
		},
		CodeSize: uint16(len(code)),
		Code:     code,
	}
}

func (m *RequestMessage) Size() uint16{
	return (uint16)(binary.Size(m.Body)) + m.CodeSize + 2
}

func (m *RequestMessage) Marshal(w io.Writer) error {
	//Write body
	err := binary.Write(w, binary.LittleEndian, m.Body)
	if err != nil {
		return err
	}
	//Write code size
	err = binary.Write(w, binary.LittleEndian, m.CodeSize)
	if err != nil {
		return err
	}

	//Write code size
	err = binary.Write(w, binary.LittleEndian, m.Code[0:m.CodeSize])
	if err != nil {
		return err
	}
	return nil
}

func (m *RequestMessage) Unmarshal(r io.Reader) error {
	//read body
	err := binary.Read(r, binary.LittleEndian, &m.Body)
	if err != nil {
		return err
	}
	//read  code size
	err = binary.Read(r, binary.LittleEndian, &m.CodeSize)
	if err != nil {
		return err
	}
	//read code
	m.Code = make([]byte, m.CodeSize)
	err = binary.Read(r, binary.LittleEndian, &m.Code)
	if err != nil {
		return err
	}
	return nil
}


/////////////Response message
type ResponseData struct {
	Tradeday 		uint32
	Value1 			int32
	Value2 			int32
}

type ResponseBody struct {
	MagicID 	uint32
	Type 		uint8
	Pixels 		uint16
	ReturnNum 	uint16
	Market 		uint32
}

type ResponseMessage struct {
	Body 		ResponseBody
	CodeSize 	uint16
	Code 		[]byte
	Data 		[]ResponseData
}

func NewResponse() *ResponseMessage { return &ResponseMessage{} }

func (m *ResponseMessage) Reset() { *m = ResponseMessage{} }

func (m *ResponseMessage) Unmarshal(r io.Reader) error {
	//read body
	err := binary.Read(r, binary.LittleEndian, &m.Body)
	if err != nil {
		return err
	}
	//read  code size
	err = binary.Read(r, binary.LittleEndian, &m.CodeSize)
	if err != nil {
		return err
	}
	//read code
	m.Code = make([]byte, m.CodeSize)
	err = binary.Read(r, binary.LittleEndian, &m.Code)
	if err != nil {
		return err
	}
	//read data
	m.Data = make([]ResponseData, m.Body.ReturnNum)
	//m.Item = make([]ResponseItem, m.Body.CurrentCount)
	err = binary.Read(r, binary.LittleEndian, &m.Data)
	if err != nil {
		return err
	}
	return nil
}