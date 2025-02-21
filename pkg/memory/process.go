package memory

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const moduleName = "d2r.exe"

type Process struct {
	handler              windows.Handle
	pid                  uint32
	moduleBaseAddressPtr uintptr
	moduleBaseSize       uint32
}

const (
	Int8  = 1 // signed 8-bit integer
	Int16 = 2 // signed 16-bit integer
	Int32 = 4 // signed 32-bit integer
	Int64 = 8 // signed 64-bit integer
)

func NewProcess() (Process, error) {
	module, err := getGameModule()
	if err != nil {
		return Process{}, err
	}

	h, err := windows.OpenProcess(0x0010, false, module.ProcessID)
	if err != nil {
		return Process{}, err
	}

	return Process{
		handler:              h,
		pid:                  module.ProcessID,
		moduleBaseAddressPtr: module.ModuleBaseAddress,
		moduleBaseSize:       module.ModuleBaseSize,
	}, nil
}

func NewProcessForPID(pid uint32) (Process, error) {
	module, found := getMainModule(pid)
	if !found {
		return Process{}, errors.New("no module found for the specified PID")
	}

	h, err := windows.OpenProcess(0x0010, false, module.ProcessID)
	if err != nil {
		return Process{}, err
	}

	return Process{
		handler:              h,
		pid:                  module.ProcessID,
		moduleBaseAddressPtr: module.ModuleBaseAddress,
		moduleBaseSize:       module.ModuleBaseSize,
	}, nil
}

func (p Process) Close() error {
	return windows.CloseHandle(p.handler)
}

func getGameModule() (ModuleInfo, error) {
	processes := make([]uint32, 2048)
	length := uint32(0)
	err := windows.EnumProcesses(processes, &length)
	if err != nil {
		return ModuleInfo{}, err
	}

	for _, process := range processes {
		module, found := getMainModule(process)
		if found {
			return module, nil
		}
	}

	return ModuleInfo{}, err
}

func getMainModule(pid uint32) (ModuleInfo, bool) {
	mi, err := GetProcessModules(pid)
	if err != nil {
		return ModuleInfo{}, false
	}
	for _, m := range mi {
		if strings.Contains(strings.ToLower(m.ModuleName), moduleName) {
			return m, true
		}
	}

	return ModuleInfo{}, false
}

func (p Process) getProcessMemory() ([]byte, error) {
	var data = make([]byte, p.moduleBaseSize)
	err := windows.ReadProcessMemory(p.handler, p.moduleBaseAddressPtr, &data[0], uintptr(p.moduleBaseSize), nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p Process) ReadBytesFromMemory(address uintptr, size uint) []byte {
	var data = make([]byte, size)
	windows.ReadProcessMemory(p.handler, address, &data[0], uintptr(size), nil)

	return data
}

type IntType uint

const (
	Uint8  = 1
	Uint16 = 2
	Uint32 = 4
	Uint64 = 8
)

func (p Process) ReadUInt(address uintptr, size IntType) uint {
	bytes := p.ReadBytesFromMemory(address, uint(size))

	return bytesToUint(bytes, size)
}

func ReadUIntFromBuffer(bytes []byte, offset uint, size IntType) uint {
	return bytesToUint(bytes[offset:offset+uint(size)], size)
}

func bytesToUint(bytes []byte, size IntType) uint {
	switch size {
	case Uint8:
		return uint(bytes[0])
	case Uint16:
		return uint(binary.LittleEndian.Uint16(bytes))
	case Uint32:
		return uint(binary.LittleEndian.Uint32(bytes))
	case Uint64:
		return uint(binary.LittleEndian.Uint64(bytes))
	}

	return 0
}
func ReadIntFromBuffer(bytes []byte, offset uint, size IntType) int {
	return bytesToInt(bytes[offset:offset+uint(size)], size)
}
func bytesToInt(bytes []byte, size IntType) int {
	switch size {
	case Int8:
		return int(int8(bytes[0]))
	case Int16:
		return int(int16(binary.LittleEndian.Uint16(bytes)))
	case Int32:
		return int(int32(binary.LittleEndian.Uint32(bytes)))
	case Int64:
		return int(int64(binary.LittleEndian.Uint64(bytes)))
	}
	return 0
}

func (p Process) ReadStringFromMemory(address uintptr, size uint) string {
	if size == 0 {
		for i := 1; true; i++ {
			data := p.ReadBytesFromMemory(address, uint(i))
			if data[i-1] == 0 {
				return string(bytes.Trim(data, "\x00"))
			}
		}
	}

	return string(bytes.Trim(p.ReadBytesFromMemory(address, size), "\x00"))
}

func (p Process) findPattern(memory []byte, pattern, mask string) int {
	patternLength := len(pattern)
	for i := 0; i < int(p.moduleBaseSize)-patternLength; i++ {
		found := true
		for j := 0; j < patternLength; j++ {
			if string(mask[j]) != "?" && string(pattern[j]) != string(memory[i+j]) {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}

	return 0
}

func (p Process) FindPattern(memory []byte, pattern, mask string) uintptr {
	if offset := p.findPattern(memory, pattern, mask); offset != 0 {
		return p.moduleBaseAddressPtr + uintptr(offset)
	}

	return 0
}

func (p Process) FindPatternByOperand(memory []byte, pattern, mask string) uintptr {
	if offset := p.findPattern(memory, pattern, mask); offset != 0 {
		// Adjust the address based on the operand value
		operandAddress := p.moduleBaseAddressPtr + uintptr(offset)
		operandValue := binary.LittleEndian.Uint32(memory[offset+3 : offset+7])
		finalAddress := operandAddress + uintptr(operandValue) + 7 // 7 is the length of the instruction
		return finalAddress
	}

	return 0
}

func (p Process) GetPID() uint32 {
	return p.pid
}

type ModuleInfo struct {
	ProcessID         uint32
	ModuleBaseAddress uintptr
	ModuleBaseSize    uint32
	ModuleHandle      syscall.Handle
	ModuleName        string
}

func GetProcessModules(processID uint32) ([]ModuleInfo, error) {
	hProcess, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, processID)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(hProcess)

	var modules [1024]windows.Handle
	var needed uint32
	if err := windows.EnumProcessModules(hProcess, &modules[0], uint32(unsafe.Sizeof(modules[0]))*1024, &needed); err != nil {
		return nil, err
	}
	count := needed / uint32(unsafe.Sizeof(modules[0]))

	var moduleInfos []ModuleInfo
	for i := uint32(0); i < count; i++ {
		var mi windows.ModuleInfo
		if err := windows.GetModuleInformation(hProcess, modules[i], &mi, uint32(unsafe.Sizeof(mi))); err != nil {
			return nil, err
		}

		var moduleName [windows.MAX_PATH]uint16
		if err := windows.GetModuleFileNameEx(hProcess, modules[i], &moduleName[0], windows.MAX_PATH); err != nil {
			return nil, err
		}

		moduleInfos = append(moduleInfos, ModuleInfo{
			ProcessID:         processID,
			ModuleBaseAddress: mi.BaseOfDll,
			ModuleBaseSize:    mi.SizeOfImage,
			ModuleHandle:      syscall.Handle(modules[i]),
			ModuleName:        syscall.UTF16ToString(moduleName[:]),
		})
	}

	return moduleInfos, nil
}

// ReadPointer reads a pointer from the specified memory address.
func (p *Process) ReadPointer(address uintptr, size int) (uintptr, error) {
	buffer := p.ReadBytesFromMemory(address, uint(size))
	if len(buffer) == 0 {
		return 0, errors.New("failed to read memory")
	}

	return uintptr(*(*uint64)(unsafe.Pointer(&buffer[0]))), nil
}

func (p Process) ReadIntoBuffer(address uintptr, buffer []byte) error {
	return windows.ReadProcessMemory(p.handler, address, &buffer[0], uintptr(len(buffer)), nil)
}

// ReadWidgetContainer reads the WidgetContainer structure.
func (p *Process) ReadWidgetContainer(address uintptr, full bool) (map[string]interface{}, error) {
	widgetPtr, err := p.ReadPointer(address+0x8, 8)
	if err != nil {
		return nil, err
	}

	widgetNameLength := p.ReadUInt(address+0x10, 4)

	widgetName := p.ReadStringFromMemory(widgetPtr, uint(widgetNameLength))
	if widgetName == "" {
		return nil, errors.New("failed to read widget name")
	}

	widget_visible := p.ReadUInt(address+0x51, 1) == 1
	widget_active := p.ReadUInt(address+0x50, 1) == 1

	result := map[string]interface{}{
		"WidgetNameString": widgetName,
		"WidgetNameLength": widgetNameLength,
		"WidgetVisible":    widget_visible,
		"WidgetActive":     widget_active,
	}

	if full {
		childWidgetsListPtr, err := p.ReadPointer(widgetPtr+0x38, 8)
		if err != nil {
			return nil, err
		}

		childWidgetSize := p.ReadUInt(widgetPtr+0x40, 4)

		widgetListPtr, err := p.ReadPointer(widgetPtr+0x68, 8)
		if err != nil {
			return nil, err
		}

		widgetListSize := p.ReadUInt(widgetPtr+0x78, 4)

		widgetList2Ptr, err := p.ReadPointer(widgetPtr+0x80, 8)
		if err != nil {
			return nil, err
		}

		widgetList2Size := p.ReadUInt(widgetPtr+0x90, 4)

		result["ChildWidgetsListPointer"] = childWidgetsListPtr
		result["ChildWidgetSize"] = childWidgetSize
		result["WidgetListPointer"] = widgetListPtr
		result["WidgetListSize"] = widgetListSize
		result["WidgetList2Pointer"] = widgetList2Ptr
		result["WidgetList2Size"] = widgetList2Size
	}

	return result, nil
}

// ReadWidgetList iterates through a list of widgets given a pointer to the list and its size.
func (p *Process) ReadWidgetList(listPointer uintptr, listSize int) (map[string]map[string]interface{}, error) {
	widgetMap := make(map[string]map[string]interface{})
	widgetSize := int(unsafe.Sizeof(uintptr(0)))

	for i := 0; i < listSize; i++ {
		widgetAddr, err := p.ReadPointer(listPointer+uintptr(i*widgetSize), 8)
		if err != nil {
			return nil, err
		}

		widgetContainer, err := p.ReadWidgetContainer(widgetAddr, false)
		if err != nil {
			return nil, err
		}

		widgetName, ok := widgetContainer["WidgetNameString"].(string)
		if !ok {
			return nil, errors.New("failed to read widget name")
		}

		widgetMap[widgetName] = widgetContainer
	}

	return widgetMap, nil
}
