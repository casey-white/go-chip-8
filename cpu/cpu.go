package cpu

import (
	"go-chip-8/memory"
)

type ProgramCounter uint16
type Stack [16]uint16
type StackPointer uint8
type Registers [16]uint8

type PCAction interface {
	isPCAction()
}

type Next struct{}
type Skip struct{}
type StepBack struct{}
type Jump struct {
	postiton uint16
}

func (Next) isPCAction()     {}
func (Skip) isPCAction()     {}
func (StepBack) isPCAction() {}
func (j Jump) isPCAction()   {}

type CPU struct {
	Mem            memory.Memory
	Registers      Registers
	ProgramCounter ProgramCounter
	StackPointer   StackPointer
	Stack          Stack
}

func NewCPU(filename string) *CPU {
	return &CPU{
		Mem:            memory.NewMemory(filename),
		Registers:      Registers{},
		ProgramCounter: 0x200,
		StackPointer:   0,
		Stack:          Stack{},
	}
}

func Reset(c *CPU) {
	c.ProgramCounter = 0x200
	c.StackPointer = 0
	c.Stack = Stack{}
	c.Registers = Registers{}
}
