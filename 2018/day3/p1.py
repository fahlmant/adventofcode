#Read lines in
with open('input') as f:
    lines = [line.rstrip('\n') for line in f]

#create fabric with each square inch = a 2d co-ord
fabric = [[0 for i in range(1200)] for j in range(1200)]
sum = 0

for line in lines:
    xcoord = int(line[line.find('@')+1:line.find(',')])
    ycoord = int(line[line.find(',')+1:line.find(':')])
    xlen = int(line[line.find(':')+2:line.find('x')])
    ylen = int(line[line.find('x')+1:])
    for x in range (xcoord, xcoord+xlen):
        for y in range(ycoord, ycoord+ylen):
            fabric[x][y] = fabric[x][y] + 1

for x in range (0, 1200):
    for y in range(0, 1200):
        if fabric[x][y] >= 2:
            sum = sum + 1

print(sum)
