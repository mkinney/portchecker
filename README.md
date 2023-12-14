portchecker - check if hosts resolve and check it the port is open

-timeout: number of seconds to try

-check:
hosts and ports to check

examples:
ip 9100
host1,host2 80,1433

returns 0 if all were successful, otherwise returns the number of failures (number of hosts not resolving plus the number of ports not open) or 125

Development:
- "go mod init" to initialize (creates go.mod)
- "go build"
- "./portchecker"

Inspired by:
- https://github.com/dddpaul/gonc
- https://stackoverflow.com/a/59621864
