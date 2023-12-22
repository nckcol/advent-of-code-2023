package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/nckcol/advent-of-code-2023/internal/utils/graph"
	"github.com/nckcol/advent-of-code-2023/internal/utils/input"
	"github.com/nckcol/advent-of-code-2023/internal/utils/list"
	"github.com/nckcol/advent-of-code-2023/internal/utils/numbers"
	"github.com/nckcol/advent-of-code-2023/internal/utils/stack"
	"github.com/nckcol/advent-of-code-2023/internal/utils/tokenizer"
)

func main() {
	input.EnsurePipeInput()
	lines, err := input.ScanLines()
	if err != nil {
		log.Fatal(err)
	}

	lookup := make(ModuleLookup, 0)

	for _, line := range lines {
		tokens, err := tokenizer.Tokenize(
			line,
			tokenizer.WithSeparators([]byte{','}),
			tokenizer.WithOperators([]string{"->", "%", "&"}),
		)

		if err != nil {
			log.Fatal(err)
		}

		operators := stack.New[string]()
		modules := stack.New[string]()
		var module *graph.NodeN[Module]

		for _, token := range tokens {
			switch {
			case token.Key == tokenizer.OPERATOR && token.Value == "->":
				moduleName, ok := modules.Pop()
				if !ok {
					log.Fatal("Unexpected end of line")
				}
				module = lookup.Ensure(moduleName)
				moduleType, _ := operators.Pop()
				switch moduleType {
				case "%":
					module.Value.Type = MODULE_TYPE_FLIP_FLOP
				case "&":
					module.Value.Type = MODULE_TYPE_CONJUNCTION
				}
			case token.Key == tokenizer.OPERATOR:
				operators.Push(token.Value)
			case token.Key == tokenizer.WORD:
				modules.Push(token.Value)
			}
		}

		for modules.Len() > 0 {
			moduleName, _ := modules.Pop()
			nextModule := lookup.Ensure(moduleName)
			nextModule.Value.Inputs[module.Value.Name] = MODULE_PULSE_LOW
			module.Children = append(module.Children, nextModule)
		}

		if operators.Len() > 0 {
			log.Fatal("Unexpected end of line")
		}
	}

	configuration := lookup.Get("broadcaster")

	// fmt.Println("First button push:")
	// runAll(configuration, "button", MODULE_PULSE_LOW, func(source string, input ModulePulse, name string) {
	// 	fmt.Printf("%s -%s-> %s\n", source, input, name)
	// })

	// fmt.Println()
	// fmt.Println("Second button push:")
	// runAll(configuration, "button", MODULE_PULSE_LOW, func(source string, input ModulePulse, name string) {
	// 	fmt.Printf("%s -%s-> %s\n", source, input, name)
	// })

	// fmt.Println()
	// fmt.Println("Third button push:")
	// runAll(configuration, "button", MODULE_PULSE_LOW, func(source string, input ModulePulse, name string) {
	// 	fmt.Printf("%s -%s-> %s\n", source, input, name)
	// })

	// fmt.Println()
	// fmt.Println("Fourth button push:")
	// runAll(configuration, "button", MODULE_PULSE_LOW, func(source string, input ModulePulse, name string) {
	// 	fmt.Printf("%s -%s-> %s\n", source, input, name)
	// })

	// lowCount := 0
	// highCount := 0
	// calc := func(source string, input ModulePulse, name string) {
	// 	if input == MODULE_PULSE_LOW {
	// 		lowCount += 1
	// 	} else {
	// 		highCount += 1
	// 	}
	// }
	// for i := 0; i < 1000000; i++ {
	// 	runAll(configuration, "button", MODULE_PULSE_LOW, calc)
	// }
	// fmt.Println(lowCount * highCount)

	// rx := lookup.Ensure("rx")
	// fmt.Println("rx:", keys(rx.Value.Inputs))

	// zh := lookup.Ensure("zh")
	// fmt.Println("zh:", keys(zh.Value.Inputs))

	vdCycle := 0
	nsCycle := 0
	bhCycle := 0
	dlCycle := 0
	for i := 0; i < 1000000; i++ {
		count := make(map[string]int)

		runAll(configuration, "button", MODULE_PULSE_LOW, func(source string, input ModulePulse, name string) {
			if input == MODULE_PULSE_HIGH && name == "zh" {
				count[source] += 1
			}
		})
		if count["vd"] == 1 {
			vdCycle = i + 1
		}
		if count["ns"] == 1 {
			nsCycle = i + 1
		}
		if count["bh"] == 1 {
			bhCycle = i + 1
		}
		if count["dl"] == 1 {
			dlCycle = i + 1
		}
		if vdCycle > 0 && nsCycle > 0 && bhCycle > 0 && dlCycle > 0 {
			break
		}
	}

	fmt.Printf("vd: %v, ns: %v, bh: %v, dl: %v\n", vdCycle, nsCycle, bhCycle, dlCycle)
	fmt.Println(numbers.LcmSlice([]int{vdCycle, nsCycle, bhCycle, dlCycle}))
}

func runAll(configuration *graph.NodeN[Module], source string, input ModulePulse, processPulse func(string, ModulePulse, string)) {
	type step struct {
		string
		*graph.NodeN[Module]
		ModulePulse
	}
	processQueue := make([]step, 0)
	processQueue = append(processQueue, step{source, configuration, input})

	for len(processQueue) > 0 {
		source := processQueue[0].string
		module := processQueue[0].NodeN
		input := processQueue[0].ModulePulse
		processQueue = processQueue[1:]

		processPulse(source, input, module.Value.Name)

		output, ok := runModule(module, source, input)
		if !ok {
			continue
		}

		next := list.Map(module.Children, func(node *graph.NodeN[Module]) step {
			return step{module.Value.Name, node, output}
		})
		slices.Reverse(next)
		processQueue = append(processQueue, next...)
	}
}

func runModule(module *graph.NodeN[Module], source string, input ModulePulse) (ModulePulse, bool) {
	switch module.Value.Type {
	case MODULE_TYPE_BROADCAST:
		return input, true

	case MODULE_TYPE_FLIP_FLOP:
		if input == MODULE_PULSE_HIGH {
			return MODULE_PULSE_LOW, false
		}
		if module.Value.State == MODULE_PULSE_LOW {
			module.Value.State = MODULE_PULSE_HIGH
		} else {
			module.Value.State = MODULE_PULSE_LOW
		}
		return module.Value.State, true

	case MODULE_TYPE_CONJUNCTION:
		module.Value.Inputs[source] = input
		for _, value := range module.Value.Inputs {
			if value == MODULE_PULSE_LOW {
				return MODULE_PULSE_HIGH, true
			}
		}
		return MODULE_PULSE_LOW, true
	}

	return MODULE_PULSE_LOW, false
}

// func keys[T any](m map[string]T) []string {
// 	result := make([]string, 0)
// 	for key := range m {
// 		result = append(result, key)
// 	}
// 	return result
// }
