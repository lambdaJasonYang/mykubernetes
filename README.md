# postgrespod.yaml

```bash
##imperative style
kubectl run postgrespod --image=postgres --restart=Never --port=5432 --env="POSTGRES_PASSWORD=root" --labels="app=mypostgres,env=prod"
```

```bash
kubectl apply -f postgrespod.yaml

## container port 5432 ==> 127.0.0.1:8098
kubectl port-forward postgrespod 8098:5432

## container port 5432 ==> postgres on 0.0.0.0:8098
kubectl port-forward postgrespod 8098:5432 --address='0.0.0.0'

## demo
psql -h 127.0.0.1 -p 8098 -U postgres
#psql -h 192.168.1.244 -p 8098 -U postgres
```