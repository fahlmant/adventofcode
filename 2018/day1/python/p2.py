import itertools
with open('input') as f:
        lines = f.readlines()

numsum = 0
numseen = {0}

for num in itertools.cycle(lines):
    numsum += int(num)
    if numsum in numseen:
        print(numsum)
        print(len(numseen))
        break
    numseen.add(numsum)

