# Contributing

## Development Setup

1. Install Go.
2. Clone this repository.
3. Run:

```bash
go mod tidy
make check
```

## Pull Request Checklist

- Keep changes focused and small.
- Add or update tests for behavior changes.
- Run `make check` before opening the PR.
- Update documentation when CLI or behavior changes.

## Commit Message Style

Use clear, imperative messages, for example:

- `fix: validate CLI arguments`
- `feat: add output file flag`
- `test: cover value conversion`
