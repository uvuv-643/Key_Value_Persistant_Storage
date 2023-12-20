# Key-Value Persistant Storage
#### Golang + MongoDB

Golang Microservice with MongoDB. \
Microservice supports operations:
1. ```Set(key string, extension string, ttl int, value any)```
2. ```Get(key string)```
3. ```Delete(key string)```

When creating new object in storage you have to specify type: 
```text```, ```photo```, ```audio```, ```video```, with its extension etc. \
When getting object from storage you receive HTTP Response 
with headers corresponding to the content. 

#### Get Request:
```GET``` 127.0.0.1:8080/:key

#### Set Request and Response:
```POST``` 127.0.0.1:8080/:key
```json
{
    "extension": "png",
    "ttl": 14400,
    "value": "base64 encoded content"
}
```
```json
{
    "success": true
}
```

#### Delete Request and Response:
```DELETE``` 127.0.0.1:8080/:key
```json
{
    "success": true
}
```

### Used Indexes
1. Hash index for ```key```, O(1) access
2. B-tree index for ```expiresAt``` O(logN * M) delete,

where N - number of documents in collection and M is number of documents
in collection which were already expired.


### Deployment
```bash
cp ./mongo/.env.example ./mongo/.env
cp ./storage/.env.example ./storage/.env
docker compose up -d
```

### Other

1. Database store path to files, but all the content stored in file system
2. Files stored locally, but it can be scaled for network file systems
3. ```Set``` request with files must contain Base64 encoded value (all other files will be treated as text/plain)
4. ```Get``` request contains not encoded raw content
5. Type of content in ```Get``` managed by headers inside response
6. For ```Set``` request Content-Type must be ```application/json```
7. Available HTTP response statuses: 200, 400, 404, 500
8. ```Set``` request with key that already exists will override previous record
9. Records with expired ```TTL``` will be deleted only when called methods ```Get```, ```Set```, ```Delete```, but it's possible to make CRON task for this