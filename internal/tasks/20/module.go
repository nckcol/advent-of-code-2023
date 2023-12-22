package main

type ModuleType int

const (
	MODULE_TYPE_BROADCAST   ModuleType = 0
	MODULE_TYPE_FLIP_FLOP   ModuleType = 1
	MODULE_TYPE_CONJUNCTION ModuleType = 2
)

func (m ModuleType) String() string {
	switch m {
	case MODULE_TYPE_BROADCAST:
		return "B"
	case MODULE_TYPE_FLIP_FLOP:
		return "%"
	case MODULE_TYPE_CONJUNCTION:
		return "&"
	}
	return "?"
}

type ModulePulse int

const (
	MODULE_PULSE_LOW  ModulePulse = 0
	MODULE_PULSE_HIGH ModulePulse = 1
)

func (m ModulePulse) String() string {
	switch m {
	case MODULE_PULSE_LOW:
		return "low"
	case MODULE_PULSE_HIGH:
		return "high"
	}
	return "?"
}

type Module struct {
	Name   string
	Type   ModuleType
	State  ModulePulse
	Inputs map[string]ModulePulse
}
