import math
import os

list1 = []
list2 = []
difference = 0

with open("1/input.txt") as f:
    data = f.read().splitlines()

    for line in data:
        num1 = line.split("   ")[0]
        num2 = line.split("   ")[1]
        list1.append(num1)
        list2.append(num2)

        print(num1, num2)

list1 = sorted(list1)
list2 = sorted(list2)

for i in range(len(list1)):
    difference += math.fabs(int(list2[i]) - int(list1[i]))

print(difference)
