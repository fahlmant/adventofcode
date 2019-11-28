with open('input') as f:
    items = [int(x) for  x in f.read().split()]

def parseTree(data):
    childrenNodes = data[0]
    metaNodes = data[1]
    otherNodes = data[2:]
    results = 0
    totals = []

    for _ in range(0, childrenNodes):
        result, total, otherNodes = parseTree(otherNodes)
        results += result
        totals.append(total)
    
    results += sum(otherNodes[:metaNodes])
    if(childrenNodes == 0):
        return (results, sum(otherNodes[:metaNodes]), otherNodes[metaNodes:])
    else:
        return(results, sum(totals[i-1] for i in otherNodes[:metaNodes] if i > 0 and i <= len(totals)), otherNodes[metaNodes:])

result, total, left = parseTree(items)
print(total)        
