import sys
import os
from datetime import datetime, timezone

TEMPLATE = """// generated at {now}, for day {day} part {part}
package main

import "advent-of-code-2021/utils"

func main() {{
    scanner, err := utils.LoadScanner("{data}")
    if err != nil {{ panic(err) }}

    _ = scanner
}}

"""

assert len(sys.argv) == 2, "newday must be run with only one argument"
day = int(sys.argv[1])
part1_path = f"day{day}/part1"
part2_path = f"day{day}/part2"

# create project folders
try:
    os.makedirs(part1_path)
    os.makedirs(part2_path)
    print(f"day {day} directories done...")
except:    
    print("directories for this day probably exist already")

def init_basic(path, **kwargs):
    with open(path + "/main.go", "w") as f:
        f.write(TEMPLATE.format(**kwargs))

# generate program templates
try:
    now = datetime.now().astimezone(timezone.utc)
    init_basic(part1_path, now=now, day=day, part=1, data=f"data/day{day}.txt")
    init_basic(part2_path, now=now, day=day, part=2, data=f"data/day{day}.txt")
    print("generating templates...")
except Exception as e:
    print("can't generate templates")
    print(e)


# day txt file
open(f"data/day{day}.txt", "w").close()
print("making empty data file...")

