def TwoSumSolution(array, target):
    store = {}

    for idx, item in enumerate(array):
        compl = target - item
        if compl in store:
            return [store[compl], idx]
        
        store[item] = idx

    return [] 