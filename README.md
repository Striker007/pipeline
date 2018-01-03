# golang tdd jenkins pipeline
https://jenkins.io/doc/book/installing/#docker

#### additional steps jenkins\docker:
```shell
docker ps
docker exec -ti ..CONTAINER ID\NAMES.. bash
apk add make go gcc g++
```
#### build \ run \ test - golang with Makefile:
```shell
make run \ make
make test
```
