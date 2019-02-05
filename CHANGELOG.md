# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),

## [Unreleased]
### Added
- This Changelog.
- Automated release process on Git Tag push.
- `--version` command line option.

### Changed
- The default listening port is now `9536` and registered as a [Prometheus Exporter Default port](https://github.com/prometheus/prometheus/wiki/Default-port-allocations).
  If you want to maintain the previous behaviour, launch the `ipsec_exporter` with the `--web.listen-address 9101`
  command line flag. 

## [0.2](https://github.com/dennisstritzke/ipsec_exporter/releases/tag/v0.2) - 2018-07-30
### Added
- Connections containing `auto=ignored` are reported as ignored (ipsec_status = 4)

## [0.1.2.1](https://github.com/dennisstritzke/ipsec_exporter/releases/tag/v0.1.2.1) - 2018-06-05
### Fixed
- Concurrent read and write operations on map containing the IPsec configuration.

## [0.1.2](https://github.com/dennisstritzke/ipsec_exporter/releases/tag/v0.1.2) - 2018-05-09
### Added
- Support for connection names that contain numbers.

### Changed
- Golang dependency management to [Glide](https://github.com/Masterminds/glide). 

## [0.1.1](https://github.com/dennisstritzke/ipsec_exporter/releases/tag/v0.1.1) - 2018-01-29
### Added
- Checking, if the config file provided is readable.

### Changed
- Warns, if there are no connections configured in the IPsec config file. 
- Warns, if IPsec status couldn't be determined for a connection.

## [0.1](https://github.com/dennisstritzke/ipsec_exporter/releases/tag/v0.1) - 2018-01-28 
### Added
- Detection of configured IPsec tunnels by reading the `ipsec.conf` file.
- Prometheus metrics, indicate if the tunnel is up, the connection is up or the tunnel is down.

[Unreleased]: https://github.com/dennisstritzke/ipsec_exporter/compare/v0.2...HEAD
[0.2]: https://github.com/dennisstritzke/ipsec_exporter/compare/v0.1.2.1...v0.2
[0.1.2.1]: https://github.com/dennisstritzke/ipsec_exporter/compare/v0.1.2...v0.1.2.1
[0.1.2]: https://github.com/dennisstritzke/ipsec_exporter/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/dennisstritzke/ipsec_exporter/compare/v0.1...v0.1.1