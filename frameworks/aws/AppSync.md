# App Sync

A managed, serverless GraphQL service

Connect to data sources in youraccount

Enterprise security features: IAM, Cognito, OIDC, API keys


Architecture UnicornTrivia (Real-Time Interactions)

            mutation -> AppSync -> subscription (question)
Admin Panel            (resolver)             Mobile App
                        DynamoDB


# Amplify

CLI serverless backends for apps
Automatically creates AWS CloudFormation templates
Deploy mobile applications ios, android, web, react native
Provide boiler plate code
Open-source to work with community roadmap


### Workshop

`amplify init`

`amplify video add`

```
MediaLive
MediaLive Primary Ingest Url: rtmp://100.21.191.232:1935/mylivestream-dev-p
MediaLive Primary Stream Key: mylivestream-dev-p

MediaLive Backup Ingest Url: rtmp://50.112.158.113:1935/mylivestream-dev-b
MediaLive Backup Stream Key: mylivestream-dev-b

MediaStore
MediaStore Output Url: https://q73s6nu23ah4rk.data.mediastore.us-west-2.amazonaws.com/p/index.m3u8
```

Setup Streaming

https://github.com/awslabs/aws-amplify-unicorntrivia-workshop/blob/master/documentation/Live_stream_doc.md
https://github.com/awslabs/aws-amplify-unicorntrivia-workshop/blob/master/documentation/admin_panel_doc.md

Repo:
https://github.com/awslabs/aws-amplify-unicorntrivia-workshop

### Monetizing a Live Stream

igeraght@amazon.com - Ivan Geraghty

Elemental MediaLive -> Ele MediaStore/Ele MediaPackage/S3 -> MediaTailor

content personalization and monetization service that allows customers to implement stitched server-side ad insertion while maintaining high quality of service


