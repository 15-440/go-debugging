# Go close/deadlock exercise

The root files are the buggy TA version. `REF/server.go` is the corrected
reference solution with optional logging.

## Before the fix

Run the test from the repository root:

```bash
go test .
```

The test should fail after 100ms because the server tries to send an ACK after
the client sends `Close`, but the client is no longer receiving responses.

## After the fix and logs

Copy the contents of `REF/server.go` into the root `server.go`. If copying the
whole file, remove its `//go:build ignore` line and the blank line after it so
Go includes the file.

Run the corrected test:

```bash
go test -race .
```

Logging is controlled by the CLI boolean `-log`. Logs are off by default:

```bash
go test -v .
```

Turn them on with:

```bash
go test -race -v . -args -log
```

You can explicitly turn them off with `-args -log=false`.
