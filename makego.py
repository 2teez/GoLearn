#!/usr/bin/env python3

import subprocess

class HyphenError(Exception):
    def __init__(self, msg, /):
        self.__msg = msg

    def __str__(self, /) -> str:
        return f'{self.__msg}'


def check_option(opt) -> str:
    if '-' not in opt:
        raise HyphenError('Options values must have an hyphen')
    return opt

def check_file_exist(file: str) -> bool:
    """ check if the file exists in the directory """
    import os
    return file in os.listdir('.')


def run_file(*args):
    # search  the directory to findout if it is a package

    subprocess.run(['gofmt', '-w', '.'])
    if check_file_exist('go.mod'):
        subprocess.run(['go', 'run', '.'])
    else:
        filename = args[-1]
        subprocess.run(['go', 'run', filename])

def create_generic_go_file(*args):
    """ create a go file with main func """

    filename = args[-1]

    if not filename.endswith('go'):
        filename = f'{filename}.go'

    if check_file_exist(filename):
        raise Exception(f"Can't overwrite the {filename}")

    with open(filename, 'w') as f:
        f.writelines(
    """
    package main
    import ("fmt")

    func main() { fmt.Println("Start Here!")}
        """)

    subprocess.run(['gofmt', '-w', '.'])

def user_input(msg: str = 'Enter: ') -> str:
    """ get user answer """

    answer = input(msg).strip().lower()
    while answer not in ['n', 'y']:
        print(f'{answer} invalid')
        answer = input(msg).strip().lower()

    return answer

def delete_go_file(*args):
    """ delete specified file """

    filename = args[-1]
    status = check_file_exist(filename)
    if status:
        if user_input(f'Want to delete file > {filename} <: ') == 'y':
            subprocess.run(['rm', filename])
    else:
        print(f'{filename} doesn\'t exists.')

def process_cli(*args):
    """ Run all the args using options in a dictionary containing function """

    options = {'-r': run_file, '-g': create_generic_go_file, '-d': delete_go_file}
   # check the data for the second parameter must have -
    option = check_option(args[1])
    return options.get(option, helper)(*args[1:])

def helper(*cli_opts):
    """ Display the options avaliable on this script"""

    print("""
        ./makego.py <options> <filename>
        Options Avaliable:
            -r  to run a go file
            -g  to make a generic go file
            -d to delete a named go file
        """)
    sys.exit(1)

def main(*args):
    """ main function to lunch the file """
    if len(args) != 3:
        # call helper function
        helper(*args)
    process_cli(*args)

if __name__ == '__main__':
    import sys

    main(*sys.argv)
