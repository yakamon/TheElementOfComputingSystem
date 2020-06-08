package symbol

import "fmt"

type Address int

const (
	start Address = 16
)

func (a *Address) String() string {
	return fmt.Sprintf("%15b", a)
}

type Table struct {
	table       map[string]int
	nextAddress Address
}

func NewTable() *Table {
	return &Table{map[string]int{}, start}
}

func (t *Table) Add(symbol string, address int) {
	t.table[symbol] = address
}

func (t *Table) Has(symbol string) bool {
	_, ok := t.table[symbol]
	return ok
}

func (t *Table) Get(symbol string) int {
	address, ok := t.table[symbol]
	assert(ok, "non-existent symbol requested")
	return address
}

func assert(cond bool, msg string) {
	if !cond {
		panic("asm/parser internal error: " + msg)
	}
}
