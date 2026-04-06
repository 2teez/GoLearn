#!/usr/bin/env python3

def main():
    """
    >>> arr = []
    >>> print(arr)
    []
    >>> arr.append([1,5,6])
    >>> print(arr)
    [[1, 5, 6]]
    >>> arr == [[1, 5, 6]]
    True
    >>> n_arr = arr[0][:2]
    >>> print(n_arr)
    [1, 5]
    """
    pass

if __name__ == '__main__':
    import doctest
    doctest.testmod(verbose=True)
