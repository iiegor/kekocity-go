package packet

import "fmt"

const (
	QTPACKET_MAXSIZE = 16384
)

type QTPacket struct {
	readPos uint16
	MsgSize uint16

	Buffer [QTPACKET_MAXSIZE]uint8
}

func NewQTPacket() *QTPacket {
	packet := &QTPacket{}
	packet.Reset()

	return packet
}

func NewQTPacketExt(_header uint8) *QTPacket {
	packet := NewQTPacket()
	packet.AddUint8(_header)

	return packet
}

func (p *QTPacket) Reset() {
	p.MsgSize = 0
	p.readPos = 2
}

func (p *QTPacket) CanAdd(_size uint16) bool {
	return (_size+p.readPos < QTPACKET_MAXSIZE-16)
}

func (p *QTPacket) CanRead(_size uint16) bool {
	return ((p.readPos + _size) < PACKET_MAXSIZE)
}

func (p *QTPacket) GetHeader() uint16 {
	p.MsgSize = uint16(uint16(p.Buffer[1]) | (uint16(p.Buffer[0]) << 8))
	return p.MsgSize
}

func (p *QTPacket) SetHeader() {
	p.Buffer[0] = uint8(p.MsgSize >> 8)
	p.Buffer[1] = uint8(p.MsgSize)
	p.MsgSize += 2
}

func (p *QTPacket) GetBuffer() [PACKET_MAXSIZE]uint8 {
	return p.Buffer
}

func (p *QTPacket) GetBufferSlice() []uint8 {
	size := p.MsgSize + 2
	return p.Buffer[2:size]
}

func (p *QTPacket) GetMsgSize() uint16 {
	return p.MsgSize
}

func (p *QTPacket) GetReadPos() uint16 {
	return p.readPos
}

// Byte
func (p *QTPacket) ReadUint8() (uint8, error) {
	if !p.CanRead(0) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 1, p.readPos, PACKET_MAXSIZE)
	}

	v := p.Buffer[p.readPos]
	p.readPos += 1
	return v, nil
}

// Short
func (p *QTPacket) ReadUint16() (uint16, error) {
	if !p.CanRead(1) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 2, p.readPos, PACKET_MAXSIZE)
	}

	v := uint16(uint16(p.Buffer[p.readPos+1]) | (uint16(p.Buffer[p.readPos]) << 8))
	p.readPos += 2
	return v, nil
}

func (p *QTPacket) ReadInt16() (int16, error) {
	value, err := p.ReadUint16()
	if err != nil {
		return 0, err
	}

	return int16(value), nil
}

// Int
func (p *QTPacket) ReadUint32() (uint32, error) {
	if !p.CanRead(3) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 4, p.readPos, PACKET_MAXSIZE)
	}

	v := uint32((uint32(p.Buffer[p.readPos+3]) | (uint32(p.Buffer[p.readPos+2]) << 8) |
		(uint32(p.Buffer[p.readPos+1]) << 16) | (uint32(p.Buffer[p.readPos]) << 24)))
	p.readPos += 4
	return v, nil
}

func (p *QTPacket) ReadInt32() (int, error) {
	value, err := p.ReadUint32()
	return int(value), err
}

func (p *QTPacket) ReadUint64() (uint64, error) {
	if !p.CanRead(7) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 8, p.readPos, PACKET_MAXSIZE)
	}

	v := uint64((uint64(p.Buffer[p.readPos+7]) | (uint64(p.Buffer[p.readPos+6]) << 8) |
		(uint64(p.Buffer[p.readPos+5]) << 16) | (uint64(p.Buffer[p.readPos+4]) << 24) |
		(uint64(p.Buffer[p.readPos+3]) << 32) | (uint64(p.Buffer[p.readPos+2]) << 40) |
		(uint64(p.Buffer[p.readPos+1]) << 48) | (uint64(p.Buffer[p.readPos]) << 56)))
	p.readPos += 8
	return v, nil
}

func (p *QTPacket) ReadInt64() (int64, error) {
	value, err := p.ReadUint64()
	return int64(value), err
}

func (p *QTPacket) ReadString() (string, error) {
	stringlen, err := p.ReadUint32()
	if err != nil {
		return "", err
	}

	if !p.CanRead(uint16(stringlen)) {
		return "", fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 7, p.readPos, PACKET_MAXSIZE)
	}

	// QString has a special value of 0xFFFFFFFF if it's NULL
	if stringlen >= 0xFFFFFFFF {
		return "", nil
	}

	v := ""
	for i := 0; uint32(i) < stringlen/uint32(2); i++ {
		val := uint16(uint16(p.Buffer[p.readPos+1]) | (uint16(p.Buffer[p.readPos]) << 8))
		p.readPos += 2

		v += string(val)
	}
	return v, nil
}

func (p *QTPacket) ReadBool() (bool, error) {
	bit, err := p.ReadUint8()
	if err != nil {
		return false, err
	}
	return (bit == 1), nil
}

func (p *QTPacket) AddUint8(_value uint8) bool {
	if !p.CanAdd(1) {
		return false
	}

	p.Buffer[p.readPos] = _value
	p.readPos += 1
	p.MsgSize += 1

	return true
}

func (p *QTPacket) AddUint16(_value uint16) bool {
	if !p.CanAdd(2) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1

	p.MsgSize += 2

	return true
}

func (p *QTPacket) AddUint32(_value uint32) bool {
	if !p.CanAdd(4) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value >> 24)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 16)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1

	p.MsgSize += 4

	return true
}

func (p *QTPacket) AddUint64(_value uint64) bool {
	if !p.CanAdd(8) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value >> 56)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 48)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 40)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 32)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 24)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 16)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1

	p.MsgSize += 8

	return true
}

func (p *QTPacket) AddBool(_value bool) bool {
	if _value {
		return p.AddUint8(1)
	}
	return p.AddUint8(0)
}

func (p *QTPacket) AddString(_value string) bool {
	stringlen := uint16(len(_value) * 2)
	if !p.CanAdd(stringlen * uint16(2)) {
		return false
	}

	p.AddUint32(uint32(stringlen))
	for i, _ := range _value {
		p.AddUint16(uint16(_value[i]))
	}

	return true
}

func (p *QTPacket) AddBuffer(_value []uint8) bool {
	size := uint16(len(_value))
	if !p.CanAdd(size) {
		return false
	}

	copy(p.Buffer[p.readPos:], _value[:])

	/*for i := 0; i < int(size); i++ {
		p.Buffer[p.readPos] = _value[i]
		p.readPos += 1
	}*/

	p.readPos += size
	p.MsgSize += size

	return true
}
