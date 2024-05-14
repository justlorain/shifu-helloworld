# shifu-helloworld

## Steps

### Step 1: Install Shifu

- **Commands**

```shell
docker ps
curl -sfL https://raw.githubusercontent.com/Edgenesis/shifu/main/test/scripts/shifu-demo-install.sh | sudo sh -
sudo kubectl get pods -A 
```

- **Proof of Work**

![1-1](./images/step-1-1.png)
![1-2](./images/step-1-2.png)
![1-3](./images/step-1-3.png)

### Step 2: Prepare shifudemos

- **Commands**

```shell
cd shifudemos
sudo kubectl run --image=nginx:1.21 nginx
sudo kubectl get pods -A | grep nginx
```

- **Proof of Work**

![2](./images/step-2.png)

### Step 3: Build and Push Image

- **Commands**

```shell
docker login
docker build -t shifu-helloworld:latest .
docker tag shifu-helloworld:latest justlorain/shifu-helloworld:latest
docker push justlorain/shifu-helloworld:latest
```
- **Proof of Work**

![3-1](./images/step-3-1.png)
![3-2](./images/step-3-2.png)

### Step 4: Apply shifu-helloworld

- **Commands**

```shell
sudo kubectl apply -f run_dir/shifu/demo_device/edgedevice-plate-reader
sudo kubectl get pods -A | grep plate
```

- **Proof of Work**

![4-1](./images/step-4-1.png)
![4-2](./images/step-4-2.png)
![4-3](./images/step-4-3.png)

### Step 5: Check Log

- **Commands** 

```shell
sudo kubectl get pods -A | grep shifu-helloworld
sudo kubectl logs shifu-helloworld-5d9cbb45fc-mwmks
```

- **Proof of Work**

![5-1](./images/step-5-1.png)
![5-2](./images/step-5-2.png)

## License

shifu-helloworld is distributed under the [Apache License 2.0](./LICENSE).
