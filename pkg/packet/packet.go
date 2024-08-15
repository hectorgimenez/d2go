package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type Packet struct {
	OpCode      uint8
	IsEncrypted bool
	Name        string
	Data        *bytes.Buffer
}

func NewPacket(isEncrypted bool, opCode uint8, name string) *Packet {
	p := &Packet{
		OpCode:      opCode,
		IsEncrypted: isEncrypted,
		Name:        name,
		Data:        new(bytes.Buffer),
	}

	// Write the opcode as the first byte
	p.Data.WriteByte(opCode)

	return p
}

func (p *Packet) FromArray(arr []byte) *Packet {
	p.Data.Write(arr)
	return p
}

func (p *Packet) ToString() string {
	var result string
	data := p.Data.Bytes()
	for i, b := range data {
		result += fmt.Sprintf("%02X ", b)
		if (i+1)%16 == 0 {
			result += "\n"
		}
	}
	return result
}

func (p *Packet) Send(conn net.Conn) error {
	_, err := conn.Write(p.Data.Bytes())
	return err
}

func (p *Packet) Receive(conn net.Conn) error {
	_, err := p.Data.ReadFrom(conn)
	return err
}

// Helper methods to write different data types to the packet
func (p *Packet) PutByte(v uint8) {
	p.Data.WriteByte(v)
}

func (p *Packet) PutShort(v uint16) {
	binary.Write(p.Data, binary.LittleEndian, v)
}

func (p *Packet) PutInt(v uint32) {
	binary.Write(p.Data, binary.LittleEndian, v)
}

func (p *Packet) PutString(s string) {
	p.Data.WriteString(s)
	p.Data.WriteByte(0) // Null-terminate the string
}

// Helper methods to read different data types from the packet
func (p *Packet) GetByte() uint8 {
	b, _ := p.Data.ReadByte()
	return b
}

func (p *Packet) GetShort() uint16 {
	var v uint16
	binary.Read(p.Data, binary.LittleEndian, &v)
	return v
}

func (p *Packet) GetInt() uint32 {
	var v uint32
	binary.Read(p.Data, binary.LittleEndian, &v)
	return v
}

func (p *Packet) GetString() string {
	var result []byte
	for {
		b, err := p.Data.ReadByte()
		if err != nil || b == 0 {
			break
		}
		result = append(result, b)
	}
	return string(result)
}
