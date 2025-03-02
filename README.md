## Localde çalıştırmak için
postgresql veritabanı, golang 1.23 sürümü ve python yüklü olması gerekmektedir.
.env dosyası açılarak (.env.example örnek alınabilir), veritabanı bilgileri doğru şekilde girilir.
```bash
go mod tidy
go run main.go &
python3 -m http.server 8000 --directory /public
```
localhost:8000 adresinden uygulamaya erişilebilir.


## Docker container üzerinde çalıştırmak için
veritabanı ve uygulama için 2 ayrı container kaldırmaktadır.
```bash
docker-compose up --build -d
```
localhost:8000 adresinden uygulamaya erişilebilir.

## Kubernetes cluster üzerinde çalıştırmak için
Docker'ın yanında kind ve kubectl kurulu olmalıdır.

```bash
# cluster oluşturulması
kind create cluster --name todo-cluster --config=kind-config.yaml

# Ingress controller için auto-fix, bu olmadığı zaman dışarıdan erişim sağlanamıyor. internette çözüm olarak bu gösteriliyor.
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
kubectl patch deployment ingress-nginx-controller -n ingress-nginx --type=json -p '[{"op": "remove", "path": "/spec/template/spec/nodeSelector"}]'
kubectl taint nodes --all node-role.kubernetes.io/control-plane-
kubectl rollout restart deployment ingress-nginx-controller -n ingress-nginx
kubectl wait --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller --timeout=120s

# ingress podlarının kontrolü 
kubectl get pods -n ingress-nginx

# deployment
kubectl apply -f k8s/postgres.yaml
kubectl apply -f k8s/todoapp.yaml
kubectl apply -f k8s/ingress.yaml

# kontrol
kubectl get all -A

# linux ise /etc/hosts, windows ise C:\Windows\System32\drivers\etc\hosts dosyasına "127.0.0.1 todo.local" satırının eklenmesi

# api kontrolü
curl http://todo.local/tasks
```
todo.local adresinden uygulamaya erişilebilir