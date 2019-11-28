#Read lines in
with open('input') as f:
    lines = [line.rstrip('\n') for line in f]

#create fabric with each square inch = a 2d co-ord
fabric = [[0 for i in range(1200)] for j in range(1200)]
sum = 0
data = [None] * len(lines)
i=0

for line in lines:
    index = int(line[line.find('#')+1:line.find('@')-1])
    xcoord = int(line[line.find('@')+1:line.find(',')])
    ycoord = int(line[line.find(',')+1:line.find(':')])
    xlen = int(line[line.find(':')+2:line.find('x')])
    ylen = int(line[line.find('x')+1:])
    data[i] = (index, ((xcoord, ycoord), (xlen, ylen)))
    i = i + 1
    for x in range (xcoord, xcoord+xlen):
        for y in range(ycoord, ycoord+ylen):
            fabric[x][y] = fabric[x][y] + 1

def allOnes(xcoord, ycoord, xlen, ylen):
    claimSum = 0
    for x in range (xcoord, xcoord+xlen):
        for y in range (ycoord, ycoord+ylen):
            claimSum = claimSum + fabric[x][y]
    if claimSum == (xlen*ylen):
        return True
    else:
        return False
            
for item in data:
    if allOnes(item[1][0][0], item[1][0][1], item[1][1][0], item[1][1][1]):
        print(item)
