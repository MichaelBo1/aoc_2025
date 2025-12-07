import sys
import os

if len(sys.argv) < 2:
    print("Error: No file provided. Usage: python script.py <filename>")
    sys.exit(1)

filepath = sys.argv[1]

with open(filepath, "r") as fhand:
    line = fhand.read() # input is expected as a single "," separated line of values.
    ranges = [x.split("-") for x in line.split(",")]

    for range in ranges:
        assert(len(range) == 2)
        lower, upper = int(range[0]), int(range[1])
        print(lower, upper)
        diff = upper - lower