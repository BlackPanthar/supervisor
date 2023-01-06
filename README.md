# supervisor
A simple supervisor to monitor node SLA

## features

The supervisor current supports the following features:

1. Check evm and cosmos node( sync check , 50% integrity check), update the performance, calculate the SLA in performance.json,the log will be saved in nodemonitor.log
2. The integrity check is configured in integrity.toml. You can disable some check or change the default address.
2. The granularity of the checks is configurable in config.toml. The supervisor will do an availability check at least once in the configured interval.
3. The schedule of the checks is randomised.







## How to use

1. fill the config.toml like the example

2. fill the integritycheck.toml like the example

3. `go build main.go`
4. `./main`

## How to test the test check case

```
cd src/check
go test
```

