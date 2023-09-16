#!/usr/bin/env bash

main() {
    echo "One for ${1:-you}, one for me."
}

main "$@"

# this also works
# main() {
#   name=${1}
#   printf "One for %s, one for me.\n" "${name:=you}"
# }
# main "$@"
