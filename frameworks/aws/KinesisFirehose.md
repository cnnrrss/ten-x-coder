## Kinesis Firehose

#### Lambda Blueprints

Kinesis Data Firehose provides the following Lambda blueprints that you can use to create a Lambda function for data transformation.

Firehose can transform using **Data Transformation** and **Record Format** techniques.

• **General Firehose Processing** — Contains the data transformation and status model described in the previous section. Use this blueprint for any custom transformation logic. \
• **Appache Log to JSON** — Parses and converts Apache log lines to JSONobjects, using predefined JSON field names. \
• **Apache Log to CSV** — Parses and converts Apache log lines to CSV format. \
• **Syslog to JSON** — Parses and converts Syslog lines to JSONobjects, using predefined JSON field names. \
• **Syslog to CSV** — Parses and converts Syslog lines to CSV format. \
• **Kinesis Data Firehose Process Record Streams as source** — Accesses the Kinesis Data Streams records in the input and returns them with a processing status. \
• **Kinesis Data Firehose CloudWatch Logs Processor** — Parses and extracts individual log events from records sent by CloudWatch Logs subscription filters.


• **Record format Conversion**: Firehose can convert JSON to Parquet or Apache ORC format before storing in S3. These formats are columnar that save space and enable faster queries compared to row-oriented formats like JSON.