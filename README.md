# NetDemo
## Docker & Kubernetes Demo

```
docker build -t netdemo . -f ./build/Dockerfile
docker run -d -p 80:80 netdemo
```