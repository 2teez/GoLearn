#!/usr/bin/env python

def add_numbers(a: int, b: int) ->int:
    """ Adding two numbers
    >>> add_numbers(4, 9)
    13
    >>> a, b = (7, 2)
    >>> add_numbers(a, b)
    >>> 9
    """
    return a + b

if __name__ == '__main__':
    import doctest
    doctest.testmod(verbose=True)
