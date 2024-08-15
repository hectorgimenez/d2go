package packet

import "encoding/binary"

type PacketInterceptor struct {
	Sniffer *PacketSniffer
}

func NewPacketInterceptor(sniffer *PacketSniffer) *PacketInterceptor {
	return &PacketInterceptor{
		Sniffer: sniffer,
	}
}

func (pi *PacketInterceptor) OnPacketReceive(data []byte, handlerOffset uint32, isDynamicLength bool) {
	if len(data) < 1 {
		return
	}

	opCode := data[0]
	packet := SnifferPacket{
		IsIncoming:      true,
		OpCode:          opCode,
		Name:            GetAckName(opCode),
		Length:          len(data),
		Data:            data,
		HandlerOffset:   handlerOffset,
		IsDynamicLength: isDynamicLength,
	}

	pi.Sniffer.AddToQueue(packet)
}

func (pi *PacketInterceptor) OnPacketSend(data []byte, retAdr uint32, callStack string) {
	if len(data) < 1 {
		return
	}

	opCode := data[0]
	packet := SnifferPacket{
		IsIncoming: false,
		OpCode:     opCode,
		Name:       GetReqName(opCode),
		Length:     len(data),
		Data:       data,
		RetAdr:     retAdr,
		CallStack:  callStack,
	}

	pi.Sniffer.AddToQueue(packet)
}

func (pi *PacketInterceptor) ParsePacket(packet *SnifferPacket) {
	var opInfo OpCodeInfo
	if packet.IsIncoming {
		opInfo = OpCodesAck[packet.OpCode]
	} else {
		opInfo = OpCodesReq[packet.OpCode]
	}

	packet.Fields = make(map[string]interface{})
	offset := 1 // Start after opcode

	for _, field := range opInfo.Fields {
		if field.Size == -1 {
			// Handle variable-length field
			// This would depend on the specific packet structure
			continue
		}
		switch field.Size {
		case 1:
			packet.Fields[field.Name] = packet.Data[offset]
		case 2:
			packet.Fields[field.Name] = binary.LittleEndian.Uint16(packet.Data[offset : offset+2])
		case 4:
			packet.Fields[field.Name] = binary.LittleEndian.Uint32(packet.Data[offset : offset+4])
			// Add more cases as needed
		}
		offset += field.Size
	}
}
