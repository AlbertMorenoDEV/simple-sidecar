# [Simple Sidecar](https://github.com/AlbertMorenoDEV/simple-sidecar)

[![codecov](https://codecov.io/gh/AlbertMorenoDEV/simple-sidecar/branch/master/graph/badge.svg)](https://codecov.io/gh/AlbertMorenoDEV/simple-sidecar)


# Start feature flag server

> go run main.go featureflagapi --with-data


# Call examples

> curl http://localhost:7983/health

> curl --header "X-Session-Token: 00000000" http://localhost:7983/parameters

> curl --header "X-Session-Token: 00000000" http://localhost:7983/parameters/parameter_1

> curl --header "X-Session-Token: 00000000" -d '{"ID":"parameter_1", "Value":"new value"}' -H "Content-Type: application/json" -X POST http://localhost:7983/parameters/test_1

> curl --header "X-Session-Token: 00000000" -d '{"ID":"test_1", "Value":"1111"}' -H "Content-Type: application/json" -X POST http://localhost:7983/parameters/test_1

> curl --header "X-Session-Token: 00000000" -X DELETE http://localhost:7983/parameters/test_1

# Run tests

> go test ./... -v --bench . --benchmem --race