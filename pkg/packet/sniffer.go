package packet

import (
	"sync"
)

type PacketSniffer struct {
	Queue           []SnifferPacket
	Listen          bool
	QueueLimit      int
	OpCodeBlackList map[uint8]bool
	mu              sync.Mutex
}

type SnifferPacket struct {
	IsIncoming      bool
	OpCode          uint8
	Name            string
	Length          int
	Data            []byte
	RetAdr          uint32
	CallStack       string
	HandlerOffset   uint32
	IsDynamicLength bool
	Fields          map[string]interface{} // To store parsed field values
}

func NewPacketSniffer() *PacketSniffer {
	return &PacketSniffer{
		Queue:           make([]SnifferPacket, 0),
		Listen:          false,
		QueueLimit:      500,
		OpCodeBlackList: make(map[uint8]bool),
	}
}

func (ps *PacketSniffer) Start() {
	ps.mu.Lock()
	ps.Listen = true
	ps.mu.Unlock()
}

func (ps *PacketSniffer) Stop() {
	ps.mu.Lock()
	ps.Listen = false
	ps.mu.Unlock()
}

func (ps *PacketSniffer) AddToQueue(packet SnifferPacket) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.Listen {
		return
	}

	if ps.OpCodeBlackList[packet.OpCode] {
		return
	}

	ps.Queue = append(ps.Queue, packet)
	if len(ps.Queue) > ps.QueueLimit {
		ps.Queue = ps.Queue[1:]
	}
}

func (ps *PacketSniffer) ClearQueue() {
	ps.mu.Lock()
	ps.Queue = make([]SnifferPacket, 0)
	ps.mu.Unlock()
}
