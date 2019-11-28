from collections import defaultdict
import heapq
with open('input') as f:
    lines = [line.rstrip('\n') for line in f]


class Elves:
    def __init__(self):
        self.work = None
        self.idle = None
        self.time = 0

    def startWork(self, work):
        self.work = work
        self.idle = False
        self.time = (ord(work) - ord('A') + 1) + 60
    
    def busy(self):
        return (self.time > 0)
    
    def step(self):
        if self.busy:
            self.time -= 1
            if self.time == 0:
                self.idle = True

graph = []
succeed = defaultdict(list)
precede = defaultdict(list)

for line in lines:
    words = line.split(" ")
    graph.append((words[1], words[7]))
    succeed[words[1]].append(words[7])
    precede[words[7]].append(words[1])

elves = [Elves() for _ in range(0,5)]

order = []
queue = []
startNodes = set(succeed.keys()) - set(precede.keys())
for node in startNodes:
    heapq.heappush(queue, node)

finaltime = 0

while len(order) != 26 :
    available_workers = [elf for elf in elves if not elf.busy()]

    for worker in available_workers:
        if queue:
            next_task = heapq.heappop(queue)
            worker.startWork(next_task)

    for worker in elves:
        worker.step()
        if worker.idle:
            order.append(worker.work)

            for node in succeed[worker.work]:
                pred = precede[node]
                if node not in order and all(p in order for p in pred):
                    heapq.heappush(queue, node)

            worker.idle = False

    finaltime += 1
print(finaltime)
