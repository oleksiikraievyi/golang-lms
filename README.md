# golang-lms test task

# Prerequisites
* docker
* golang 1.22
* linux + bash

# 1. Run build
```
make build
```

# 2. Run docker build
```
make docker-build
```

# 3. Run tests
```
make tests
```

# 4. Run server
```
make run
```

or 

```
docker run -p {internal_port}:{external_port} lms:0.0.1
```

Access swagger endpoint **/swagger/index.html** for api docs 

# 4. Clean up
```
make clean
```

