import unittest

from makego import main, process_cli

class TestCliFunctions(unittest.TestCase):
    def test_main_cli_function(self):
        expected = main(*['py_script.py','-d', 'main.go'])
        got = None
        self.assertEqual(expected, got)

    def test_process_cli_function(self):
        expected = process_cli(*['py_script.py','-d', 'main.go'])
        got = 2
        self.assertEqual(got, expected)
