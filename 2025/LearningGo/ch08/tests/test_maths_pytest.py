import pytest

from ch08.mymaths.mymaths import add_numbers

def test_add_number():
    """ adding two numbers """
    assert add_numbers(4, 8), 12
