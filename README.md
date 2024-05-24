# Init Oracle Monitor

Initia Oracle Monitor program for Initia chain validator operator.

## Monitor List

- Oracle Vote (VoteExtension)
- Oracle Double Sign

## Roadmap (TODO)

- Invalid price monitoring and alert
- Specify a monitoring height range
- Prometheus
- Alert
  - Slack
  - Pagerduty

## Support

- Alert
  - Telegram
- JSON-API
  - localhost:<listen_port>
  - Result Example
  ```json
  {
    "state": [
        {
            "status": true,
            "height": 297453,
            "block_sign": true,
            "oracle_sign": true,
            "double_sign": false
        },
        {
            "status": true,
            "height": 297454,
            "block_sign": true,
            "oracle_sign": true,
            "double_sign": false
        },
  }
  ```

## Quick Guide

1. Build

```bash
go build
```

2. Configure config.toml file

```bash
# You can get a example of config.toml file by below command
cp config.toml.example config.toml
```

3. Execute

```bash
./init-oracle-mon
```
