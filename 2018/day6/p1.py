from collections import defaultdict
with open('input.1') as f:
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

coords = {coord_id: point for coord_id, point in enumerate(points, start=1)}

for i in range(0,xMax+1):
    for j in range(0, yMax+1):
        min_dists = sorted([(mannDist((x,y),(i,j)), coord_id) for coord_id, (x, y) in coords.items()])
        
        if( len(min_dists) == 1 or min_dists[0][0] != min_dists[1][0]):
            coord = min_dists[0][1]
            regionSizes[coord] += 1
        
            if i == 0 or i == xMax or j == 0 or j == yMax:
                infiniteAreaPoints.add(coord)

maxArea = 0
for item in regionSizes:
    if(item not in infiniteAreaPoints):
        maxArea = max(maxArea, regionSizes[item])

print(maxArea)