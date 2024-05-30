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
    "status": {
        "status": true,
        "oracle_miss_cnt": "0 / 30",
        "uptime": "100%",
        "WindowSize": 30,
        "OracleMissCnt": 0
    },
    "state": [
        {
            "height": 464351,
            "block_sign": true,
            "oracle_sign": true,
            "oracle_double_sign": false
        },
        {
            "height": 464352,
            "block_sign": true,
            "oracle_sign": true,
            "oracle_double_sign": false
        },
        ...
    ]
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
