// Package vm contains code related to the virtual machine that we created
// starting on day 4
package vm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

// VM is an object representing the virtual machine
type VM struct {
	Instructions []string
	Addr         int
	Acc          int
	Visited      map[int]bool
}

// NewVM creates a new VM from an input file of instructions
func NewVM(file string) *VM {
	instructions := input.Slice(file)
	return &VM{instructions, 0, 0, make(map[int]bool)}
}

// FromSlice returns a new VM created using the instructions passed in
func FromSlice(instructions []string) *VM {
	return &VM{instructions, 0, 0, make(map[int]bool)}
}

// Copy returns a copy of the VM
func (vm *VM) Copy() *VM {
	instructions := make([]string, len(vm.Instructions))
	copy(instructions, vm.Instructions)
	visited := make(map[int]bool, len(vm.Visited))
	return &VM{instructions, vm.Addr, vm.Acc, visited}
}

// Execute the instructions within the VM and return whether the program succesfully finished
func (vm *VM) Execute() (valid bool) {
	for vm.Addr < len(vm.Instructions) {
		if vm.Visited[vm.Addr] {
			return false
		}
		vm.Visited[vm.Addr] = true

		parts := strings.Split(vm.Instructions[vm.Addr], " ")
		value, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "nop":
			vm.Addr++
		case "acc":
			vm.Acc += value
			vm.Addr++
		case "jmp":
			vm.Addr += value
		}
	}
	return true
}

// Print visualizes the current state of the program
func (vm *VM) Print() {
	fmt.Printf("Accumulator: %v\n", vm.Acc)
	if vm.Addr > 0 {
		fmt.Printf("%v:  %v\n", vm.Addr-1, vm.Instructions[vm.Addr-1])
	}
	fmt.Printf("%v:* %v\n", vm.Addr, vm.Instructions[vm.Addr])
	if vm.Addr < len(vm.Instructions)-1 {
		fmt.Printf("%v:  %v\n", vm.Addr+1, vm.Instructions[vm.Addr+1])
	}
	fmt.Println("-------------------------")
}
