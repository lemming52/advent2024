# Advent of Code 2024 Solutions

In search of wandering historians.

My set of solutions for [Advent of Code 2024](https://adventofcode.com/2024), written in Golang because _I want yet more practice_.

## Running

My inputs from the challenges are all stored in the `inputs` directory, and at the time of writing these files are effectively hardcoded into the running.

To run a particular day (i.e. _?_, ...) use either the name or the day
```sh
go run main.go -challenge hysteria
go run main.go -challenge 1
```

To run all days
```
go run main.go -all
```

### Challenge Days

Day | Challenge |Day | Challenge
----|-----------|----|----------
1 | `hysteria` | 14 | ``
2 | `rednosereports` | 15 | ``
3 | `mullitover` | 16 | ``
4 | `ceressearch` | 17 | ` `
5 | `printqueue` | 18 | ` `
6 | `` | 19 | ` `
7 | `` | 20 | ` `
8 | `` | 21 | ` `
9 | `` | 22 | ` `
10 | `` | 23 | ` `
11 | `` | 24 | ` `
12 | `` | 25 | ` `
13 | ``

### Adding new template

To template out a new day, from the root directory run
```sh
python3 scripts/template.py <day_number> <challenge_name>
```
