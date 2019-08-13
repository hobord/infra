http mock
---------

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

peeling module
--------------

- It is getting the session id and request url, and requestid
- get the the specific parameters save into the session


Redirection module
------------------

- getting url, requestID
- using redirection stack by requestId
- make logic
- response (redirect to url)