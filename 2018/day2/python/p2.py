import collections

#Read lines in
with open('input') as f:
    lines = f.readlines()

for i in lines:
    for j in lines:
        diff = [x for x,y in zip(i,j) if x==y]
        if len(i)-len(diff)==1:
            print("".join(diff))
            break
