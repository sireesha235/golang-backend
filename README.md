# golang-backend

## Local Development:
Command to Build a docker image 
```shell
docker build -t golang-backend .
```
```shell
docker run -d -p 9090:9090 <IMAGEID>

```
* api endpoints:
  * `/health`
  * `/servertime`

`/health` endpoint returns the status of service.

`/servertime` returns the Server Epoch time and Time Zone

Once the container started app can be accessed at below endpoints in local
```html
http://localhost:9090/health
```

```html
http://localhost:9090/servertime
```

