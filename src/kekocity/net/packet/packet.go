package packet

import "fmt"

const (
	PACKET_MAXSIZE = 16384
)

type Packet struct {
	readPos uint16
	MsgSize uint16

	Buffer [PACKET_MAXSIZE]uint8
}

// NewPacket creates a new Packet with no header
func NewPacket() *Packet {
	packet := &Packet{}
	packet.Reset()

	return packet
}

// NewPacketExt creates a new Packet with message header
func NewPacketExt(_header uint8) *Packet {
	packet := NewPacket()
	packet.AddUint8(_header)

	return packet
}

func (p *Packet) Reset() {
	p.MsgSize = 0
	p.readPos = 2
}

func (p *Packet) CanAdd(_size uint16) bool {
	return (_size+p.readPos < PACKET_MAXSIZE-16)
}

func (p *Packet) CanRead(_size uint16) bool {
	return ((p.readPos + _size) < PACKET_MAXSIZE)
}

func (p *Packet) GetHeader() uint16 {
	p.MsgSize = uint16(uint16(p.Buffer[0]) | (uint16(p.Buffer[1]) << 8))
	return p.MsgSize
}

func (p *Packet) SetHeader() {
	p.Buffer[0] = uint8(p.MsgSize >> 8)
	p.Buffer[1] = uint8(p.MsgSize)
	p.MsgSize += 2
}

func (p *Packet) GetBuffer() [PACKET_MAXSIZE]uint8 {
	return p.Buffer
}

func (p *Packet) GetBufferSlice() []uint8 {
	return p.Buffer[0:p.MsgSize]
}

func (p *Packet) GetMsgSize() uint16 {
	return p.MsgSize
}

func (p *Packet) ReadUint8() (uint8, error) {
	if !p.CanRead(0) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 1, p.readPos, PACKET_MAXSIZE)
	}

	v := p.Buffer[p.readPos]
	p.readPos += 1
	return v, nil
}

func (p *Packet) ReadUint16() (uint16, error) {
	if !p.CanRead(1) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 2, p.readPos, PACKET_MAXSIZE)
	}

	v := uint16(uint16(p.Buffer[p.readPos]) | (uint16(p.Buffer[p.readPos+1]) << 8))
	p.readPos += 2

	return v, nil
}

func (p *Packet) ReadInt16() (int16, error) {
	value, err := p.ReadUint16()
	if err != nil {
		return 0, err
	}

	return int16(value), nil
}

func (p *Packet) ReadUint32() (uint32, error) {
	if !p.CanRead(3) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 4, p.readPos, PACKET_MAXSIZE)
	}

	v := uint32((uint32(p.Buffer[p.readPos]) | (uint32(p.Buffer[p.readPos+1]) << 8) |
		(uint32(p.Buffer[p.readPos+2]) << 16) | (uint32(p.Buffer[p.readPos+3]) << 24)))
	p.readPos += 4

	return v, nil
}

func (p *Packet) ReadInt32() (int, error) {
	value, err := p.ReadUint32()
	return int(value), err
}

func (p *Packet) ReadUint64() (uint64, error) {
	if !p.CanRead(7) {
		return 0, fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 8, p.readPos, PACKET_MAXSIZE)
	}

	v := uint64((uint64(p.Buffer[p.readPos]) | (uint64(p.Buffer[p.readPos+1]) << 8) |
		(uint64(p.Buffer[p.readPos+2]) << 16) | (uint64(p.Buffer[p.readPos+3]) << 24) |
		(uint64(p.Buffer[p.readPos+4]) << 32) | (uint64(p.Buffer[p.readPos+5]) << 40) |
		(uint64(p.Buffer[p.readPos+6]) << 48) | (uint64(p.Buffer[p.readPos+7]) << 56)))
	p.readPos += 8

	return v, nil
}

func (p *Packet) ReadInt64() (int64, error) {
	value, err := p.ReadUint64()
	return int64(value), err
}

func (p *Packet) ReadString() (string, error) {
	stringlen, err := p.ReadUint16()
	if err != nil {
		return "", err
	}

	if !p.CanRead(uint16(stringlen)) {
		return "", fmt.Errorf("packet: unable to read %d bytes, current read position %d, buffer length %d", 7, p.readPos, PACKET_MAXSIZE)
	}

	v := string(p.Buffer[p.readPos : p.readPos+uint16(stringlen)])
	p.readPos += uint16(stringlen)

	return v, nil
}

func (p *Packet) ReadBool() (bool, error) {
	bit, err := p.ReadUint8()
	if err != nil {
		return false, err
	}
	return (bit == 1), nil
}

func (p *Packet) AddUint8(_value uint8) bool {
	if !p.CanAdd(1) {
		return false
	}

	p.Buffer[p.readPos] = _value
	p.readPos += 1
	p.MsgSize += 1

	return true
}

func (p *Packet) AddUint16(_value uint16) bool {
	if !p.CanAdd(2) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1

	p.MsgSize += 2

	return true
}

func (p *Packet) AddUint32(_value uint32) bool {
	if !p.CanAdd(4) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 16)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 24)
	p.readPos += 1

	p.MsgSize += 4

	return true
}

func (p *Packet) AddUint64(_value uint64) bool {
	if !p.CanAdd(8) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 16)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 24)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 32)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 40)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 48)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 56)
	p.readPos += 1

	p.MsgSize += 8

	return true
}

func (p *Packet) AddBool(_value bool) bool {
	if _value {
		return p.AddUint8(1)
	}
	return p.AddUint8(0)
}

func (p *Packet) AddString(_value string) bool {
	stringlen := uint16(len(_value))
	if !p.CanAdd(stringlen) {
		return false
	}

	p.AddUint16(uint16(stringlen))
	for i, _ := range _value {
		p.Buffer[p.readPos+uint16(i)] = _value[i]
	}

	p.readPos += stringlen
	p.MsgSize += uint16(stringlen)

	return true
}

func (p *Packet) AddBuffer(_value []uint8) bool {
	size := uint16(len(_value))
	if !p.CanAdd(size) {
		return false
	}

	for i := 2; i < int(size+2); i++ {
		p.Buffer[p.readPos] = _value[i]
		p.readPos += 1
	}

	p.MsgSize += size

	return true
}

func (p *Packet) ToString() string {
	c := p.Buffer
  n := -1
  for i, b := range c {
      if b == 0 {
          break
      }
      n = i
  }
  return string(c[:n+1])
}
