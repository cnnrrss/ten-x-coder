# IOT

## Provisioning and Creation

- Enable a device with an _existing certificate_ or have AWS IOT create and register one for you
- Define device **shadow service** to get and set the state of a device over MQTT or HTTP
- Assign a **policy** attached to the certificate and Define unique indentifier for the thing (device).
- Define set of attributes for the thing, including existing thing **types** and **groups**


### Rule actions
AWS IOT rule actions are used to specifiy what to do when a rule is triggered. You can define actions to write data to DynamoDB or Kinesis Streams. Or invoke lambda, but the rule actions aren't required as part of provisioning.