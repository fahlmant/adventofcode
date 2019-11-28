from collections import defaultdict
with open('input') as f:
    lines = [line.rstrip('\n') for line in f]

points = set()
infiniteAreaPoints = set()
regionSizes = defaultdict(int)

xMax = 0
yMax = 0

for line in lines:
    x = int(line[:line.find(',')])
    xMax = max(xMax, x)
    y = int(line[line.find(',')+1:])
    yMax = max(yMax, y)
    points.add((x,y))

def mannDist(point1, point2):
    return (abs(point1[0] - point2[0]) + abs(point1[1] - point2[1]))

sharedSize = 0

for i in range(0,xMax+1):
    for j in range(0, yMax+1):
        isLess = sum(mannDist((x,y),(i,j)) for (x,y) in points) < 10000
        if isLess:
            sharedSize += 1
        

print(sharedSize)