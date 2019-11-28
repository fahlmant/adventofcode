import collections

#Read lines in
with open('../input') as f:
    lines = f.readlines()

#Init vars
numTwos = 0
numThrees = 0
i = 0
results = [None]*len(lines)

#Read each line and count each character into a Counter object
def getResults(line):
    result = collections.Counter(line)
    return result

#For each value in key-value pair, check if there is a 2 or 3
for line in lines:
    results[i] = getResults(line)
    if 2 in (results[i].values()):
        numTwos = numTwos + 1
    if 3 in (results[i].values()):
        numThrees = numThrees + 1
    i = i+1

print(numTwos)
print(numThrees)
print(numTwos*numThrees)
