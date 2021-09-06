## header-dump

```
docker build -t header-dump:latest .
docker run --name header-dump --rm -ti -p 8080:8080 header-dump:latest

curl --header "Host: foobar.com" localhost:8080/
```
