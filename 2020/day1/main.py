def product(arr):
    prod = 1
    for i in range(len(arr)):
        prod *= arr[i]
    return prod

def reportPair(arr, k):
    result = []

    while arr:
        num = arr.pop()
        diff = k - num
        if diff in arr:
            result.extend([num, diff])
            return result
    return []


def reportTriplet(arr, k):
    result = []

    while arr:
        num = arr.pop()
        diff = k - num
        test = reportPair(arr.copy(), diff)

        if len(test) == 2:
            result.extend([num, test[0], test[1]])
            return result
    return []



with open('nums.txt') as file:
    testArray = file.read().splitlines()

testArray = list(map(int, testArray))


print(testArray)

ptOne = reportPair(testArray.copy(), 2020)
ptTwo = reportTriplet(testArray.copy(), 2020)

print(ptOne)
print(product(ptOne))


print(ptTwo)
print(product(ptTwo))
