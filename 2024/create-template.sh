#/bin/bash

SESSION=$(cat session.txt)
YEAR=2024

mkdir $1 

curl -s -b "session=$SESSION" "https://adventofcode.com/$YEAR/day/$1/input" -o $1/input.txt

touch $1/solve.go && cat template.go > $1/solve.go
