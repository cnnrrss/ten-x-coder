# GKE quickstart

*Create a cluster*

`gcloud container clusters create [CLUSTER_NAME]`

*Auth cluster*
`gcloud container clusters get-credentials [CLUSTER_NAME]`

*Deploy app to cluster*

`kubectl create deployment hello-server --image=gcr.io/google-samples/hello-app:1.0`

Note, hello-app is packaged as docker container image, the image is stored in Google container registry.

*Expose the deployment*
```bash
kubectl expose deployment hello-server \
--type LoadBalancer \
--port 80 \
--target-port 8080
```

*Inspect pods*
`kubectl get pods`

*Inspect service*
`kubectl get service hello-server`

#### Clean up

*Delete service*
`kubectl delete service hello-server`

*Delete cluster*
`gcloud container clusters delete [CLUSTER_NAME]`
