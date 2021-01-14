# go-crud

# 1: Create MySql in local
```bash
  $ . scripts/docker-up.sh
```

# 2: Start the Server
```bash
  $ . script/start.sh
  # Successful Output: 
  # DB selected successfully..
  # Drop tasks created successfully..
  # Drop users created successfully..
  # Table tasks created successfully..
  # Table users created successfully..
  # Listening On http://localhost:8090 ...
```

# 3: API
## All tasks API need login (i.e. Auth header) to access
##
| API   |      Param Type      |  Params |
|----------|:-------------:|------|
| GET /signup | QueryString | username:string password:string |
| GET /signin | QueryString | username:string password:string |
| POST /tasks | Form Data | JSON { name, due, completion, status } |
| GET /tasks/:id | QueryString | id:int |
| PUT /tasks/:id | QueryString AND Form Data | id:int, JSON { name, due, completion, status }|
| DEL /tasks/:id | QueryString | id:int |
    