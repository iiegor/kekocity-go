package packet

type INetMessageWriter interface {
	WritePacket() IPacket
}

type INetMessageReader interface {
	GetHeader() uint8
	ReadPacket(IPacket) error
}

type IPacket interface {
	Reset()
	CanAdd(_size uint16) bool

	GetHeader() uint16
	SetHeader()

	GetBuffer() [PACKET_MAXSIZE]uint8
	GetBufferSlice() []uint8
	GetMsgSize() uint16

	ReadUint8() (uint8, error)
	ReadUint16() (uint16, error)
	ReadInt16() (int16, error)
	ReadUint32() (uint32, error)
	ReadInt32() (int, error)
	ReadUint64() (uint64, error)
	ReadInt64() (int64, error)
	ReadString() (string, error)
	ReadBool() (bool, error)

	AddUint8(_value uint8) bool
	AddUint16(_value uint16) bool
	AddUint32(_value uint32) bool
	AddUint64(_value uint64) bool
	AddBool(_value bool) bool
	AddString(_value string) bool
	AddBuffer(_value []uint8) bool
}
