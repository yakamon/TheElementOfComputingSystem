package asm

import (
	"os"

	"github.com/yakamon/ossu-computer-science/TheElementsOfComputingSystem/hack/asm/parser"
	"github.com/yakamon/ossu-computer-science/TheElementsOfComputingSystem/hack/asm/symbol"
)

func Assemble(file *os.File) {
	parser := parser.New(file)
	symbolTable := symbol.NewTable()
}
