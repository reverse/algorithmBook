def SingleNumberHashMap(nums):
    store = {}

    for item in nums:
        if item in store:
            store[item] += 1
            continue
        
        store[item] = 1 

    
    for i in store:
        if store[i] == 1:
            return i