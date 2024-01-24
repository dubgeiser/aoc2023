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

### Day 4, part 2
Going by the need to loop over the cards first a second time after building the `Card` slice, suggests that there's probably an easier way to go about this.

### Day 5, part 2
Took the easy route and let the program run for a couple of minutes... Answer was found < 3', I'll take it for now :-D.

Think it should be rewritten with calculating the min. location for ranges or something (haven't yet thought about it)

### Day 7
Very tired, keeping it "simple"... at least for that little part of my brain that is still somewhat functioning.
I feel that there's got to be some simple nifty bitwise operator stuff in there to solve this though.
After all; card games have been programmed in home computers with 2KB memory and all...

### Day 9
I assume there a bunch of (ML) algo's for predicting the next element in a sequence

### Day 12
Wanted to brute force part 1 by "just trying all possible combinations"... which lead me to a serious running-around-in-circles-brainf*ck that I eventually broke through by introducing a multiple-base-number `decrement()` func.
Runs slow af though, takes around 5 seconds :-D
But lead to building a `Permutations()` func that I'll probably can refactor in a more re-usable lib func.

### Day 13
Bit ugly with the separate `CheckGrid()` call.
Part 2 took a while, because of reference to the slices in `isPerfectReflection()` and the `calcLines()` calculation was not sufficient to count the 1-off mirroring correctly, so I rewrote it from scratch...

### Day 17
Some refactoring can be done to remove the enormous code duplication between both parts.
part 1's 3 and part 2's 4 and 10 can probably be extracted in some min/max arguments.
Or the conditions (the RELAX() from Algo book on Dijkstra) can be passed as functors.

### Day 20
Design-wise, I'm not too happy with the general Module struct.
I'd much rather have FlipFlop and Conjuctions structs, that support the same interface Module.
But at the time, it was too cumbersome to get the design clean enough... 
Would be nice to revisit it and clean it up.

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

## lib

### grids
Not easy-to-use just yet...
`NewGrid()` constructor was naive and did not grow out of necessity.
`GridFromFile()` is "better" but restricts to `string` type...
`[][]byte` or `grids.ByteGrid` is probably better, maybe some generic T is possible

It needs NESW (or U, D, L, R) `directions` next to `allDirections`

`Abs()` probably needs to go in a different module (`ints`, `intmath`, or something along those lines.)

