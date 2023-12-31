# Advent of code 2023

[Advent of Code](https://adventofcode.com/2023/about) is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

This year I'm trying to solve the puzzles in Go. I'm not a Go developer, so I'm using this as an opportunity to learn the language deeper. Sorry if the code is not very beautiful/idiomatic 🤪.

## How to run

```bash
cd internal/tasks/01
input.txt | go run main.go
```

## Benchmarking

You need to have [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) installed.

```bash
cd internal/tasks/01
go test -bench=. -count 10 | tee new.txt
benchstat old.txt new.txt
```
