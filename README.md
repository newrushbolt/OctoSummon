# OctoSummon

**OctoSummon** is an prometheus-compatible alerts processing router.
Is purposed to give some advanced paging logic similar to Zabbix/Pagerduty.

### Prometheus Alertmanager config

```
---

route:
    receiver: 'OctoSummon'

receivers:
    - name: 'OctoSummon'
      webhook_configs:
        - send_resolved: true
          url: "http://10.10.0.1:8000/alerts"
```

### Envs

**OCTOSUMMON_DEBUG** - enable debug logging
`export OCTOSUMMON_DEBUG="true"`


#### Manual build and run

```
make deps && make && make install
$GOPATH/bin/OctoSummon
```

### Docker build and run

```
docker build -t OctoSummon .
docker run -d OctoSummon
```
