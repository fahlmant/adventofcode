from collections import deque, defaultdict

with open('input') as f:
    data = f.read().split()
numPlayers = data[0]
lastMarble = data[6]

circle = deque([0])
players = defaultdict(int)
player = 1

for marble in range(1, (int(lastMarble)*100) +1):
    if marble % 23 == 0:
        circle.rotate(-7)
        players[player] += marble + circle.pop()
    else:
        circle.rotate(2)
        circle.append(marble)
    player += 1
    if player > int(numPlayers):
        player = 1

    
if players:
    print(max(players.values()))