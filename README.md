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

## auth

### connect redis

```
	store := pool.RedisSession()
	
  //create echo server
	e := echo.New()
	e.Use(store)
```

### use nats 

```
	apiV1 := e.Group("/api/v1")
	apiV1.Use(resp.SessionAuth)
```
### nats

#### import
```
  import (
    "github.com/Viva-con-Agua/echo-pool/nats"

  )
```
#### initial
```
	nats.NatsConnect()
```
#### use

https://github.com/nats-io/nats.go#encoded-connections

### responses models

#### import 
```
import (
	"github.com/Viva-con-Agua/echo-pool/resp"
)
```

#### use
```
resp.InternelServerError() 

{
  "message": "Internel server error, please check logs"
}
```

```
resp.Conflict() 

{
  "message": "Models already exists"
}
```

```
resp.Created() 

{
  "message": "Successful created"
}
```

```
resp.Unauthorized() 

{
  "message": "Not authenticated"
}
```

```
resp.NoContent() 

{
  "message": "Not found",
  "uuid": uuid
}
```

```
resp.Updated() 

{
  "message": "Successful updated"
}
```

```
resp.Deleted() 

{
  "message": "Successful deleted",
  "uuid": uuid
}
```



