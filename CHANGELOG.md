# Changelog

## [3.0.0](https://github.com/jensschulze/php-fpm_exporter/compare/v2.3.1...v3.0.0) (2026-01-27)


### ⚠ BREAKING CHANGES

* remove `--prometheus.const-label`

### Features

* add `--web.phpfpm-metrics-only` flag ([1435d9b](https://github.com/jensschulze/php-fpm_exporter/commit/1435d9b72fcbbfd97eeec5e77d9afc6fbc33c157))


### Bug Fixes

* **deps:** update module github.com/sirupsen/logrus to v1.9.4 ([6037755](https://github.com/jensschulze/php-fpm_exporter/commit/60377557faaec16ae08830031c3789a40a5ef78e))
* **deps:** update module github.com/sirupsen/logrus to v1.9.4 ([5a76f33](https://github.com/jensschulze/php-fpm_exporter/commit/5a76f331b579bc1dc8295d9e4175762f9faeef0e))
* **deps:** update module github.com/spf13/cobra to v1.10.2 ([4ddace5](https://github.com/jensschulze/php-fpm_exporter/commit/4ddace5cff0178b7f94540a059cc41a8fe835bd6))
* **deps:** update module github.com/spf13/cobra to v1.10.2 ([091fcdb](https://github.com/jensschulze/php-fpm_exporter/commit/091fcdb077a1586b2ecf278af9beff980cf7a463))

## [2.3.1](https://github.com/jensschulze/php-fpm_exporter/compare/v2.3.0...v2.3.1) (2025-10-22)


### Bug Fixes

* apply const labels to runtime metrics ([9782bb5](https://github.com/jensschulze/php-fpm_exporter/commit/9782bb5dff46a0486abda1c545521ac4b0e35fe0))

## [1.0.1](https://github.com/jensschulze/php-fpm_exporter/compare/v1.0.0...v1.0.1) (2025-10-11)


### Bug Fixes

* **ci:** mixup of Docker Hub and Github repo name ([cf9a95d](https://github.com/jensschulze/php-fpm_exporter/commit/cf9a95d981c93a60a7418c3a0ecab9375aaa34b8))

## 1.0.0 (2025-10-11)


### ⚠ BREAKING CHANGES

* `pid_hash` is being removed in favour of `child` to avoid high cardinality explosion. In turn this means processes and their state changes can't be identified anymore. If you're using this behaviour please open an issue.

### Features

* add custom labels ([a67b9f4](https://github.com/jensschulze/php-fpm_exporter/commit/a67b9f485bb7f9d38d7644be8ca210aab2c3b324))
* add scrape_url as prometheus label for each metric ([#72](https://github.com/jensschulze/php-fpm_exporter/issues/72)) ([505fe34](https://github.com/jensschulze/php-fpm_exporter/commit/505fe34b2c5e6e4ec5010df1c082beeb2666a030))
* build with Go 1.14.10 ([0a03abe](https://github.com/jensschulze/php-fpm_exporter/commit/0a03abee4b4bac9657483faed4ed86dd84fcab04))
* create multiarch docker images using buildx ([#226](https://github.com/jensschulze/php-fpm_exporter/issues/226)) ([1d33182](https://github.com/jensschulze/php-fpm_exporter/commit/1d3318235be16eb8514bf72f172b128a400d7def))
* customization ([9712014](https://github.com/jensschulze/php-fpm_exporter/commit/97120144521d340b571a73aeb5d15a444469cc42))
* publish docker images to GitHub Registry ([8647200](https://github.com/jensschulze/php-fpm_exporter/commit/86472005d053d8b1bbf27139ce6a443c35b47f53))


### Bug Fixes

* align proposed parameter with actual name. ([#208](https://github.com/jensschulze/php-fpm_exporter/issues/208)) ([877a9ff](https://github.com/jensschulze/php-fpm_exporter/commit/877a9ffa0a7dd78323fb230e7842bd9d551987ba))
* allow 10 seconds to shutdown, fixing shutdown error ([#204](https://github.com/jensschulze/php-fpm_exporter/issues/204)) ([ea4a929](https://github.com/jensschulze/php-fpm_exporter/commit/ea4a929fa9123c6c820ec8703bb77d3d44cbeacc))
* **ci:** remove integration tests ([e50e8c1](https://github.com/jensschulze/php-fpm_exporter/commit/e50e8c108b7a3aeab19c604633c7aec2f884e5b9))
* **ci:** remove wrong directories from tests ([02a8580](https://github.com/jensschulze/php-fpm_exporter/commit/02a8580b52982c84e01cb52a86fe19ea7fb70f65))
* generate gauge metrics for all states ([#173](https://github.com/jensschulze/php-fpm_exporter/issues/173)) ([26cc9ad](https://github.com/jensschulze/php-fpm_exporter/commit/26cc9ada6672b09a58528d8851a91a137887e937))
* high cardinality of pid_hash ([#124](https://github.com/jensschulze/php-fpm_exporter/issues/124)) ([0d25732](https://github.com/jensschulze/php-fpm_exporter/commit/0d25732143e2777409f7d44d2b34eb4eb79546b7))
* Invalid PHP-FPMs request uri encoding ([#37](https://github.com/jensschulze/php-fpm_exporter/issues/37)) ([8b64a58](https://github.com/jensschulze/php-fpm_exporter/commit/8b64a58900a3a599ce1c0086951e478998dc6b3e))
* json: cannot unmarshal number ([#28](https://github.com/jensschulze/php-fpm_exporter/issues/28)) ([03d8708](https://github.com/jensschulze/php-fpm_exporter/commit/03d87088d79054ac2f736d6b30cb0bd05a11f37d))
* Label pool.Name on metric was missing ([510d087](https://github.com/jensschulze/php-fpm_exporter/commit/510d087da9a3892044f6a6cd9fc08f3ee8c9b02f))
* publish step for semantic release missing ([4dd1801](https://github.com/jensschulze/php-fpm_exporter/commit/4dd180146368fd4a233bf24701b483e7e1f64b6c))
* release pipeline failure ([a727363](https://github.com/jensschulze/php-fpm_exporter/commit/a7273638add963fb3605ee4a687d9e80289a8a03))
* supply tar.gz archive ([#172](https://github.com/jensschulze/php-fpm_exporter/issues/172)) ([0eaefb8](https://github.com/jensschulze/php-fpm_exporter/commit/0eaefb844b73202456b668a2cc9fd18bfce451b5)), closes [#171](https://github.com/jensschulze/php-fpm_exporter/issues/171)
* support capturing SIGTERM signal to shutdown gracefully ([#82](https://github.com/jensschulze/php-fpm_exporter/issues/82)) ([271748e](https://github.com/jensschulze/php-fpm_exporter/commit/271748ee09756cb9ade06e4d2e64414524540010))
* trigger release with arm ([65a8273](https://github.com/jensschulze/php-fpm_exporter/commit/65a8273ea259e3c1ef91bdc6967db710abf1cf70))
* Typo FPM_REQUEST_INFO ([#134](https://github.com/jensschulze/php-fpm_exporter/issues/134)) ([cf49da4](https://github.com/jensschulze/php-fpm_exporter/commit/cf49da44660b9e33d31d368fc510169e8574f7c2))
