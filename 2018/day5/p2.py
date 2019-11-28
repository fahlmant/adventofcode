from collections import defaultdict
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

def remove_values_from_list(the_list, val):
    return [value for value in the_list if value != val]

alphabet = 'abcdefghijklmnopqrstuvwxyz'
best = len(line)
bestletter = None

for letter in alphabet:
    polymer = line
    done = 1
    polymer = remove_values_from_list(polymer, letter)
    polymer = remove_values_from_list(polymer, letter.upper())
    while(done != 0):
        (polymer, done) = reduce(polymer)
    if(len(polymer) < best):
        best = len(polymer)
        bestletter = letter

print(best)
