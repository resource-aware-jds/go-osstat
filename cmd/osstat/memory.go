package main

import (
	"github.com/resource-aware-jds/go-osstat/memory"
)

type memoryGenerator struct {
	memory *memory.Stats
	err    error
}

func (gen *memoryGenerator) Get() {
	gen.memory, gen.err = memory.Get()
}

func (gen *memoryGenerator) Error() error {
	return gen.err
}
