# OpenTelemetry Metric

The `opentelemetry` serializer converts Telegraf metrics into the [OpenTelemetry Metrics protobuf format][otel-format]. When used with the `opentelemetry` output, this will 

[otel-format]: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/metrics/v1/metrics.proto

### Configuration


```toml
[[inputs.file]]
  files = ["stdout"]

  ## Describe variables using the standard SampleConfig style.
  ##   https://github.com/influxdata/telegraf/wiki/SampleConfig
  example_option = "example_value"

  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ##   https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
  data_format = "opentelemetry"
```

#### example_option

If an option requires a more expansive explanation than can be included inline
in the sample configuration, it may be described here.

### Metrics

Metrics are processed as batches. TODO

### Example

**Example Input**
```
cpu,cpu=cpu0,host=loaner time_active=202224.15999999992,time_guest=30250.35,time_guest_nice=0,time_idle=1527035.04 1568760922000000000
cpu,cpu=cpu0,host=loaner usage_active=31.249999981810106,usage_guest=2.083333333080696,usage_guest_nice=0,usage_idle=68.7500000181899 1568760922000000000
cpu,cpu=cpu1,host=loaner time_active=201890.02000000002,time_guest=30508.41,time_guest_nice=0,time_idle=264641.18 1568760922000000000
cpu,cpu=cpu1,host=loaner usage_active=12.500000010610771,usage_guest=2.0833333328280585,usage_guest_nice=0,usage_idle=87.49999998938922 1568760922000000000
```

**Example Output (as JSON)** 
```json
{
  "metrics": [
    {
      "name": "time_active",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 202224.15999999992
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 201890.02000000002
          }
        ]
      }
    },
    {
      "name": "time_guest",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 30250.35
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 30508.41
          }
        ]
      }
    },
    {
      "name": "time_guest_nice",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 0
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 0
          }
        ]
      }
    },
    {
      "name": "time_idle",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 1527035.04
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 264641.18
          }
        ]
      }
    },
    {
      "name": "usage_active",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 31.249999981810106
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 12.500000010610771
          }
        ]
      }
    },
    {
      "name": "usage_guest",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 2.083333333080696
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 2.0833333328280585
          }
        ]
      }
    },
    {
      "name": "usage_guest_nice",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 0
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 0
          }
        ]
      }
    },
    {
      "name": "usage_idle",
      "data": {
        "data_points": [
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu0"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 68.7500000181899
          },
          {
            "labels": [
              {
                "key": "cpu",
                "value": "cpu1"
              },
              {
                "key": "host",
                "value": "loaner"
              },
            ],
            "time_unix_nano": "1568760922000000000",
            "value": 87.49999998938922
          }
        ]
      }
    }
  ]
}
```