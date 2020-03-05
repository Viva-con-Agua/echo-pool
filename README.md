# echo-pool

## responses models

```
pool.InternelServerError() 

{
  "message": "Internel server error, please check logs"
}
```

```
pool.Conflict() 

{
  "message": "Models already exists"
}
```

```
pool.Created() 

{
  "message": "Successful created"
}
```

```
pool.Unauthorized() 

{
  "message": "Not authenticated"
}
```

```
pool.NoContent() 

{
  "message": "Not found",
  "uuid": uuid
}
```

```
pool.Updated() 

{
  "message": "Successful updated"
}
```

```
pool.Deleted() 

{
  "message": "Successful deleted",
  "uuid": uuid
}
```



