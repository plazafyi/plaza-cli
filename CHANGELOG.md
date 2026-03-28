# Changelog

## 0.2.0 (2026-03-28)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/plazafyi/plaza-cli/compare/v0.1.0...v0.2.0)

### Features

* **api:** api update ([6263b6f](https://github.com/plazafyi/plaza-cli/commit/6263b6ff0c88377138b68e6a16e7cbfbb000d74a))
* **api:** api update ([1d4dbf4](https://github.com/plazafyi/plaza-cli/commit/1d4dbf4e38e101a9e787cb487b9b5396bf60dc16))
* **api:** api update ([11c4fa4](https://github.com/plazafyi/plaza-cli/commit/11c4fa408b62a53c7f1a3b5645944c9800431eb6))
* **api:** api update ([e0cfeb7](https://github.com/plazafyi/plaza-cli/commit/e0cfeb7d1f61403062ddb134f3c87ed59e23b758))
* set CLI flag constant values automatically where `x-stainless-const` is set ([9dfbdaf](https://github.com/plazafyi/plaza-cli/commit/9dfbdaf00c7059af0d9dc46278dda13e58d2ddf8))


### Bug Fixes

* cli no longer hangs when stdin is attached to a pipe with empty input ([a0a80ea](https://github.com/plazafyi/plaza-cli/commit/a0a80ea0f7e9efa365ee53d46ce1a5139fc3786d))
* fix for off-by-one error in pagination logic ([a087dae](https://github.com/plazafyi/plaza-cli/commit/a087dae75e0dbb1654e8d8d7710e11596e1e4684))


### Chores

* **ci:** skip lint on metadata-only changes ([b83eaec](https://github.com/plazafyi/plaza-cli/commit/b83eaec313d6a95b12e807d3e9faf1dc485e7659))
* **internal:** update gitignore ([bfc57a1](https://github.com/plazafyi/plaza-cli/commit/bfc57a1456f8deceb216fb82c06acc664fbbe556))
* **internal:** update multipart form array serialization ([5cef7e1](https://github.com/plazafyi/plaza-cli/commit/5cef7e13df9ad05099815863114914a34a26f615))
* omit full usage information when missing required CLI parameters ([097fe19](https://github.com/plazafyi/plaza-cli/commit/097fe197e14d07f18be2da436c172cd9d1d9f8cd))
* **tests:** bump steady to v0.19.4 ([35b14fc](https://github.com/plazafyi/plaza-cli/commit/35b14fc1627bddcf05e8652158aeb1d1ec05f9a8))
* **tests:** bump steady to v0.19.5 ([7eab4a7](https://github.com/plazafyi/plaza-cli/commit/7eab4a7562c76a9deabb8cdd25caf6f33c1bb402))
* **tests:** bump steady to v0.19.6 ([508f853](https://github.com/plazafyi/plaza-cli/commit/508f853509084fabc9e7b30cacaf83ac49a715d2))
* **tests:** bump steady to v0.19.7 ([df70688](https://github.com/plazafyi/plaza-cli/commit/df706884a6f089fa774400993297357c413da159))

## 0.1.0 (2026-03-20)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/plazafyi/plaza-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** api update ([fb054c7](https://github.com/plazafyi/plaza-cli/commit/fb054c79c4484e6736f0dcf9eeca13cd5a78d069))
* initial SDK generation ([745c96a](https://github.com/plazafyi/plaza-cli/commit/745c96ab5533cb2a7105970688578d558d79c125))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([caea46d](https://github.com/plazafyi/plaza-cli/commit/caea46db06fbbbbb60ff8bb0bddf26bbf7b8bc47))
* better support passing client args in any position ([4b5637e](https://github.com/plazafyi/plaza-cli/commit/4b5637eac19a57fdd1d6039ee927a55d09a8c584))
* improve linking behavior when developing on a branch not in the Go SDK ([ea6d279](https://github.com/plazafyi/plaza-cli/commit/ea6d279f6ab123e05fc011e7e1395e5e78590116))
* improved workflow for developing on branches ([c699ac9](https://github.com/plazafyi/plaza-cli/commit/c699ac9a5f5a219feb1caa46a3fd6af583d0bf9a))
* no longer require an API key when building on production repos ([eeeb97e](https://github.com/plazafyi/plaza-cli/commit/eeeb97e80522482717801b04100669d9bd91bc39))


### Chores

* configure new SDK language ([06e4233](https://github.com/plazafyi/plaza-cli/commit/06e4233f00294b83a6c4c8febeeb6995551cbce8))
* configure new SDK language ([e0c67da](https://github.com/plazafyi/plaza-cli/commit/e0c67da095c21a42c7f7730e5b8260216a83d88f))
* configure new SDK language ([92e096b](https://github.com/plazafyi/plaza-cli/commit/92e096b5949afc2ba5b64b4ad868c626822e207b))
* **internal:** codegen related update ([6bcf0db](https://github.com/plazafyi/plaza-cli/commit/6bcf0dbbbc682b2d6a987f58fa2b40c8a0d2a2d6))
* **internal:** tweak CI branches ([4a0f98d](https://github.com/plazafyi/plaza-cli/commit/4a0f98db732c23c3951e3f9e2fe86d0c6bd671d0))
* update SDK settings ([f523be7](https://github.com/plazafyi/plaza-cli/commit/f523be7e0116e1b298a398cb39f1422a00c7db0d))
* update SDK settings ([6353c09](https://github.com/plazafyi/plaza-cli/commit/6353c0923e68ef98f5542502511ee9d160280afc))


### Refactors

* **tests:** switch from prism to steady ([309372d](https://github.com/plazafyi/plaza-cli/commit/309372d5fb9ccf19a1de363a67a38668a13ebd46))
