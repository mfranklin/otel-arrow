# otel-arrow-adapter

Adapter used to convert OTEL batches to/from OTEL Arrow batches in both directions.

See [OTEP 0156](https://github.com/lquerel/oteps/blob/main/text/0156-columnar-encoding.md) for more details on the OTEL Arrow protocol.

## Packages

| Package   | Description                                                                                                                  |
|-----------|------------------------------------------------------------------------------------------------------------------------------|
| air       | Arrow Intermediate Representation used to translate batch of row-oriented entities into columnar-oriented batch of entities. |
| benchmark | Benchmark infrastructure to compare OTLP and OTLP Arrow representations.                                                     |
| datagen   | Metrics, logs, and traces generator (fake data).                                                                             |
| otel      | Conversion functions to translate OTLP entities into OTLP Arrow Event and vice versa.                                        |

## Tools

| Tool              | Description                                                      |
|-------------------|------------------------------------------------------------------|
| logs_gen          | Generates fake logs and store them as protobuf binary files.     |
| metrics_gen       | Generates fake metrics and store them as protobuf binary files.  |
| traces_gen        | Generates fake traces and store them as protobuf binary files.   |
| logs_benchmark    | Benchmark tool to compare OTLP and OTLP Arrow representations.   |
| metrics_benchmark | Benchmark tool to compare OTLP and OTLP Arrow representations.   |
| traces_benchmark  | Benchmark tool to compare OTLP and OTLP Arrow representations.   |

## Status [WIP]

### Arrow Intermediate Representation (framework to convert row-oriented structured data to Arrow columnar data)
- [X] Values (supported types: bool, i[8|16|32|64], u[8|16|32|64], f[32|64], string, binary, list, struct)
- [X] Fields
- [X] Record
- [X] Record Builder
- [X] Record Repository
- [X] Generate Arrow records
  - [X] Scalar values
  - [X] Struct values
  - [X] List values (except list of list)
- [X] Optimizations
  - [X] Dictionary encoding for string fields
  - [X] Dictionary encoding for binary fields
  - [X] Multi-field sorting (string field)
  - [X] Multi-field sorting (binary field)

### OTLP --> OTLP Arrow
  - **General**
    - [X] Complex attributes
    - [X] Complex body
  - **OTLP metrics --> OTLP_ARROW events**
    - [X] Gauge
    - [X] Sum
    - [X] Summary
    - [X] Histogram
    - [X] Exponential histogram
    - [X] Univariate metrics to multivariate metrics
    - [ ] Aggregation temporality
    - [ ] Exemplar
  - **OTLP logs --> OTLP_ARROW events**
    - [X] Logs
  - **OTLP trace --> OTLP_ARROW events**
    - [X] Trace
    - [X] Links
    - [X] Events

### OTLP Arrow --> OTLP
  - **General**
    - [ ] Complex attributes
    - [ ] Complex body
  - **OTLP_ARROW events --> OTLP metrics**
    - [ ] Gauge
    - [ ] Sum
    - [ ] Summary
    - [ ] Histogram
    - [ ] Exponential histogram
    - [ ] Univariate metrics to multivariate metrics
    - [ ] Aggregation temporality
    - [ ] Exemplar
  - **OTLP_ARROW events --> OTLP logs**
    - [ ] Logs
  - **OTLP_ARROW events --> OTLP trace**
    - [ ] Trace
    - [ ] Links
    - [ ] Events

### Protocol
  - [X] OTLP proto
  - [X] Event service
  - [x] BatchEvent producer
  - [X] BatchEvent consumer
  - [ ] gRPC service implementation

### Benchmarks 
  - Fake data generator
    - [X] ExportMetricsServiceRequest (except for histograms and summary)
    - [X] ExportLogsServiceRequest
    - [X] ExportTraceServiceRequest 
  - Framework to compare OTLP and OTLP_ARROW performances (i.e. size and time)
    - [X] General framework
    - [X] Compression algorithms (lz4 and zstd)
    - [X] Console output
    - [X] Export CSV
` - [X] OTLP batch creation + serialization + compression + decompression + deserialization
  - [X] OTLP_ARROW batch creation + serialization + compression + decompression + deserialization
` - [ ] Performance and memory optimizations
  - [ ] Check memory leaks (e.g. Arrow related memory leaks)

### Tools
  - [X] logs_gen
  - [X] metrics_gen
  - [X] traces_gen
  - [WIP] logs_benchmark
  - [WIP] metrics_benchmark
  - [X] traces_benchmark

### CI
  - [ ] GitHub Actions to build, test, check at every commit.

### Integration
  - [ ] Integration with Open Telemetry Collector.

### Feedback to implement
  - [ ] Feedback provided by @atoulme 