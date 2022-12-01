#  Areas nay seem redundant, but I coded to exhibit the different parts.

def countValid(arr):
    count = 0

    for i in range(len(arr)):
        temp = arr[i].split(":")
        min = int(temp[0][0:temp[0].index("-")])
        char = temp[0][-1]
        max = int(temp[0][temp[0].index("-") + 1:(temp[0].index(char)) - 1])
        pw = temp[1][1:]

        if min <= pw.count(char) <= max:
            count += 1
    return count


def countToboggan(arr):
    count = 0

    for i in range(len(arr)):
        temp = arr[i].split(":")
        min = int(temp[0][0:temp[0].index("-")])-1
        char = temp[0][-1]
        max = int(temp[0][temp[0].index("-") + 1:(temp[0].index(char)) - 1])-1
        pw = temp[1][1:]

        if not(pw[min] == char and pw[max] == char) and (pw[min] == char or pw[max] == char):
            count+= 1
    return count


with open('data.txt') as file:
    testArray = file.read().splitlines()

print(countValid(testArray.copy()))
print(countToboggan(testArray.copy()))
