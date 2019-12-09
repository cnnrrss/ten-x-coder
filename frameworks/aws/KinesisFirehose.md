# Kinesis Firehose

Kinesis Data Firehose is a **fully managed** service for delivering _near_ real-time streaming data to destinations such as **S3**, **Redshift**, **Elasticsearch**, and **Splunk**. With Firehose you don't need to write apps or manage resources. Amazon Kinesis Data Firehose provides a simple way to capture, transform, and load streaming data with just a few clicks in the AWS Management Console.

Configure producers to send data to Firehose and it automatically delivers the data to the destination that you specified. Types of producers:
- Kinesis data stream
- Kinesis Agent
- Kinesis Data Firehose API

Other Producers...
- CloudWatch Logs
- CloudWatch Events
- AWS IOT

You can also configure Firehose to transform data before delivery.

### Destinations
- S3: when s3 is selected as destination, and source record S3 backup is enabled, untransformed incoming data can be delivered to a separate S3 bucket and errors are delivered to processing-failed and errors folder in S3 bucket.
- Redshift:
- Elasticsearch:
- Splunk:

### Delivery Limits
Kinesis firehose data delivery frequency limits

| |S3|ES|Splunk|
|-|--|--|------|
|**Buffer size**|1-128 MB|1-100 MB|5 MB|
|**Buffer interval**|60-900 seconds|60-900 seconds|60 seconds|

- Each account can have up to 50 delivery streams per region.
- Kinesis Data Firehose delivery stream stores data records for up to **24 hours** in case the delivery destination is unavailable.
- Maximum size of a record sent to Kinesis Data Firehose, before base64-encoding, is **1000 KiB**
- `PutRecordBatch` operation can take up to 500 records per call or 4 MiB per call, whichever is smaller. This limit cannot be changed.

#### Lambda Blueprints

Kinesis Data Firehose provides the following Lambda blueprints that you can use to create a Lambda function for data transformation.

Firehose can transform using **Data Transformation** and **Record Format** techniques.

• **General Firehose Processing** — Contains the data transformation and status model described in the previous section. Use this blueprint for any custom transformation logic. \
• **Apache Log to JSON** — Parses and converts Apache log lines to JSONobjects, using predefined JSON field names. \
• **Apache Log to CSV** — Parses and converts Apache log lines to CSV format. \
• **Syslog to JSON** — Parses and converts Syslog lines to JSONobjects, using predefined JSON field names. \
• **Syslog to CSV** — Parses and converts Syslog lines to CSV format. \
• **Kinesis Data Firehose Process Record Streams as source** — Accesses the Kinesis Data Streams records in the input and returns them with a processing status. \
• **Kinesis Data Firehose CloudWatch Logs Processor** — Parses and extracts individual log events from records sent by CloudWatch Logs subscription filters.

• **Record format Conversion**: Firehose can convert JSON to Parquet or Apache ORC format before storing in S3. These formats are columnar that save space and enable faster queries compared to row-oriented formats like JSON.

Backup the syslog streaming data:
- When s3 is selected as destination, an source record S3 backup is enabled, untransformed incoming data can be delivered to a separate S3 bucket

Capture transform failures into same S3 bucket to address audit:
- Data Transformation failures are delivered to processing-failed folder