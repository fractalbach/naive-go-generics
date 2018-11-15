import sys


if len(sys.argv) != 3:
    sys.exit("Need 2 argument: \n\t 1: Before \n\t 2: AFTER")

oldname = sys.argv[1]
newname = sys.argv[2]

for line in sys.stdin:
    print(line.replace(oldname, newname), end='')
