package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	Size  int
}

type RangeMap struct {
	SourceStart      int
	DestinationStart int
	Size             int
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("You should pipe input to stdin.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	chunk := make([]string, 0)
	seedRanges := make([]Range, 0)

	var (
		seedToSoil            []RangeMap
		soilToFertilizer      []RangeMap
		fertilizerToWater     []RangeMap
		waterToLight          []RangeMap
		lightToTemperature    []RangeMap
		temperatureToHumidity []RangeMap
		humidityToLocation    []RangeMap
	)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			chunk = append(chunk, line)
			continue
		}

		if strings.HasPrefix(chunk[0], "seeds:") {
			parsedSeeds := strings.Split(chunk[0], " ")
			for i := 1; i < len(parsedSeeds); i += 2 {
				start, err := strconv.Atoi(parsedSeeds[i])
				if err != nil {
					log.Fatalln("Cannot parse input", err)
				}
				size, err := strconv.Atoi(parsedSeeds[i+1])
				seedRanges = append(seedRanges, Range{
					Start: start, Size: size})
			}
			slices.SortFunc(seedRanges, compareRanges)
		} else if chunk[0] == "seed-to-soil map:" {
			seedToSoil = parseMapChunk(chunk)
		} else if chunk[0] == "soil-to-fertilizer map:" {
			soilToFertilizer = parseMapChunk(chunk)
		} else if chunk[0] == "fertilizer-to-water map:" {
			fertilizerToWater = parseMapChunk(chunk)
		} else if chunk[0] == "water-to-light map:" {
			waterToLight = parseMapChunk(chunk)
		} else if chunk[0] == "light-to-temperature map:" {
			lightToTemperature = parseMapChunk(chunk)
		} else if chunk[0] == "temperature-to-humidity map:" {
			temperatureToHumidity = parseMapChunk(chunk)
		} else {
			log.Fatalln("Cannot parse input")
		}

		chunk = make([]string, 0)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	humidityToLocation = parseMapChunk(chunk)

	// fmt.Println(seedRanges)

	mapSequence := [][]RangeMap{
		seedToSoil,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation,
	}

	current := seedRanges

	for _, currentMap := range mapSequence {
		current = mapValues(current, currentMap)
		// fmt.Println("Done ", i, ": ", current)
	}

	fmt.Println(len(current))
	fmt.Println(current[0])
}

func parseMapChunk(chunk []string) []RangeMap {
	result := make([]RangeMap, 0, len(chunk)-2)
	for i := 1; i < len(chunk); i++ {
		if curentRange, err := parseRange(chunk[i]); err == nil {
			result = append(result, curentRange)
		} else {
			log.Fatalln("Cannot parse input", err)
		}
	}
	slices.SortFunc(result, compareRangeMaps)
	return result
}

func parseRange(source string) (RangeMap, error) {
	parsedRange := strings.Split(source, " ")
	currentRange := RangeMap{}
	var err error
	currentRange.DestinationStart, err = strconv.Atoi(parsedRange[0])
	if err != nil {
		return RangeMap{}, err
	}
	currentRange.SourceStart, err = strconv.Atoi(parsedRange[1])
	if err != nil {
		return RangeMap{}, err
	}
	currentRange.Size, err = strconv.Atoi(parsedRange[2])
	if err != nil {
		return RangeMap{}, err
	}
	return currentRange, nil
}

func compareRanges(a, b Range) int {
	return a.Start - b.Start
}

func compareRangeMaps(a, b RangeMap) int {
	return a.SourceStart - b.SourceStart
}

func mapValues(ranges []Range, currentMap []RangeMap) []Range {
	boundaries := make([]int, 0, len(currentMap)*2)
	for _, m := range currentMap {
		boundaries = append(boundaries, m.SourceStart, m.SourceStart+m.Size-1)
	}

	var subRanges []Range
	for _, r := range ranges {
		subRanges = append(subRanges, r.Split(boundaries...)...)
	}

	j := 0
	for i := 0; i < len(subRanges); i++ {
		r := &subRanges[i]
		for j+1 < len(currentMap) && currentMap[j+1].SourceStart <= r.Start {
			j += 1
		}
		m := currentMap[j]

		if r.Start+r.Size-1 < m.SourceStart || r.Start > m.SourceStart+m.Size-1 {
			continue
		}

		r.Start = (r.Start - m.SourceStart) + m.DestinationStart
	}
	slices.SortFunc(subRanges, compareRanges)
	return subRanges
}

func (r Range) Split(boundaries ...int) []Range {
	var result []Range
	slices.Sort(boundaries)
	for _, boundary := range boundaries {
		if boundary <= r.Start {
			continue
		}
		if boundary >= r.Start+r.Size {
			break
		}
		result = append(result, Range{Start: r.Start, Size: boundary - r.Start})
		r = Range{Start: boundary, Size: r.Size + r.Start - boundary}
	}
	return append(result, r)
}
