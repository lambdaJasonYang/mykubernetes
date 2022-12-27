
```bash
docker build -t goaplineweb .
```

```bash
docker container run -p 3333:5000 --name looperA goaplineweb
```

# Debugging

```bash
docker rm looperA
docker container run -d -it --name looperA goaplineweb sh -c 'while true; do date; sleep 1; done'
# docker container run -p 3333:5000 -d -it --name looperA goaplineweb sh -c 'while true; do date; sleep 1; done'
docker exec -it 7e436 sh 
```