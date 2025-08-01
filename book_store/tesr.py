def merge_sort(arr):
    if len(arr) <= 1:
        return arr

    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])

    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0

    # Fusionne les deux tableaux triés
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1

    # Ajoute le reste
    result.extend(left[i:])
    result.extend(right[j:])
    return result

# Exemple :
print(merge_sort([5, 2, 9, 1, 3, 4]))