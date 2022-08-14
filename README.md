Assignment submission for chaos-theory by Manish Dhal


## Start Project

```
docker compose up
```

## SSH into pod

#### Note :- If any issues on connecting to the pod from localhost, then ssh into pod and execute the commands


```
➜  ~ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS                   PORTS                               NAMES
52510ad33635   go-chaos_api   "go run cmd/main/mai…"   4 minutes ago   Up 4 minutes             127.0.0.1:9010->9010/tcp            go-chaos_api_1
43a36128d23c   mysql:latest   "docker-entrypoint.s…"   4 minutes ago   Up 4 minutes (healthy)   0.0.0.0:3306->3306/tcp, 33060/tcp   db
➜  ~ docker exec -it 52510ad33635 /bin/sh
```

## Add Entry

```
curl --location --request POST 'localhost:9010/lists' \
--header 'Content-Type: application/json' \
--data-raw '{
    "TempKey" : "Naruto",
    "TempValue": "Village Hidden in the Leaf"
}'
```

## Get All Entries

```
curl --location --request GET 'localhost:9010/lists'
```
