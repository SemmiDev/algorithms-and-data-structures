def binnarySearch(list, item):
	low = 0
	high = len(list)-1

	while low <= high:
		mid = (low + high) // 2
		guess = list[mid]
		if guess == item:
			return mid
		if guess > item:
			high = mid - 1
		else:
			low = mid + 1
	return None


def bubbleSort(list):
	for i in range(len(list)):
		for j in range(len(list)-1):
			if list[j] > list[j+1]:
				temp = list[j]
				list[j] = list[j+1];
				list[j+1] = temp
	return list;


# my_list = [1,3,5,7,9,13]
# angka = 9
# print(f"angka {angka} berada pada index {binnarySearch(my_list, angka)}")

result = bubbleSort([2,7,6,3,8,12,10])
print(result)