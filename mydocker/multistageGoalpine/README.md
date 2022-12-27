
```bash
docker build -t multistagegoweb .
```

```bash
docker container run -p 3333:5000 --name looperA multistagegoweb
```

# Debugging

```bash
docker rm looperA
docker container run -d -it --name looperA multistagegoweb sh -c 'while true; do date; sleep 1; done'
#docker container run -p 3333:5000 -d -it --name looperA multistagegoweb sh -c 'while true; do date; sleep 1; done'
docker exec -it 7e436 sh 
```

# Extra

`CGO_ENABLED=0` is very important when doing a multistage build with a go binary so that your binary doesnt dynamically link to the container's libc

```bash
RUN CGO_ENABLED=0 go build ...
```