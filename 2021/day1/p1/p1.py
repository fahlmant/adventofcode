with open("../input") as f:
    lines = f.readlines()

lower = 0
higher = 0
answer = 0
for l in lines:
    print(l)
    higher = int(l)

    if lower < higher:
        answer += 1

    lower = int(higher)

print(answer - 1)