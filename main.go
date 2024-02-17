package main

import (
	"errors"
	"fmt"
	"log"
)

// StackVM represents a simple stack-based virtual machine.
type StackVM struct {
	stack []int
}

// NewStackVM initializes a new StackVM.
func NewStackVM() *StackVM {
	return &StackVM{}
}

// Push value to the stack.
func (vm *StackVM) Push(value int) {
	vm.stack = append(vm.stack, value)
}

// remove and returns the top value from the stack.
func (vm *StackVM) Pop() (int, error) {
	if len(vm.stack) == 0 {
		return 0, errors.New("stack underflow")
	}
	top := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return top, nil
}

// Add pops the top two values from the stack, adds them, and pushes the result back.
func (vm *StackVM) Add() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for add operation")
	}
	a, err := vm.Pop()
	if err != nil {
		return err
	}
	b, err := vm.Pop()
	if err != nil {
		return err
	}
	result := a + b
	vm.Push(result)
	return nil
}

// Multiply pops the top two values from the stack, multiplies them, and pushes the result back.
func (vm *StackVM) Multiply() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for multiply operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := a * b
	vm.Push(result)
	return nil
}

// Subtraction pops the top two values from the stack, subtracts the second from the first, and pushes the result back.
func (vm *StackVM) Subtraction() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for subtraction operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := a - b
	vm.Push(result)
	return nil
}

// Division pops the top two values from the stack, divides the first by the second, and pushes the result back.
func (vm *StackVM) Division() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for division operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	if b == 0 {
		return errors.New("division by zero")
	}

	result := a / b
	vm.Push(result)
	return nil
}

// Modulus pops the top two values from the stack, calculates the modulus of the first by the second, and pushes the result back.
func (vm *StackVM) Modulus() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for modulus operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	if b == 0 {
		return errors.New("modulus by zero")
	}

	result := a % b
	vm.Push(result)
	return nil
}

// Negation pops the top value from the stack, negates it, and pushes the result back.
func (vm *StackVM) Negation() error {
	if len(vm.stack) < 1 {
		return errors.New("not enough operands for negation operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	result := -a
	vm.Push(result)
	return nil
}

// Increment increments the top value on the stack.
func (vm *StackVM) Increment() error {
	if len(vm.stack) < 1 {
		return errors.New("not enough operands for increment operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	result := a + 1
	vm.Push(result)
	return nil
}

// Decrement decrements the top value on the stack.
func (vm *StackVM) Decrement() error {
	if len(vm.stack) < 1 {
		return errors.New("not enough operands for decrement operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	result := a - 1
	vm.Push(result)
	return nil
}

// EqualTo pops the top two values from the stack, checks if they are equal, and pushes the result back.(0 or 1)
func (vm *StackVM) EqualTo() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for equal to operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := 0
	if a == b {
		result = 1
	}

	vm.Push(result)
	return nil
}

// NotEqualTo pops the top two values from the stack, checks if they are not equal, and pushes the result back.
func (vm *StackVM) NotEqualTo() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for not equal to operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := 0
	if a != b {
		result = 1
	}

	vm.Push(result)
	return nil
}

// GreaterThan pops the top two values from the stack, checks if the first is greater than the second, and pushes the result back.
func (vm *StackVM) GreaterThan() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for greater than operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := 0
	if a > b {
		result = 1
	}

	vm.Push(result)
	return nil
}

// LessThan pops the top two values from the stack, checks if the first is less than the second, and pushes the result back.
func (vm *StackVM) LessThan() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for less than operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := 0
	if a < b {
		result = 1
	}

	vm.Push(result)
	return nil
}

// GreaterThanOrEqual pops the top two values from the stack, checks if the first is greater than or equal to the second, and pushes the result back.
func (vm *StackVM) GreaterThanOrEqual() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for greater than or equal to operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := 0
	if a >= b {
		result = 1
	}

	vm.Push(result)
	return nil
}

// LessThanOrEqual pops the top two values from the stack, checks if the first is less than or equal to the second, and pushes the result back.
func (vm *StackVM) LessThanOrEqual() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for less than or equal to operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := 0
	if a <= b {
		result = 1
	}

	vm.Push(result)
	return nil
}

// LogicalAnd pops the top two values from the stack, performs logical AND, and pushes the result back.
func (vm *StackVM) LogicalAnd() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for logical AND operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := a & b
	vm.Push(result)
	return nil
}

// LogicalOr pops the top two values from the stack, performs logical OR, and pushes the result back.
func (vm *StackVM) LogicalOr() error {
	if len(vm.stack) < 2 {
		return errors.New("not enough operands for logical OR operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	b, err := vm.Pop()
	if err != nil {
		return err
	}

	result := a | b
	vm.Push(result)
	return nil
}

// LogicalNot pops the top value from the stack, performs logical NOT, and pushes the result back.
func (vm *StackVM) LogicalNot() error {
	if len(vm.stack) < 1 {
		return errors.New("not enough operands for logical NOT operation")
	}

	a, err := vm.Pop()
	if err != nil {
		return err
	}

	result := ^a
	vm.Push(result)
	return nil
}

// Print prints the top value of the stack.
func (vm *StackVM) Print() {
	if len(vm.stack) == 0 {
		fmt.Println("Stack is empty.")
	} else {
		fmt.Println("Top of the stack:", vm.stack[len(vm.stack)-1])
	}
}

// Run executes a sequence of instructions.
func (vm *StackVM) Run(instructions []string) error {
	for _, instr := range instructions {
		switch instr {
		case "ADD":
			err := vm.Add()
			if err != nil {
				return err
			}
		case "MULTIPLY":
			err := vm.Multiply()
			if err != nil {
				return err
			}
		case "SUB":
			err := vm.Subtraction()
			if err != nil {
				return err
			}
		case "DIV":
			err := vm.Division()
			if err != nil {
				return err
			}
		case "MOD":
			err := vm.Modulus()
			if err != nil {
				return err
			}
		case "NEG":
			err := vm.Negation()
			if err != nil {
				return err
			}
		case "INC":
			err := vm.Increment()
			if err != nil {
				return err
			}
		case "DEC":
			err := vm.Decrement()
			if err != nil {
				return err
			}

		case "EQ":
			err := vm.EqualTo()
			if err != nil {
				return err
			}
		case "NEQ":
			err := vm.NotEqualTo()
			if err != nil {
				return err
			}
		case "GT":
			err := vm.GreaterThan()
			if err != nil {
				return err
			}
		case "LT":
			err := vm.LessThan()
			if err != nil {
				return err
			}
		case "GTE":
			err := vm.GreaterThanOrEqual()
			if err != nil {
				return err
			}
		case "LTE":
			err := vm.LessThanOrEqual()
			if err != nil {
				return err
			}

		case "PRINT":
			vm.Print()

		case "LOGICAL_AND":
			err := vm.LogicalAnd()
			if err != nil {
				return err
			}
		case "LOGICAL_OR":
			err := vm.LogicalOr()
			if err != nil {
				return err
			}
		case "LOGICAL_NOT":
			err := vm.LogicalNot()
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("unknown instruction: %s", instr)
		}
	}
	return nil
}

func main() {
	// Example usage
	vm := NewStackVM()

	// Push values onto the stack
	vm.Push(5)
	vm.Push(10)
	vm.Push(0)

	// Execute MULTIPLY and ADD instructions
	err := vm.Run([]string{"MULTIPLY", "ADD"})
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Print the result(the top value of the stack)
	vm.Print() // Expected: 5     (10*0) ---> 0+5

	// Push values onto the stack
	vm.Push(0) // False
	vm.Push(1) // True

	// Execute LOGICAL_AND instruction
	err = vm.Run([]string{"LOGICAL_AND"})
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Print the result
	vm.Print() // Expected: 0 (False)

	vm.Push(5)
	vm.Print()
	vm.Push(10)
	vm.Print()

	// Execute comparison operations
	err = vm.Run([]string{"LT"})
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Print the results
	vm.Print()

	// Push values onto the stack
	vm.Push(10)
	vm.Push(5)

	// Execute SUB (subtraction) operation
	err = vm.Run([]string{"SUB", "INC", "DEC"})
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Print the result
	vm.Print()

}
