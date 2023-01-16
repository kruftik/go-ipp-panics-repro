# go-ipp / http.Response.Body strange reads example

* all `go-ipp` prepare request logic is omitted to simplify the example
* `cupsRequestBody` is a `wireshark`-dumped http-body of `GetPrinters` request to CUPS.
* `printers.conf` contains **1446** printers defined.

## TestMainDirect test

direct http.Response.Body parsing, without buffering

```bash

# run the test 100 times, multiple (or even all) fails are expected
go test -run TestMainDirect -count 100
```

## TestMainPreread test

read entire http.Response.Body into a byte slice, parse bytes.NewReader of that slice.

```bash

# run the test 100 times, none of the tries are expected to fail
go test -run TestMainDirect -count 100
```
