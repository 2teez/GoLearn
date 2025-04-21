#!/usr/bin/env bash
# Description: Make a script that lunches golang program
# Date: 21/04/25
# Author: omitida

function help() {
    echo "
    ./makego [options] <filename>

    options
    =======
    -b  build and run a go file. This generate an executable file.
    -d  delete a go file.
    -g  generate and run a generic go file.
    -h  print the help function of the script.
    "
}

# write out a basic go file
function generate_go_file() {
    filename="${1}"

    if [[ -e "${filename}" ]]; then
        echo "The file ${filename} exists, and can't be overwritten."
        exit 1
    fi

    # basic golang line to write to file
    echo "
        package main

        import (
            \"fmt\"
        )

        func main() {
            fmt.Println(\"Start Here...\")
        }
        " > "${filename}"
}

# check if the file has an extension check
# if it doesn't add one: Which should be go
function file_extension_check() {
    filename=${1}
    expected_extension="${2}"
    ext="${filename##*.}"
    if [[ "${ext}" != "${expected_extension}" ]]; then
        filename="${filename}.go"
    fi
}

function format_code() {
    filename=${1}
    gofmt "${filename}" > "${filename}_tmp"
    cat "${filename}_tmp" > "${filename}"
    rm "${filename}_tmp"
}

# check and start the script here
[[ "${#}" -ne 2 ]] && help

# options string to use
optstring=":b:d:g:h"

while getopts "${optstring}" opt; do
    case "${opt}" in
        b)
            filename="${OPTARG}"

            # call the function with parameters
            file_extension_check "${filename,,}" "go"
            generate_go_file "${filename}"
            format_code "${filename}"

            go mod init $(basename "${PWD}")
            go mod tidy

            go build .
            go run .
        ;;
        d)
            filename="${OPTARG}"
            filename=$(ls | grep -i "${filename}")

            while read -p "Do you want to delete the ${filename}? [yn]: " -r ans; do
                case "${ans,,}" in
                    y)
                        rm -rf "${filename}"
                        echo "${filename} deleted."
                        exit 0
                        ;;
                    n)
                        echo "${filename} NOT deleted."
                        exit 0
                        ;;
                    *)
                    echo "Use either y or n."
                    continue
                    ;;
                esac
            done
        ;;
        g)
            filename="${OPTARG}"

            # call the function with parameters
            file_extension_check "${filename,,}" "go"
            generate_go_file "${filename}"
            format_code "${filename}"

            go run "${filename}"
        ;;
        h)
            help
            exit 1
        ;;
        *)
            echo "invalid option specified."
            exit 1
        ;;
    esac
done
