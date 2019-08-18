  # Google Cloud CLI cheat sheet

### switch projects
`gcloud config set project test-environment` \
`gcloud config set project prod`

### delete pub sub subscription
`gcloud pubsub subscriptions delete test-subscriber`

### create pub/sub subscription
```bash
gcloud pubsub subscriptions create test-subscriber \
--topic=feed \
--topic-project=test-environment
```

```bash
gcloud pubsub subscriptions create subscribe-feed \
--topic=feed \
--topic-project=prod-env
```

Examples:
```bash
  # List all pods in ps output format.
  kubectl get pods

  # List all pods in ps output format with more information (such as node name).
  kubectl get pods -o wide

  # List a single replication controller with specified NAME in ps output format.
  kubectl get replicationcontroller web

  # List a single pod in JSON output format.
  kubectl get -o json pod web-pod-13je7

  # List a pod identified by type and name specified in "pod.yaml" in JSON output format.
  kubectl get -f pod.yaml -o json

  # Return only the phase value of the specified pod.
  kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}}

  # List all replication controllers and services together in ps output format.
  kubectl get rc,services

  # List one or more resources by their type and names.
  kubectl get rc/web service/frontend pods/web-pod-13je7

  # List all resources with different types.
  kubectl get all

  # describe job
  kubectl describe job feeder-blah-12345 --namespace=it
```

`kubectl config set-context $(kubectl config current-context) --namespace=prod`

`docker exec -i -t my-app_1 psql -U postgres -d dev`

### Build docker image
```bash
docker build -t connor/my-app .
docker run -i -t us.gcr.io/<prod-env>/my-app:latest sh
docker ps
docker cp ~/Downloads/test-environment.json 35fa70b036e0:/
export GOOGLE_APPLICATION_CREDENTIALS=./test-environment.json
echo $GOOGLE_APPLICATION_CREDENTIALS
python3 -m my-app-module --gcp_project=test-environment-170018 \
  --pg_connection='postgresql://postgres:happy-happy-joy-joy@10.163.104.234/dev' \
  --subscription=test-subscription-1

docker run -i -t connor/my-app:latest python3 -m my-app-module --gcp_project=test-environment \
 --pg_connection='postgresql://postgres:happy-happy-joy-joy@10.163.104.234/dev' \
 --subscription=test-subscription-1

docker run -v /tmp:/tmp -e GOOGLE_APPLICATION_CREDENTIALS=/tmp/test-environment.json -i -t connor/my-app:latest python3 -m my-app-module --gcp_project=test-environment \
--pg_connection="postgresql://postgres:happy-happy-joy-joy@10.163.104.234/dev" \
--subscription=test-subscription-1

docker run -v /tmp:/tmp -e GOOGLE_APPLICATION_CREDENTIALS=/tmp/test-environment.json -i -t connor/my-app:latest python3 -m subscriber
```

## Postgres
### connect to cloud

`gcloud sql connect development`
`gcloud sql connect production`


**some posgres magic**

List all relations:

`\dt+ *.*` \
`\c name-space`


```sql
WITH "names"("name") AS (
  SELECT n.nspname AS "name"
    FROM pg_catalog.pg_namespace n
      WHERE n.nspname !~ '^pg_'
        AND n.nspname <> 'information_schema'
) SELECT "name",
  pg_catalog.has_schema_privilege("awe_user", "name", 'CREATE') AS "create",
  pg_catalog.has_schema_privilege("awe_user", "name", 'USAGE') AS "usage"
    FROM "names";
```

```sql
SELECT grantee, privilege_type 
FROM information_schema.role_table_grants 
WHERE table_name='config.table_name';
```

```sql
SELECT * FROM information_schema.role_table_grants 
WHERE table_name='table_name' OR table_name='table_name_2';
```

```sql
GRANT INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER 
ON config.table_name TO config_group;
```

## Secrets

**Create from file**
```bash
kubectl create secret generic storage-key \
--from-file=~/gcloud/key.json
```

```bash
kubectl create secret generic storage-key \
--namespace=test \
--from-file=./gcloud/key.json
```


## Pods

**Delete**
```bash
kubectl delete pod my-app-deployment-99d44696f-n54bs \
--grace-period=0 \
--force
```