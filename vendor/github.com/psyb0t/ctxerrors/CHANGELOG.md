# Changelog

All notable changes per release. Versions follow [semver](https://semver.org)
pre-1.0 conventions: minor bumps may include breaking API changes (called out
explicitly), patch bumps are docs / build / fixes only.

## v0.3.2 — 2026-07-27

Self-hosted README badges.

- **Coverage / version / license badges are self-rendered SVGs** served from
  `raw.githubusercontent.com/psyb0t/ctxerrors/badges/*.svg` — no third-party
  render service. `make test-coverage` writes the percentage to
  `coverage-percent.txt`, the pipeline uploads it, and a `badges` job bakes it
  into the SVG. The CI badge is switched to GitHub's native `badge.svg`. No
  library code changed.

## v0.3.1 — 2026-07-26

README badges.

- pkg.go.dev reference + GitHub Actions CI status badges. No library code
  changed.

## v0.3.0 — 2026-05-11

- Added an error mapping layer.

## v0.2.3 — 2026-01-19

- Removed the external logrus dependency; use stdlib `log/slog` for nil-error
  warnings. No API changes.

## v0.2.2 — 2026-01-16

- Updated the Go version in the CI pipeline.

## v0.2.1 — 2026-01-16

- Modernized tooling and updated dependencies.

## v0.2.0 — 2025-09-10

- Renamed the error struct to `CTXError`.

## v0.1.0 — 2025-09-07

- Initial release.
