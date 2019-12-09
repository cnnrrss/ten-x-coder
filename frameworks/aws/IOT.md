# IoT

## Provisioning and Creation

- Enable a device with an _existing certificate_ or have AWS IOT create and register one for you
- Define device **shadow service** to get and set the state of a device over MQTT or HTTP
- Assign a **policy** attached to the certificate and Define unique indentifier for the thing (device).
- Define set of attributes for the thing, including existing thing **types** and **groups**

### Rules for AWS IoT

You can use rules to support tasks like these:

- Augment or filter data received from a device.
- Write data received from a device to an Amazon DynamoDB - database.
- Save a file to Amazon S3.
- Send a push notification to all users using Amazon SNS.
- Publish data to an Amazon SQS queue.
- Invoke a Lambda function to extract data.
- Process messages from a large number of devices using - Amazon Kinesis.
- Send data to the Amazon Elasticsearch Service.
- Capture a CloudWatch metric.
- Change a CloudWatch alarm.
- Send the data from an MQTT message to Amazon Machine - Learning to make predictions based on an Amazon ML model.
- Send a message to a Salesforce IoT Input Stream.
- Send message data to an AWS IoT Analytics channel.
- Start execution of a Step Functions state machine.
- Send message data to an AWS IoT Events input.

IOT Rules address filtering and integration with DynamoDB for persistent storage, Elasticsearch for search, S3 for storage, SNS for notifications, etc..

### Rule actions

AWS IOT rule actions are used to specifiy what to do when a rule is triggered. You can define actions to write data to DynamoDB or Kinesis Streams. Or invoke lambda, but the **rule actions aren't required as part of provisioning**.

### AWS IoT Artifacts

- **Device Gateway** - to securely and efficiently communicate with AWS IoT
- **Message broker** - to provide a secure mechanism for devices and AWS IoT applications to publish and receive messages from each other
- **Rules engine** - Provides message processing and integration with other AWS services (like S3, DynamoDB, Lambda)
- **Device shadow** - A JSON document (not store) used to store and retrieve current state information for a device.
- **Device Shadow service** - to provide persistent representations of your devices in the AWS Cloud. You can publish updated state information to a device's shadow, and your device can synchronize its state when it connects.
- **Device Provisioning service** - Allows you to provision devices using a template that describes the resources required for your device: a thing, a certificate, and one or more policies.
- **Custom Authentication service** - You can define custom authorizers that allow you to manage your own authentication and authorization strategy using a custom authentication service and a Lambda function.
- **Jobs service** - Allows you to define a set of remote operations that are sent to and executed on one or more devices connected to AWS IoT.
- **Security and Identity service** - Provides shared responsibility for security in the AWS Cloud.
- **Registry** - Organizes the resources associated with each device in the AWS Cloud. ou register your devices and associate up to three custom attributes with each one. You can also associate certificates and MQTT client IDs with each device to improve your ability to manage and troubleshoot them.
- **Group registry** - Groups allow you to manage several devices at once by categorizing them into groups. Groups can also contain groups—you can build a hierarchy of groups.