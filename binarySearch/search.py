def Search(array, target):
    left, right = 0, len(array)-1

    while left <= right:
        midpoint = (left+right)//2
        num = array[midpoint]

        if num == target:
            return midpoint
        elif target < num:
            right = midpoint - 1 
        else:
            left = midpoint + 1

    return -1 