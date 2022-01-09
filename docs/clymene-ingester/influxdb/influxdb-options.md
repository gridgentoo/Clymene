## Clymene-ingester

Clymene-ingester consumes from Kafka and send to db.

```
Clymene-ingester [flags]
```

### Options

```
--admin.http.host-ports string                  The host:ports (e.g. 127.0.0.1:15694 or :15694) for the admin server, including health check, /metrics, etc. (default ":15694")
--clymene-ingester.deadlockInterval duration    Interval to check for deadlocks. If no messages gets processed in given time, clymene-ingester app will exit. Value of 0 disables deadlock check. (default 0s)
--clymene-ingester.parallelism string           The number of messages to process in parallel (default "1000")
--influxdb.bucket string                        influx bucket, A bucket is a named location where time series data is stored
--influxdb.http.http-request-timeout duration   HTTP request timeout in sec (default 10s)
--influxdb.org string                           influx organization, An organization is a workspace for a group of users.
--influxdb.tls.ca string                        Path to a TLS CA (Certification Authority) file used to verify the remote server(s) (by default will use the system truststore)
--influxdb.tls.cert string                      Path to a TLS Certificate file, used to identify this process to the remote server(s)
--influxdb.tls.enabled                          Enable TLS when talking to the remote server(s)
--influxdb.tls.key string                       Path to a TLS Private Key file, used to identify this process to the remote server(s)
--influxdb.tls.server-name string               Override the TLS server name we expect in the certificate of the remote server(s)
--influxdb.tls.skip-host-verify                 (insecure) Skip server's certificate chain and host name verification
--influxdb.token string                         Use the Authorization header and the Token scheme
--influxdb.url string                           the influxdb url (default "http://localhost:8086")
--influxdb.write.batch-size uint                Maximum number of points sent to server in single request (default 5000)
--influxdb.write.default-tags string            Tags added to each point during writing. separated by , (TAG1=VALUE1,TAG2=VALUE2)
--influxdb.write.exponential-base uint          The base for the exponential retry delay (default 2)
--influxdb.write.flush-interval uint            Interval, in ms, in which is buffer flushed if it has not been already written (by reaching batch size) (default 1000)
--influxdb.write.max-retries uint               Maximum count of retry attempts of failed writes (default 5)
--influxdb.write.max-retry-interval uint        The maximum delay between each retry attempt in milliseconds (default 125000)
--influxdb.write.max-retry-time uint            The maximum total retry timeout in millisecond (default 180000)
--influxdb.write.precision duration             Precision to use in writes for timestamp. In unit of duration: time.Nanosecond, time.Microsecond, time.Millisecond, time.Second (default 1ns)
--influxdb.write.retry-buffer-limit uint        Maximum number of points to keep for retry. Should be multiple of BatchSize (default 50000)
--influxdb.write.retry-interval uint            Default retry interval in ms, if not sent by server (default 5000)
--influxdb.write.use-gzip                       Whether to use GZip compression in requests
--kafka.consumer.authentication string          Authentication type used to authenticate with kafka cluster. e.g. none, kerberos, tls, plaintext (default "none")
--kafka.consumer.brokers string                 The comma-separated list of kafka brokers. i.e. '127.0.0.1:9092,0.0.0:1234' (default "127.0.0.1:9092")
--kafka.consumer.client-id string               The Consumer Client ID that clymene-ingester will use (default "clymene")
--kafka.consumer.encoding string                The encoding of metrics ("json", "protobuf") consumed from kafka (default "protobuf")
--kafka.consumer.group-id string                The Consumer Group that clymene-ingester will be consuming on behalf of (default "clymene")
--kafka.consumer.kerberos.config-file string    Path to Kerberos configuration. i.e /etc/krb5.conf (default "/etc/krb5.conf")
--kafka.consumer.kerberos.keytab-file string    Path to keytab file. i.e /etc/security/kafka.keytab (default "/etc/security/kafka.keytab")
--kafka.consumer.kerberos.password string       The Kerberos password used for authenticate with KDC
--kafka.consumer.kerberos.realm string          Kerberos realm
--kafka.consumer.kerberos.service-name string   Kerberos service name (default "kafka")
--kafka.consumer.kerberos.use-keytab            Use of keytab instead of password, if this is true, keytab file will be used instead of password
--kafka.consumer.kerberos.username string       The Kerberos username used for authenticate with KDC
--kafka.consumer.plaintext.mechanism string     The plaintext Mechanism for SASL/PLAIN authentication, e.g. 'SCRAM-SHA-256' or 'SCRAM-SHA-512' or 'PLAIN' (default "PLAIN")
--kafka.consumer.plaintext.password string      The plaintext Password for SASL/PLAIN authentication
--kafka.consumer.plaintext.username string      The plaintext Username for SASL/PLAIN authentication
--kafka.consumer.protocol-version string        Kafka protocol version - must be supported by kafka server
--kafka.consumer.tls.ca string                  Path to a TLS CA (Certification Authority) file used to verify the remote server(s) (by default will use the system truststore)
--kafka.consumer.tls.cert string                Path to a TLS Certificate file, used to identify this process to the remote server(s)
--kafka.consumer.tls.enabled                    Enable TLS when talking to the remote server(s)
--kafka.consumer.tls.key string                 Path to a TLS Private Key file, used to identify this process to the remote server(s)
--kafka.consumer.tls.server-name string         Override the TLS server name we expect in the certificate of the remote server(s)
--kafka.consumer.tls.skip-host-verify           (insecure) Skip server's certificate chain and host name verification
--kafka.consumer.topic string                   The name of the kafka topic to consume from (default "clymene")
```

###### Auto generated by spf13/cobra on 7-Jan-2022