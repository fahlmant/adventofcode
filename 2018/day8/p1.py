with open('input') as f:
    items = [int(x) for  x in f.read().split()]

def parseTree(data):
    childrenNodes = data[0]
    metaNodes = data[1]
    otherNodes = data[2:]
    results = 0

    for _ in range(0, childrenNodes):
        result, otherNodes = parseTree(otherNodes)
        results += result
    
    results += sum(otherNodes[:metaNodes])
    return (results, otherNodes[metaNodes:])

result, left = parseTree(items)
print(result)        
