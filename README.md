# Jinli's AdventOfCode Adventure Time

Mostly Go, maybe Rust or something else in the future?

## Installation
```bash
cd golang
go mod tidy
go mod vendor
```

## Usage

Currently, the usage only looks at 2024 as I have not touched any other years.

This requires your own input.txt in each yr/day dir to work for the *actual* parts. For the prompt portions, this can
be executed without providing your own input text data.

```bash
# Run day 1 part 1 prompt
go run ./cmd/runner 1 part1 prompt
# Run day 1 part 1 actual after placing input.txt in yr2024/day01/ folder.
go run ./cmd/runner 1 part1 actual
# Run day 1 part 2 prompt
go run ./cmd/runner 2 part1 prompt
# Run day 1 part 2 actual after placing input.txt in yr2024/day01/ folder.
go run ./cmd/runner 1 part1 actual
```
