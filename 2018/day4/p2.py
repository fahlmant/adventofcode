from collections import defaultdict
#Read lines in
with open('input') as f:
    lines = [line.rstrip('\n') for line in f]

#Chronological sort
lines.sort()

def getTimestamp(line):
    date = line[1:11]
    time = line[12:17]
    return int(time.split(':')[1])

guardsWithTime = defaultdict(int)
guardsTotal = defaultdict(int)
guardNum = None
sleepTime = None
for line in lines:
    time = getTimestamp(line)
    if 'begins shift' in line:
        guardNum = int(line.split()[3][1:])
        asleep = None
    elif 'falls asleep' in line:
        sleepTime = time
    elif 'wakes up' in line:
        for t in range(sleepTime, time):
            guardsWithTime[(guardNum, t)] += 1
            guardsTotal[guardNum] += 1

def bestTime(guard, guardsAndTime):
    bestMinute = 0
    for i in range(0,59):
        if(guardsAndTime[(guard, i)] > guardsAndTime[(guard, bestMinute)]):
            bestMinute = i
    print(bestMinute)
    return bestMinute
            
mostSlept = (max(guardsWithTime.items(), key=lambda a: a[1]))
sleepiestGuard = mostSlept[0][0]
sleepiestMinute = mostSlept[0][1]

print(sleepiestGuard*sleepiestMinute)