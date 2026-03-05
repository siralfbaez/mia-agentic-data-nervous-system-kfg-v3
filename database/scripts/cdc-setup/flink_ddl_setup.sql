CREATE TABLE signal_events (
    signal_id STRING,
    source_system STRING,
    payload STRING,
    ts TIMESTAMP(3),
    WATERMARK FOR ts AS ts - INTERVAL '5' SECOND
) WITH (
    'connector' = 'kafka',
    'topic' = 'mia.signals.raw',
    'properties.bootstrap.servers' = '${CONFLUENT_BOOTSTRAP}',
    'format' = 'avro-confluent',
    'avro-confluent.url' = '${SCHEMA_REGISTRY_URL}'
);
