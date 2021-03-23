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

def findSmalles(arr):
	smallest = arr[0]
	smallest_index = 0
	for i in range(1, len(arr)):
		if arr[i] < smallest:
			smallest = arr[i]
			smallest_index = i
	return smallest_index

def selectionSort(arr):
	newArr = []
	for i in range(len(arr)):
		smallest = findSmalles(arr)
		newArr.append(arr.pop(smallest))
	return newArr


# my_list = [1,3,5,7,9,13]
# angka = 9
# print(f"angka {angka} berada pada index {binnarySearch(my_list, angka)}")

#result = bubbleSort([2,7,6,3,8,12,10])
#print(result)

print(selectionSort([10,9,8,7,6,5,4,3,2,1]))