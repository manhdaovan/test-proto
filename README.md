## Context
- Same `echo.proto` file, **with no `go_package` option**
- Generate `service1` and it client into `service/service1` directory
- Generate `service2` and it client into `service/service2` directory
- Start server with service from `service/service1`
- Call client1 from `service/service1`
- Call client2 from `service/service2`

## Expect
- Both client1 and client2 called successfully.

## Result
- As expect.
- Detail:
  - Server log:
    ```
	--- start serving --- 
	in ------ msg:"message from client 1" order:1 
	in ------ msg:"message from client 2" order:2 
    ```
  - client1 log:
    ```
	--- start request client1 --- 
	out ------ msg:"msg from server" order:1 
	```
  - client2 log:
    ```
	--- start request client2 --- 
	out ------ msg:"msg from server" order:2 
	```

## NOTE:
If **`go_package` option** was given in `echo.proto` file, `option go_package="github.com/manhdaovan/service/service1";` for example, then service2 cannot be generated correctly in other service.
=> so:
- the `option go_package` must be explicit and consistent if have. 
- the `option go_package`, if have, SHOUD NOT include multiple level.
  - For example `option go_package="github.com/manhdaovan/test_proto/service";`, SHOULD be change to `option go_package="service";`. All other level (`github.com/manhdaovan/test_proto/` in this case) SHOULD be defined when executing `protoc` command.
  - Instead of:
    ```
    // echo.proto
    option go_package="github.com/manhdaovan/test_proto/service";
    // protoc executing
    $protoc --go_out=plugins=grpc:${GOPATH}/src/ ./echo.proto;
	```
  - SHOULD be changed to:
    ```
    // echo.proto
    option go_package="service"; // or NO defined this option
    // protoc executing
    $protoc --go_out=plugins=grpc:${GOPATH}/src/github.com/manhdaovan/test_proto/service ./echo.proto;
	```