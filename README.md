# echo-pool

## install

```
go get github.com/Viva-con-Agua/echo-pool/...       
```

## use 

```
import (
	"github.com/Viva-con-Agua/echo-pool/pool"
)
```

## redis 

### connect redis

```
	store := pool.RedisSession("172.2.150.1:6379")
	
  //create echo server
	e := echo.New()
	e.Use(store)
```

### use session in api

```
	apiV1 := e.Group("/api/v1")
	apiV1.Use(pool.SessionAuth)
```


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



