END {
    printf "One for " ($0 ? $0 : "you") ", one for me." > "/dev/stderr"
}
