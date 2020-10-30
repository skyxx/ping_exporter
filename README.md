# ping_exporter
[![Build Status](https://travis-ci.org/skyxx/ping_exporter.svg)](https://travis-ci.org/skyxx/ping_exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/skyxx/ping_exporter)](https://goreportcard.com/report/github.com/skyxx/ping_exporter)

Prometheus exporter for ICMP echo requests using https://github.com/digineo/go-ping

This is a simple server that scrapes go-ping stats and exports them via HTTP for
Prometheus consumption. The go-ping library is build and maintained by Digineo GmbH.
For more information check the [source code][go-ping].

[go-ping]: https://github.com/digineo/go-ping

## Getting Started

### Config file

Targets can be specified in a YAML based config file:

```yaml
targets:
  - 8.8.8.8
  - 8.8.4.4
  - 2001:4860:4860::8888
  - 2001:4860:4860::8844
  - google.com
  

ping:
  interval: 2s
  timeout: 3s
  history-size: 42
  payload-size: 120
```

### Exported metrics

- `ping_rtt_best_ms`:          Best round trip time in millis
- `ping_rtt_worst_ms`:         Worst round trip time in millis
- `ping_rtt_median_ms`:        Median round trip time in millis
- `ping_rtt_mean_ms`:          Mean round trip time in millis
- `ping_rtt_std_deviation_ms`: Standard deviation in millis
- `ping_packet_sent`:          Number of sent Packets
- `ping_packet_loss`:          Number of loss Packets

Each metric has labels `ip` (the target's IP address), `ip_version`
(4 or 6, corresponding to the IP version), and `target` (the target's
name).

Additionally, a `ping_up` metric reports whether the exporter
is running (and in which version).


### Shell

To run the exporter:

```console
$ ./ping_exporter [options] target1 target2 ...
```

or

```console
$ ./ping_exporter --config.path my-config-file [options]
```

Help on flags:

```console
$ ./ping_exporter --help
```

Getting the results for testing via cURL:

```console
$ curl http://localhost:9427/metrics
```

## Contribute

Simply fork and create a pull-request. We'll try to respond in a timely fashion.

## License

MIT License, Copyright (c) 2018
Philip Berndroth [pberndro](https://twitter.com/pberndro)
Daniel Czerwonk [dan_nrw](https://twitter.com/dan_nrw)
