import unittest
from ch08.mymaths.mymaths import add_numbers

class Test_MathsClassMode(unittest.TestCase):
    """ The unittest covers the test for mymaths """

    def test_add_numbers(self):
        """ test the method that adds two numbers """

        expected = add_numbers(7, 2)
        got = 9
        self.assertEqual(expected, got)

if __name__ == '__main__':
    """ lauches main function """
    unittest.main()
