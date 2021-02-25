# OpenTelemetry Metric

The `opentelemetry` serializer converts Telegraf metrics into the [OpenTelemetry Metrics protobuf format][otel-format]. When used with the `opentelemetry` output, the 

[otel-format]: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/metrics/v1/metrics.proto

This description explains at a high level what the serializer does and
provides links to where additional information about the format can be found.

### Configuration

This section contains the sample configuration for the serializer.  Since the
configuration for a serializer is not have a standalone plugin, use the `file`
or `http` outputs as the base config.

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
  data_format = "example"
```

#### example_option

If an option requires a more expansive explanation than can be included inline
in the sample configuration, it may be described here.

### Metrics

The optional Metrics section contains details about how the serializer converts
Telegraf metrics into output.

### Example

**Example Input**
```
cpu,cpu=cpu0,host=loaner time_active=202224.15999999992,time_guest=30250.35,time_guest_nice=0,time_idle=1527035.04 1568760922000000000
cpu,cpu=cpu0,host=loaner usage_active=31.249999981810106,usage_guest=2.083333333080696,usage_guest_nice=0,usage_idle=68.7500000181899 1568760922000000000
cpu,cpu=cpu1,host=loaner time_active=201890.02000000002,time_guest=30508.41,time_guest_nice=0,time_idle=264641.18 1568760922000000000
cpu,cpu=cpu1,host=loaner usage_active=12.500000010610771,usage_guest=2.0833333328280585,usage_guest_nice=0,usage_idle=87.49999998938922 1568760922000000000
cpu,cpu=cpu2,host=loaner time_active=201382.78999999998,time_guest=30325.8,time_guest_nice=0,time_idle=264686.63 1568760922000000000
cpu,cpu=cpu2,host=loaner usage_active=15.999999993480742,usage_guest=1.9999999999126885,usage_guest_nice=0,usage_idle=84.00000000651926 1568760922000000000
cpu,cpu=cpu3,host=loaner time_active=198953.51000000007,time_guest=30344.43,time_guest_nice=0,time_idle=265504.09 1568760922000000000
cpu,cpu=cpu3,host=loaner usage_active=10.41666667424579,usage_guest=0,usage_guest_nice=0,usage_idle=89.58333332575421 1568760922000000000
cpu,cpu=cpu-total,host=loaner time_active=804450.5299999998,time_guest=121429,time_guest_nice=0,time_idle=2321866.96 1568760922000000000
cpu,cpu=cpu-total,host=loaner usage_active=17.616580305880305,usage_guest=1.036269430422946,usage_guest_nice=0,usage_idle=82.3834196941197 1568760922000000000
```

**Example Output (as JSON)** 
```
{
  "metrics": [
  {
    "name": "metric name",
    "description": "my metric description",
    "unit": "ms",
    "data": {
      "data_points": [
        {
          "labels": [
            {
            "key": "",
            "value": ""
            }
          ],
          "start_time_unix_nano": "0",
          "time_unix_nano": "123154123",
          "value": 22
        }
      ]
    }
  }
]
}
```