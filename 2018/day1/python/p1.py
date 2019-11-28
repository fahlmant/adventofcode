with open('input') as f:
    lines = f.readlines()

sum = 0

for line in lines:
    sum += int(line)

print(sum)
