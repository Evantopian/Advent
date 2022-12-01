def validatePass(arr, con):
    vArray = []
    batch = {}

    for i in arr:
        if i:
            fields = i.split()

            for f in fields:
                key, val = f.split(':')
                batch[key] = val
        else:
            test = True

            for x in con:
                if not (x in batch):
                    test = False
                    break

            if test:
                vArray.append(batch)
            batch = {}
    return vArray


def rangeTest(v, min, max):
    return min <= int(v) <= max


def conditionsPass(arr, cTest):
    count = 0

    for test in arr:
        valid = True

        # years
        if not rangeTest(test[cTest[0]], 1920, 2002) or \
                not rangeTest(test[cTest[1]], 2010, 2020) or \
                not rangeTest(test[cTest[2]], 2020, 2030):
            valid = False

        # height
        temp = test[cTest[3]][:-2]
        if test[cTest[3]].endswith('in'):
            if not rangeTest(temp, 59, 76):
                valid = False

        elif test[cTest[3]].endswith('cm'):
            if not rangeTest(temp, 150, 193):
                valid = False

        else:
            valid = False

        temp = test[cTest[4]]

        # hair color
        if not test[cTest[4]][0] == '#':
            valid = False

        for x in temp[1:]:
            if x not in '0123456789abcdef':
                valid = False
                break

        # eye color
        if test[cTest[5]] not in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
            valid = False

        # passport id
        if len(test[cTest[6]]) != 9:
            valid = False

        for x in test[cTest[6]]:
            if x not in '0123456789':
                valid = False
            break

        if valid:
            count += 1
    return count


def main():
    with open("batch.txt") as file:
        batchArray = file.read().splitlines()

    cTest = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
    reqArr = validatePass(batchArray.copy(), cTest)

    print(reqArr, "\n" + str(len(reqArr) + 1))
    print(conditionsPass(reqArr.copy(), cTest) + 1)
    
    
main()
