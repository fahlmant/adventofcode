with open('input') as f:
    lines = f.readlines()

maxX = 100
maxY = 100

points = []
for line in lines:
    posX = int(line[10:15])
    posY = int(line[17:23])
    velX = int(line[36:38])
    velY = int(line[40:42])
    points.append([posX, posY, velX, velY])

def printGrid(grid, X, Y):
    for i in range(-X, X):
        for j in range(-Y, Y):
            print(grid[i+int(maxX/2)][j+int(maxY/2)], end='')
        print('\n')

def move(grid, points):
    pass
    for point in points:
        if point[0][0] < (maxX/2) and point[0][0] > -maxX/2 and point[0][1] < maxY/2 and point[0][1] > -maxY/2:
            grid[point[0][0] + int(maxX/2)][point[0][1]+int(maxY/2)] = '.'
            tmp = list(point[0])
            tmp[0] = point[0][0] + point[1][0]
            tmp[1] = point[0][1] + point[1][1]
            point = (tuple(tmp),point[1])
            grid[point[0][0] + int(maxX/2)][point[0][1]+ int(maxY/2)] = '*'

grid = [['.' for i in range(maxX)] for j in range(maxY)]

for point in points:
    if point[0] < int(maxX/2) and point[0] > (int(maxX/2)*-1) and point[1] < int(maxY/2) and point[1] > (int(maxY/2)*-1):
        grid[point[0]+int(maxX/2)][point[1]+ int(maxY/2)] = '*'

for _ in range(4):
    printGrid(grid, 50,50)
    move(grid, points)
    #print('iter')
