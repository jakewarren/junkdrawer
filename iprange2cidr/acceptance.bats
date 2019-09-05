@test "copy & paste range from whois results" {
    run go run main.go 19.136.100.96 - 19.136.100.127
    [ "${lines[0]}" == '19.136.100.96/27' ]
}

@test "specify two IPs with no delimeter" {
    run go run main.go 172.217.0.0 172.217.255.255
    [ "${lines[0]}" == '172.217.0.0/16' ]
}