package gomem

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

func TestGetProcessFromName(t *testing.T) {
	name := executableName()

	process, err := GetProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	if process.ID == 0 {
		t.Errorf("unexpected process id")
	}

	if process.Name != name {
		t.Errorf("unexpected process name")
	}
}

func TestProcessOpen(t *testing.T) {
	name := executableName()

	process, _ := GetProcessFromName(name)

	handle, err := process.Open()

	if err != nil {
		t.Errorf(err.Error())
	}

	if handle == 0 {
		t.Errorf("unexpected handle id")
	}
}

func TestProcessReadByte(t *testing.T) {
	name := executableName()

	var value = (byte)(0x42)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadByte(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadUInt32(t *testing.T) {
	name := executableName()

	var value = (uint32)(42)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadUInt32(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadUInt64(t *testing.T) {
	name := executableName()

	var value = (uint64)(42)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadUInt64(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadFloat32(t *testing.T) {
	name := executableName()

	var value = (float32)(42.0)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadFloat32(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadFloat64(t *testing.T) {
	name := executableName()

	var value = (float64)(42.0)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadFloat64(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadString16(t *testing.T) {
	name := executableName()

	var value = [16]byte{1, 2}
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadString16(valuePtr)
	fmt.Println(assertValue)

	if err != nil {
		t.Errorf(err.Error())
	}

	if string(value[:]) != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessWriteByte(t *testing.T) {
	name := executableName()

	var (
		value    = (byte)(0x42)
		valuePtr = (uintptr)(unsafe.Pointer(&value))
		newValue = (byte)(0x43)
	)

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	err = process.WriteByte(valuePtr, newValue)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != newValue {
		t.Errorf("unexpected value")
	}
}

func TestGetModuleNotFound(t *testing.T) {
	name := executableName()

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	ptr, err := process.GetModule("unknown.dll")

	if err.Error() != "module not found" {
		t.Errorf(err.Error())
	}

	if (ptr) == 0 {
		t.Errorf("unexpected value")
	}
}

func TestIsKeyDown(t *testing.T) {
	value := IsKeyDown(0x20) // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes

	if value != false {
		t.Errorf("unexpected value")
	}
}

func executableName() string {
	path, _ := os.Executable()

	return filepath.Base(path)
}
