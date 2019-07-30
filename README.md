# go-server-complex

Include base components a backend server need.

Features:
* HTTP Service Framework: use [gin](https://github.com/gin-gonic/gin).
* Configure function: use [toml](https://github.com/BurntSushi/toml).
* Log function: based on [logrus](https://github.com/sirupsen/logrus).
* Include [UUID](https://github.com/google/uuid) for each request in log message.



## Installation:
### 404 Packages:
* golang.org/x/sys/unix  
mirror: https://github.com/golang/sys

### Build
```shell
git clone git@github.com:kmiku7/go-server-complex.git
cd go-server-complex
export GOPATH=`pwd`
govendor sync

govendor install server
./bin/server -config-path ./config/demo.toml &

curl http://127.0.0.1:8066/ping/
```
