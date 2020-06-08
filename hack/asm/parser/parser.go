package parser

import (
	"bufio"
	"os"
)

type Pos int

type CommandType int

const (
	Address CommandType = iota
	Compute
	Symbol
)

type Parser struct {
	scanner *bufio.Scanner
	pos     Pos
}

func New(file *os.File) *Parser {
	scanner := bufio.NewScanner(file)
	return &Parser{scanner, 0}
}

func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) Advance() {

}

type (
	Command interface {
		Pos() Pos
		Type() CommandType
	}

	AddressCommand struct {
		Value string
	}

	ComputeCommand struct {
		Comp string
		Dest string
		Jump string
	}

	SymbolCommand struct {
		Symbol string
	}
)

func (a *AddressCommand) Type() CommandType {
	return Address
}

func (c *ComputeCommand) Type() CommandType {
	return Compute
}

func (l *SymbolCommand) Type() CommandType {
	return Symbol
}

func assert(cond bool, msg string) {
	if !cond {
		panic("asm/parser internal error: " + msg)
	}
}
