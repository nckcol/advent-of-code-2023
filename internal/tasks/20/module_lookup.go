package main

import "github.com/nckcol/advent-of-code-2023/internal/utils/graph"

type ModuleLookup map[string]*graph.NodeN[Module]

func (l ModuleLookup) Get(name string) *graph.NodeN[Module] {
	return l[name]
}

func (l ModuleLookup) Ensure(name string) *graph.NodeN[Module] {
	if l[name] == nil {
		l[name] = &graph.NodeN[Module]{Value: Module{Name: name, Inputs: make(map[string]ModulePulse)}}
	}

	return l[name]
}
