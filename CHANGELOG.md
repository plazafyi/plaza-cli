# Changelog

## 0.2.0 (2026-05-01)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/plazafyi/plaza-cli/compare/v0.1.0...v0.2.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([1d71156](https://github.com/plazafyi/plaza-cli/commit/1d71156f29107340af0bfa9704dcef36a42970ea))
* **api:** api update ([fbd7b3d](https://github.com/plazafyi/plaza-cli/commit/fbd7b3d39a85f5d0687494ae0d4cf0a7dd304a82))
* **api:** api update ([6263b6f](https://github.com/plazafyi/plaza-cli/commit/6263b6ff0c88377138b68e6a16e7cbfbb000d74a))
* **api:** api update ([1d4dbf4](https://github.com/plazafyi/plaza-cli/commit/1d4dbf4e38e101a9e787cb487b9b5396bf60dc16))
* **api:** api update ([11c4fa4](https://github.com/plazafyi/plaza-cli/commit/11c4fa408b62a53c7f1a3b5645944c9800431eb6))
* **api:** api update ([e0cfeb7](https://github.com/plazafyi/plaza-cli/commit/e0cfeb7d1f61403062ddb134f3c87ed59e23b758))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([9e72554](https://github.com/plazafyi/plaza-cli/commit/9e72554732a26308d0de97a0cb58a6bb1759ce6b))
* binary-only parameters become CLI flags that take filenames only ([dfc1330](https://github.com/plazafyi/plaza-cli/commit/dfc13309168bac56e07bfedbeec70f6d4a117fe8))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([1928348](https://github.com/plazafyi/plaza-cli/commit/1928348c4c4d47369ee188c4c3d19c2961bd3604))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([da142ab](https://github.com/plazafyi/plaza-cli/commit/da142ab976c1829c1bfb132f828b4123f3ceeccc))
* **cli:** send filename and content type when reading input from files ([b9bb425](https://github.com/plazafyi/plaza-cli/commit/b9bb42562b9d80137fd389d586674c9b05c537cb))
* set CLI flag constant values automatically where `x-stainless-const` is set ([9dfbdaf](https://github.com/plazafyi/plaza-cli/commit/9dfbdaf00c7059af0d9dc46278dda13e58d2ddf8))
* support passing path and query params over stdin ([1b5e07c](https://github.com/plazafyi/plaza-cli/commit/1b5e07cfa779c4d3eb430162a72af6cfd33b3547))


### Bug Fixes

* cli no longer hangs when stdin is attached to a pipe with empty input ([a0a80ea](https://github.com/plazafyi/plaza-cli/commit/a0a80ea0f7e9efa365ee53d46ce1a5139fc3786d))
* **cli:** correctly load zsh autocompletion ([325bbac](https://github.com/plazafyi/plaza-cli/commit/325bbace68cb27491688a2e854be6e95b5b364b4))
* **cli:** fix incompatible Go types for flag generated as array of maps ([86ad372](https://github.com/plazafyi/plaza-cli/commit/86ad3722c20b42fe945f9abb0644f62682a7acb6))
* fall back to main branch if linking fails in CI ([f379781](https://github.com/plazafyi/plaza-cli/commit/f37978149b096ae23684672f9861c4c4d4f3d0fd))
* fix for failing to drop invalid module replace in link script ([094e403](https://github.com/plazafyi/plaza-cli/commit/094e403b6b0f6f9efc044b8c0a95819d158db0fa))
* fix for off-by-one error in pagination logic ([a087dae](https://github.com/plazafyi/plaza-cli/commit/a087dae75e0dbb1654e8d8d7710e11596e1e4684))
* fix quoting typo ([31cfee5](https://github.com/plazafyi/plaza-cli/commit/31cfee5119603357a293ec3ce162ed15d50e4189))
* flags for nullable body scalar fields are strictly typed ([34c1070](https://github.com/plazafyi/plaza-cli/commit/34c10709ee9a120149d95199dd997ce8a006d1df))
* handle empty data set using `--format explore` ([0e3f4cd](https://github.com/plazafyi/plaza-cli/commit/0e3f4cd86ffc0b508bba15cefaa6bf0c458e5705))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([62ba831](https://github.com/plazafyi/plaza-cli/commit/62ba8310c27725ad10bfd9349204cc92c2eb890d))


### Chores

* add documentation for ./scripts/link ([a98a364](https://github.com/plazafyi/plaza-cli/commit/a98a3642d83048980c9635891087f572e6e5c6c6))
* **ci:** skip lint on metadata-only changes ([b83eaec](https://github.com/plazafyi/plaza-cli/commit/b83eaec313d6a95b12e807d3e9faf1dc485e7659))
* **ci:** support manually triggering release workflow ([47f5699](https://github.com/plazafyi/plaza-cli/commit/47f569940b6077cc63c1cb4d2f996516c2926531))
* **cli:** additional test cases for `ShowJSONIterator` ([1279fc7](https://github.com/plazafyi/plaza-cli/commit/1279fc798797695ecd270a05551988c2121f2382))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([2d3680c](https://github.com/plazafyi/plaza-cli/commit/2d3680c1e31c92607a52a48c54cdd13c2af68b82))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([37e351c](https://github.com/plazafyi/plaza-cli/commit/37e351cfd8881afb29b9091265c91179131c9937))
* **cli:** switch long lists of positional args over to param structs ([923e5c9](https://github.com/plazafyi/plaza-cli/commit/923e5c9b6139f3d7e598a34b55ca941ec2228b0c))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([6c4a074](https://github.com/plazafyi/plaza-cli/commit/6c4a0744e462d546fa3b44c97cbeb7b285cc0449))
* **internal:** more robust bootstrap script ([bc91d4d](https://github.com/plazafyi/plaza-cli/commit/bc91d4db25889d08ca0c0d769dabde5b2528ec26))
* **internal:** update gitignore ([bfc57a1](https://github.com/plazafyi/plaza-cli/commit/bfc57a1456f8deceb216fb82c06acc664fbbe556))
* **internal:** update multipart form array serialization ([5cef7e1](https://github.com/plazafyi/plaza-cli/commit/5cef7e13df9ad05099815863114914a34a26f615))
* mark all CLI-related tests in Go with `t.Parallel()` ([0416eef](https://github.com/plazafyi/plaza-cli/commit/0416eef6aed42d27b5fe895007f755c1d6b062ec))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([da317b2](https://github.com/plazafyi/plaza-cli/commit/da317b2bbb933eb383848f88bc879affc4c12940))
* omit full usage information when missing required CLI parameters ([097fe19](https://github.com/plazafyi/plaza-cli/commit/097fe197e14d07f18be2da436c172cd9d1d9f8cd))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([cad0853](https://github.com/plazafyi/plaza-cli/commit/cad0853b47b63947e8233a57560ec679165195ae))
* **tests:** bump steady to v0.19.4 ([35b14fc](https://github.com/plazafyi/plaza-cli/commit/35b14fc1627bddcf05e8652158aeb1d1ec05f9a8))
* **tests:** bump steady to v0.19.5 ([7eab4a7](https://github.com/plazafyi/plaza-cli/commit/7eab4a7562c76a9deabb8cdd25caf6f33c1bb402))
* **tests:** bump steady to v0.19.6 ([508f853](https://github.com/plazafyi/plaza-cli/commit/508f853509084fabc9e7b30cacaf83ac49a715d2))
* **tests:** bump steady to v0.19.7 ([df70688](https://github.com/plazafyi/plaza-cli/commit/df706884a6f089fa774400993297357c413da159))
* **tests:** bump steady to v0.20.1 ([0567eb8](https://github.com/plazafyi/plaza-cli/commit/0567eb83a2bd06e84d08455a61c1669a435ab6bf))
* **tests:** bump steady to v0.20.2 ([05b9a87](https://github.com/plazafyi/plaza-cli/commit/05b9a870cfbbb985aac110c691c7427db1de03fb))
* **tests:** bump steady to v0.22.1 ([3bab44b](https://github.com/plazafyi/plaza-cli/commit/3bab44b71741d3f61c494857825452dbca20863f))

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
