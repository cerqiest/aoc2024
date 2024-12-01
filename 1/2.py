list1 = []
list2 = []
similarity = 0

with open("1/input.txt") as f:
    data = f.read().splitlines()

    for line in data:
        num1 = line.split("   ")[0]
        num2 = line.split("   ")[1]
        list1.append(num1)
        list2.append(num2)

        print(num1, num2)

for i in range(len(list1)):
    similarity += list2.count(list1[i]) * int(list1[i])

print(similarity)
