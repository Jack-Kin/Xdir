package msg6030

import (
	"encoding/binary"
	"io"
)

//proto
const Name = "6030"

///////////////request message
type RequestBody struct {
	Version			uint16	// 版本			Short
	ClientType 		uint8	// 客户端代号		Char
	ClientID		uint32	// 客户端版本号	Int
	FreeType 		uint8 	// 是否收费		Char
	KlineType		uint8  	// K线类型		Byte
	RequestMode		uint8	// 请求类型		Byte
	FuQuan 			uint8	// 复权类型		Byte
	Start			uint32	// 开始日期		uInt
	End				uint32	// 结束日期		uInt
	Size			uint32	// 请求根数		uInt
	Market 			uint32 	// 市场ID		uInt
}

type RequestMessage struct {
	Body 		RequestBody
	CodeSize 	uint16 		// 代码字符串长度	ushort
	Code 		[]byte   	// 代码	String
}

func NewRequest(version uint16, clientType uint8,
	clientID uint32, freeType uint8, klineType uint8,
	requestmode uint8, fuquan uint8, start uint32,
	end uint32, size uint32, market uint32,
	code []byte) *RequestMessage {
	return &RequestMessage{
		Body:     RequestBody{
			Version: 		version,
			ClientType: 	clientType,
			ClientID: 		clientID,
			FreeType: 		freeType,
			KlineType: 		klineType,
			RequestMode: 	requestmode,
			FuQuan: 		fuquan,
			Start: 			start,
			End: 			end,
			Size: 			size,
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

// 只写了version 5
type ResponseData5 struct {
	Tradeday 			uint32		// 日期		uInt，例如1705251447
	Open 				int32		// 开盘		Int
	High 				int32		// 最高		Int
	Low 				int32 		// 最低		Int
	Close 				int32 		// 收盘		Int
	Volume 				uint32  	// 成交量	无符号Large Int, 月, 季，半年，年周期单位是万手, 其他是手，成交量根据请求可能带复权 科创板单位股
	Amount 				uint32 		// 成交额	无符号Large Int, 单位是万元
	TurnoverRate		uint32		// 换手率	Int   计算得到的换手率扩大100倍，保留2位小数
	Settle 				uint32 		// 结算价	Int
	Interest        	uint64  	// 持仓量	UINT64
}

type ResponseBody5 struct {
	// xml文件中ver字段不需要
	IsPay 				uint8
	KlineType 			uint8
	RequestMode 		uint8
	Fuquan 				uint8
	Scale 				uint8
	IpoPrice 			uint32
	IpoDate				uint32
	TotalNum  			uint32
	ReturnNum 			uint16
	Market 				uint32
}

type ResponseMessage struct {
	// ver = 5
	Body5 		ResponseBody5
	CodeSize5 	uint16
	Code5		[]byte
	Data5		[]ResponseData5
}

func NewResponse() *ResponseMessage { return &ResponseMessage{} }

func (m *ResponseMessage) Reset() { *m = ResponseMessage{} }

func (m *ResponseMessage) Unmarshal(r io.Reader) error {
	//read body
	err := binary.Read(r, binary.LittleEndian, &m.Body5)
	if err != nil {
		return err
	}
	//read  code size
	err = binary.Read(r, binary.LittleEndian, &m.CodeSize5)
	if err != nil {
		return err
	}
	//read code
	m.Code5 = make([]byte, m.CodeSize5)
	err = binary.Read(r, binary.LittleEndian, &m.Code5)
	if err != nil {
		return err
	}
	//read data
	m.Data5 = make([]ResponseData5, m.Body5.ReturnNum)
	//m.Item = make([]ResponseItem, m.Body.CurrentCount)
	err = binary.Read(r, binary.LittleEndian, &m.Data5)
	if err != nil {
		return err
	}
	return nil
}