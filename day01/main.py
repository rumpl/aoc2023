import re
f = open("input.txt", "r")

def convert(s):
    d = {}

    nums = ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
    letters = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

    for n in nums:
        idxs = [m.start() for m in re.finditer(n, s)]
        for idx in idxs:
            d[idx] = int(n)

    for i, n in enumerate(letters):
        idxs = [m.start() for m in re.finditer(n, s)]
        for idx in idxs:
            d[idx] = i + 1

    return d[min(d.keys())] * 10 + d[max(d.keys())]

num = 0
for line in f:
    num += convert(line)

print(num)
