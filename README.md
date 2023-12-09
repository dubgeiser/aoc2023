# Advent Of Code 2023
My solutions for Advent Of Code 2023

This year, I'm trying Go... a whole new/different programming language for me.

Code is up for grabs, public domain.

Every day will be a directory, with 1 `main.go` file that has 2 functions, `part1` and `part2`.
I cannot explain where the inspiration comes from.

Puzzles should normally be found on https://adventofcode.com/2023/


## TODO / REVISIT

### Day 3
Solution looks too complicated, can probably done a lot more concise

### Day 5, part 2
Took the easy route and let the program run for a couple of minutes... Answer was found < 3', I'll take it for now :-D.

Think it should be rewritten with calculating the min. location for ranges or something (haven't yet thought about it)

### Day 7
Very tired, keeping it "simple"... at least for that little part of my brain that is still somewhat functioning.
I feel that there's got to be some simple nifty bitwise operator stuff in there to solve this though.
After all; card games have been programmed in home computers with 2KB memory and all...


## Notes

### Day 6
While parsing either the times or distances depending on `i` being 0 or 1, I took the simple way with a conditional.
times and distances could also be referenced so that `i` can be used as an index to assign to either times or distances.
Not convinced that this is necessary here, but it's an idiom worth keeping in mind, I feel.
```Go
for _, sn := range sNrs {
    if n, err := strconv.Atoi(sn); err == nil {
        if i == 0 {
            s.times = append(s.times, n)
        } else {
            s.distances = append(s.distances, n)
        }
    }
}

// VERSUS

targets := [2]*[]int{&s.times, &s.distances}
for _, sn := range sNrs {
    if n, err := strconv.Atoi(sn); err == nil {
        *targets[i] = append(*targets[i], n)
```
