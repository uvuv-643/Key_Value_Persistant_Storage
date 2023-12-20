# Key-Value Persistant Storage
#### Golang + MongoDB

Golang Microservice with MongoDB. \
Microservice supports operations:
1. ```Set(key string, type string, ttl int, value any)```
2. ```Get(key string)```
3. ```Delete(key string)```

When creating new object in storage you have to specify type: 
```text```, ```photo```, ```audio```, ```video```, etc. \
When getting object from storage you receive HTTP Response 
with headers corresponding to the content. 

#### Set Request and Response:
```json
{
    "key": "hello-world",
    "type": "photo",
    "ttl": 14400,
    "value": "base64 encoded content"
}
```
```json
{
    "success": true
}
```

#### Get Request:
```json
{
    "key": "hello-world"
}
```

#### Delete Request and Response:
```json
{
    "key": "hello-world"
}
```
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