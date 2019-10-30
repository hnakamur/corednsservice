corednsservice
==============

corednsservice is a wrapper executable to install CoreDNS as a service on Windows.

An example `corednsservice.yml`.
Put this in the same directory as `corednsservice.exe` and modify it appropriately.

```
name: CoreDNS
display_name: CoreDNS service
description: CoreDNS service for local development.
exec: "C:\\CoreDNS\\coredns.exe"
args: ["-conf", "Corefile"]
dir: "C:\\CoreDNS"
stdout:
  filename: "C:\\CoreDNS\\coredns.log"
  maxsize: 100
  maxbackups: 50
  maxage: 30
  compress: true
```
