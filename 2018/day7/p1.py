from collections import defaultdict
import heapq
with open('input') as f:
    lines = [line.rstrip('\n') for line in f]

graph = []
succeed = defaultdict(list)
precede = defaultdict(list)


for line in lines:
    words = line.split(" ")
    graph.append((words[1], words[7]))
    succeed[words[1]].append(words[7])
    precede[words[7]].append(words[1])

queue = []
startNodes = set(succeed.keys()) - set(precede.keys())
for node in startNodes:
    heapq.heappush(queue, node)

def allInList(nodes, list):
    return(all(item in list for item in nodes))

finalOrder = []
while(len(queue) != 0):
    item = heapq.heappop(queue)
    finalOrder.append(item)
    for node in succeed[item]:
        pre = precede[node]
        if ((node not in finalOrder) and allInList(pre, finalOrder)):
            heapq.heappush(queue, node)
