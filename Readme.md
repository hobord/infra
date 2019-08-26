## http mock

- It has some route
- All route can get some get query parameters
- Create a session if not exists
- Save cookie

ENV:

- PORT: 8100
- SESSION_GRPC_SERVER: 10.20.35.111:30645
  - SESSION_TTL: 0
  - SESSION_COOKIE_KEY: session
  - SESSION_REQUEST_GET_KEY: session
  - SESSION_REQUEST_POST_KEY: session
-

flow

- call the peeling module (send the request url, and session id, requestId) response ok
- call redirection module, if redirect then redirect and send session cookie
- check the reverse proxy config and get content by rules (send session id) and response with result and session cookie and request id as response id.

## peeling module

- It is getting the session id and request url, and requestid
- get the the specific parameters save into the session

## Redirection module

- getting url, requestID
- using redirection stack by requestId
- make logic
- response (redirect to url)

## Stress test

https://github.com/tsenart/vegeta

```
echo "GET http://127.0.0.1:8100/" | vegeta attack -duration=60s -workers=250 | tee results.bin | vegeta report
  vegeta report -type=json results.bin > metrics.json
  cat results.bin | vegeta plot > plot.html
  cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
```

```
echo 'GET http://localhost:8100' | \
    vegeta attack -rate 5000 -duration 10m | vegeta encode | \
    jaggr @count=rps \
          hist\[100,200,300,400,500\]:code \
          p25,p50,p95:latency \
          sum:bytes_in \
          sum:bytes_out | \
    jplot rps+code.hist.100+code.hist.200+code.hist.300+code.hist.400+code.hist.500 \
          latency.p95+latency.p50+latency.p25 \
          bytes_in.sum+bytes_out.sum
```
