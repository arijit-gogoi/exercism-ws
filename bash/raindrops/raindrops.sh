#!/usr/bin/env bash

set -o errexit
set -o nounset

main() {
    [[ $# -ne 1 ]] && exit 1

    local num=$1
    local str=""

    ((num % 3 == 0)) && str+="Pling"
    ((num % 5 == 0)) && str+="Plang"
    ((num % 7 == 0)) && str+="Plong"

    # If the variable str is defined and has a value,
    # the code will print the value of str to the console.
    # If the variable str is not defined,
    # the code will print the value of num to the console.
    echo "${str:-$num}"
}

main "$@"
