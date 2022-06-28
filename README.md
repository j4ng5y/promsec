# PromSec

A central service to collect security metrics.

## Config File Schema

```yaml
version: string # default: v1
kind: string # default: PromsecConfig
server:
  host: string # default: 0.0.0.0
  port: uint # default: 5001
  endpoint: string # default: /metrics
  read_timeout: duration # default: 5s
  read_header_timeout: duration # default: 5s
  write_timeout: duration # default: 5s
  idle_timeout: duration # default: 10s
```
