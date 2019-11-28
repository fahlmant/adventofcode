line = open('input').read().splitlines()[0]

def polarityCheck(x,y):
    return((x.isupper() and y.islower()) or (x.islower() and y.isupper()))

def sameChar(x,y):
    return (x.lower() == y.lower())

def reduce(poly):
    numchanged = 0
    finalline = []
    for letter in poly:
        if(len(finalline) == 0):
            finalline.append(letter)
        else:
            if(sameChar(letter, finalline[-1]) and (polarityCheck(letter, finalline[-1]))):
                finalline.pop()
                numchanged += 1
            else:
                finalline.append(letter)
    return (finalline, numchanged)

numchanged = 1
result = line
while(numchanged != 0):
    (result, numchanged) = reduce(result)


print(len(result))
