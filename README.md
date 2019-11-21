# Start feature flag server

> go run main.go featureflagapi --with-data


# Call examples

> curl --header "X-Session-Token: 00000000" http://localhost:7983/parameters
> curl --header "X-Session-Token: 00000000" http://localhost:7983/parameters/parameter_1
> curl --header "X-Session-Token: 00000000" -d '{"ID":"test_1", "Value":"1111"}' -H "Content-Type: application/json" -X POST http://localhost:7983/parameters/test_1
> curl --header "X-Session-Token: 00000000" -X DELETE http://localhost:7983/parameters/test_1