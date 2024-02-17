# Stack-Based Virtual Machine in Go

This is a simple implementation of a stack-based virtual machine (VM) in the Go programming language. The VM supports basic arithmetic, logical, and comparison operations.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.

### Running the Program

1. Clone this repository:

   ```bash
   git clone https://github.com/anouaressa/stack-based-vm.git

2. Change into the project directory:
    ```bash 
    cd stack-based-vm
3. ```bash 
    go run main.go
### Example Usage
The main program (main.go) demonstrates the usage of the stack-based virtual machine. Feel free to modify the instructions in the Run method to test different operations.

```
// Push values onto the stack
vm.Push(5)
vm.Push(10)

// Execute ADD instruction
err := vm.Run([]string{"ADD"})
if err != nil {
    log.Fatal("Error:", err)
}

// Execute MULTIPLY instruction
err = vm.Run([]string{"MULTIPLY"})
if err != nil {
    log.Fatal("Error:", err)
}

// Print the result
vm.Print()

```

# Available Operations

### Arithmetic Operations:
ADD (Addition)
SUB (Subtraction)
MULTIPLY (Multiplication)
DIV (Division)
MOD (Modulus)
Unary operations (e.g., negation, increment, decrement)
### Logical Operations:
LOGICAL_AND (Logical AND)
LOGICAL_OR (Logical OR)
LOGICAL_NOT (Logical NOT)
### Comparison Operations:
EQ (Equal to)
NEQ (Not equal to)
GT (Greater than)
LT (Less than)
GTE (Greater than or equal to)
LTE (Less than or equal to)

## Extending the VM
You can extend the virtual machine by adding more operations or modifying existing ones based on your requirements.

