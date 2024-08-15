package packet

import (
	"encoding/binary"
	"fmt"
)

// Opcode constant for InteractWithUnit
const OpCodeInteractWithUnit = 19

// InteractWithUnit creates a packet for interacting with a unit in the game
func InteractWithUnit(targetUnitId, executeeUnitId uint32) (*Packet, error) {
	opCodeInfo, exists := OpCodesReq[OpCodeInteractWithUnit]
	if !exists {
		return nil, fmt.Errorf("opcode %d for InteractWithUnit not found in OpCodesReq", OpCodeInteractWithUnit)
	}

	p := NewPacket(false, uint8(OpCodeInteractWithUnit), opCodeInfo.Name)

	// Ensure the packet size matches the expected size
	expectedSize := opCodeInfo.Size
	if expectedSize != 9 { // 1 byte opcode + 4 bytes targetUnitId + 4 bytes executeeUnitId
		return nil, fmt.Errorf("unexpected packet size for InteractWithUnit: got %d, want 9", expectedSize)
	}

	// Write the target unit ID (4 bytes)
	if err := binary.Write(p.Data, binary.LittleEndian, targetUnitId); err != nil {
		return nil, fmt.Errorf("failed to write targetUnitId: %w", err)
	}

	// Write the executee unit ID (4 bytes)
	if err := binary.Write(p.Data, binary.LittleEndian, executeeUnitId); err != nil {
		return nil, fmt.Errorf("failed to write executeeUnitId: %w", err)
	}

	return p, nil
}

// ParseInteractWithUnit parses a received InteractWithUnit packet
func ParseInteractWithUnit(p *Packet) (targetUnitId, executeeUnitId uint32, err error) {
	if p.OpCode != OpCodeInteractWithUnit {
		return 0, 0, fmt.Errorf("invalid opcode for InteractWithUnit: got %d, want %d", p.OpCode, OpCodeInteractWithUnit)
	}

	// Read target unit ID
	if err := binary.Read(p.Data, binary.LittleEndian, &targetUnitId); err != nil {
		return 0, 0, fmt.Errorf("failed to read targetUnitId: %w", err)
	}

	// Read executee unit ID
	if err := binary.Read(p.Data, binary.LittleEndian, &executeeUnitId); err != nil {
		return 0, 0, fmt.Errorf("failed to read executeeUnitId: %w", err)
	}

	return targetUnitId, executeeUnitId, nil
}
