with open("input.txt") as f:
    data = f.read().strip().split()

fullyContains = 0
rangesOverlaps = 0

for line in data:
    elves = line.split(",")
    ranges = [list(map(int, elf.split("-"))) for elf in elves]
    print(ranges)

    elf1Start, elf1End = ranges[0]
    elf2Start, elf2End = ranges[1]

    cond1 = (elf1Start <= elf2Start) and (elf1End >= elf2End)
    cond2 = (elf1Start >= elf2Start) and (elf1End <= elf2End)

    if cond1 or cond2:
        fullyContains += 1
    
    cond3 = elf1Start > elf2End
    cond4 = elf1End < elf2Start

    if not (cond3 or cond4):
        rangesOverlaps += 1

print("Fully contains:", fullyContains)
print("Ranges overlaps:", rangesOverlaps)