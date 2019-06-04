def swap(arr, i, j):
	arr[i], arr[j] = arr[j], arr[i]
	return arr

def permutations():
	return swap([1,3], 0, 1)

print(permutations())