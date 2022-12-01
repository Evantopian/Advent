with open("id.txt") as file:
    idArr = file.read().splitlines()

for i in idArr:
    idArr[idArr.index(i)] = int(i.replace("B", "1").replace("R", "1").replace("F", "0").replace("L", "0"), 2)
idArr.sort()

print(idArr)

# part 1
print(max(idArr))

# part 2
idRange = set(range(min(idArr), max(idArr)))
print(idRange - set(idArr))
