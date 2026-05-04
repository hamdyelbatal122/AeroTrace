# AeroTrace

[![CI](https://github.com/hamdyelbatal122/AeroTrace/actions/workflows/ci.yml/badge.svg)](https://github.com/hamdyelbatal122/AeroTrace/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/hamdyelbatal122/AeroTrace)](https://github.com/hamdyelbatal122/AeroTrace/blob/main/go.mod)
[![License](https://img.shields.io/github/license/hamdyelbatal122/AeroTrace)](LICENSE)

AeroTrace is a lightweight Go CLI that evaluates Skylark/Starlark-like `.sky` files and emits Kubernetes manifests as YAML.

## Features

- Generate Kubernetes objects from reusable `.sky` functions.
- Built-in object output helper (`output_type`) for consistent YAML generation.
- Supports `load(...)` across `.sky` modules.
- Optional file output via CLI (`-o`) for CI/CD pipelines.

## Requirements

- Go (version defined in `go.mod`)

## Quick Start

```bash
git clone https://github.com/hamdyelbatal122/AeroTrace.git
cd AeroTrace
go mod tidy
go build -o aerotrace .
```

Generate YAML to stdout:

```bash
./aerotrace foo.sky
```

Generate YAML to file:

```bash
./aerotrace -o manifests.yaml foo.sky
```

## Development

Use the provided Make targets:

```bash
make build
make test
make vet
make check
```

## Project Structure

```text
.
├── .github/workflows/ci.yml   # GitHub Actions CI
├── k8s/v1/api.sky             # Kubernetes API helper functions
├── foo.sky                    # Example input manifest source
├── main.go                    # CLI and Skylark execution engine
├── main_test.go               # Unit tests
├── Makefile                   # Common local tasks
└── README.md
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## Security

See [SECURITY.md](.github/SECURITY.md).

## Code of Conduct

See [CODE_OF_CONDUCT.md](.github/CODE_OF_CONDUCT.md).

## License

Licensed under [LICENSE](LICENSE).
