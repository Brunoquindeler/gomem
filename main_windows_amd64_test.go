package main

import (
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

func TestSanity(t *testing.T) {
	main()
}

func TestGetFromProcessName(t *testing.T) {
	name := executableName()

	process, err := GetFromProcessName(name)

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

	process, _ := GetFromProcessName(name)

	handle, err := process.Open()

	if err != nil {
		t.Errorf(err.Error())
	}

	if handle == 0 {
		t.Errorf("unexpected handle id")
	}
}

func TestProcessRead(t *testing.T) {
	name := executableName()
	
	testInt32 := 42
	offset := (uintptr)(unsafe.Pointer(&testInt32))

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}
	
	var buffer uintptr

	process.Open()
	ptr, err := process.Read(offset, buffer, 4)

	if err != nil {
		t.Errorf(err.Error())
	}

	if (int)(*ptr) != testInt32 {
		t.Errorf(err.Error())
	}
}

func executableName() string {
	path, _ := os.Executable()

	return filepath.Base(path)
}
