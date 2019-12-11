def merge(arr, start, mid, end):
    left = arr[start:mid+1]
    right = arr[mid+1:end+1]
    i, j, k = 0, 0, start

    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            arr[k] = left[i]
            i += 1
        else:
            arr[k] = right[j]
            j += 1
        k += 1

    while i < len(left):
        arr[k] = left[i]
        i += 1
        k += 1

    while j < len(right):
        arr[k] = right[j]
        j += 1
        k += 1

    return arr

def mergesort(arr, l, r):
    """
    >>> mergesort([3,2,1],0,2)
    [1, 2, 3]
    >>> mergesort([3,2,1,0,1,2,3,5,4],0,8)
    [0, 1, 1, 2, 2, 3, 3, 4, 5]
    """

    if l < r:
        m = (l + r) // 2

        mergesort(arr, l, m)
        mergesort(arr, m+1, r)
        merge(arr, l, m, r)

        return arr

if __name__ == "__main__":
    arr = [3, 2, 1, 0, 1, 2, 3, 5, 4]
    print(arr)
    print(mergesort(arr, 0, len(arr)))
