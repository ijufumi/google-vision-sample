# Backend

T.B.D

## How to test

### How to generage mock

First, install `mockery`.

```bash
go install github.com/vektra/mockery/v2@latest
```

Then, generate `mock`.

> Stub codes of Storage API

```bash
mockery --srcpkg=cloud.google.com/go/storage --output=internal/infrastructures/google/clients/stubs --outpkg=stubs --all --case=snake --testonly=true --disable-version-string=true
```