from time import time


def contain(collection, target):
    """Determine whether list contains target"""
    return target in collection


def bs_contains(orderd, target):
    """
    Binary search for determine ordered list contains target
    :return position if target in collection
    """
    low = 0
    high = len(orderd) - 1

    while low <= high:
        mid = (low + high)//2
        if target == orderd[mid]:
            return mid
        elif target < orderd[mid]:
            high = mid - 1
        else:
            low = mid + 1
    return -(low + 1)


def insert_in_place_simple(ordered, target):
    """Insert target in collection"""
    for i in range(len(ordered)):
        if target < ordered[i]:
            ordered.insert(i, target)
            return
    ordered.append(target)


def insert_in_place_bs(ordered, target):
    """Insert target in collection"""
    idx = bs_contains(ordered, target)
    if idx < 0:
        ordered.insert(-(idx + 1), target)
        return
    ordered.insert(idx, target)


def check_performance(function):
    """Check performace of search method"""
    n = 1024
    while n < 999999:
        sorted_list = list(range(n))

        now = time()
        function(sorted_list, n/4)
        done = time()

        print('{0} {1}'.format(n, (done-now)*1000))

        n *= 2


if __name__ == '__main__':
    check_performance(insert_in_place_simple)
    print ('=====================================')
    check_performance(insert_in_place_bs)
    print ('=====================================')
    check_performance(contain)
    print ('=====================================')
    check_performance(bs_contains)