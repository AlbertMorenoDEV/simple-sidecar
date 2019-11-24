# [Simple Sidecar](https://github.com/AlbertMorenoDEV/simple-sidecar)

[![codecov](https://codecov.io/gh/AlbertMorenoDEV/simple-sidecar/branch/master/graph/badge.svg)](https://codecov.io/gh/AlbertMorenoDEV/simple-sidecar)
[![build](https://github.com/AlbertMorenoDEV/simple-sidecar/workflows/Build%20and%20Test/badge.svg)](https://github.com/AlbertMorenoDEV/simple-sidecar/actions?query=workflow%3A%22Build+and+Test%22)
[![release](https://img.shields.io/github/v/release/AlbertMorenoDEV/simple-sidecar.svg)](https://github.com/AlbertMorenoDEV/simple-sidecar/releases/latest)


# Build

> make build

# Run

> ./simple-sidecar

To configure the service you can use either environment variables or a .env file.

# Stop

Just press Ctrl+C and a graceful shutdown with 15 seconds (default) will be executed.

# Configuration

- SS_PARAMETERS_PORT (default 7983)
- SS_PARAMETERS_WRITE_TIMEOUT (default 15s)
- SS_PARAMETERS_READ_TIMEOUT (default 15s)
- SS_PARAMETERS_IDLE_TIMEOUT (default 60s)
- SS_PARAMETERS_GRACEFUL_TIMEOUT (default 15s): The duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 60s
- SS_DEBUG_MODE (default false)
- SS_AUTH_TOKENS (default empty): Set accepted auth tokens separeted by coma. Clients must send tokens through HTTP Header X-Session-Token.

# Run tests

> make test

# Create new version

> git tag v0.1.0-alpha

> git push --tags

After that release workflow will create a new version using goreleaser.

# Parameters call examples

> curl http://localhost:7983/health

> curl --header "X-Session-Token: 00000000" http://localhost:7983/parameters

> curl --header "X-Session-Token: 00000000" http://localhost:7983/parameters/parameter_1

> curl --header "X-Session-Token: 00000000" -d '{"ID":"parameter_1", "Value":"new value"}' -H "Content-Type: application/json" -X POST http://localhost:7983/parameters/test_1

> curl --header "X-Session-Token: 00000000" -d '{"ID":"test_1", "Value":"1111"}' -H "Content-Type: application/json" -X POST http://localhost:7983/parameters/test_1

> curl --header "X-Session-Token: 00000000" -X DELETE http://localhost:7983/parameters/test_1