# Kinesis Firehose

Kinesis Data Firehose is a **fully managed** service for delivering real-time streaming data to destinations such as **S3**, **Redshift**, **ES**, and **Splunk**. With Firehose you don't need to write apps or manage resources.

Configure producers to send data to Firehose and it automatically delivers the data to the destination that you specified. Types of producers:
- Kinesis data stream
- Kinesis Agent
- Kinesis Data Firehose API

Other Producers...
- CloudWatch Logs
- CloudWatch Events
- AWS IOT

You can also configure Firehose to transform data before delivery.

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

Apply transform of syslog data to JSON and CSV format:
- Data transformation from syslog to JSON and CSV can be performed through Lambda blueprints

Backup the syslog streaming data:
- When s3 is selected as destination, an source record S3 backup is enabled, untransformed incoming data can be delivered to a separate S3 bucket

Capture transform failures into same S3 bucket to address audit:
- Data Transformation failures are delivered to processing-failed folder