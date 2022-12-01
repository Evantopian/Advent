with open("ans.txt") as file:
    ans = file.read().strip().split("\n\n")

k1 = 0
k2 = 0
for a in ans:
    temp = set(a.replace("\n", ""))
    k1 += len(temp)

    groups = list(a.split("\n"))

    x = set(groups[0])
    for i in range(len(groups) - 1):
        x = set(x) & set(groups[i + 1])

    k2 += len(x)
    groups = []

print("Part 1: %s" % k1, "\nPart 2: %s" % k2)
