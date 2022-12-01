def singleSlope(arr, x, y):
    yPos = 0
    xPos = 0

    trees = 0
    while yPos < len(arr):
        row = arr[yPos]

        if row[xPos % len(row)] == '#':
            trees += 1
        yPos += y
        xPos += x
    return trees


with open("map.txt") as file:
    mapArray = file.read().splitlines()

print(singleSlope(mapArray.copy(), 3, 1))

slopes = [(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)]
prod = 1

for s in slopes:
    prod *= singleSlope(mapArray.copy(), s[1], s[0])
print(prod)
