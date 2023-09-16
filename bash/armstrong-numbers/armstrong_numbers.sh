#!/usr/bin/env bash

main() {
    declare -i squaresum=0

    while read -r -n 1 digit; do
        squaresum+=$((digit ** ${#1}))
    done <<<"$1"

    # -eq checks numeric equality while == checks for string equality
    [[ $1 -eq squaresum ]] && echo "true" || echo "false"
}

main "$@"
