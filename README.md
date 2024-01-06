# pod_info_go
A simple golang app for displaying information about a Kubernetes pod hosting the app

<br/>

### To Run locally
```shell
cd app
go run main.go
```
<br/>

### To build the Docker image
```shell
docker build . -t pod_info_go
```
<br/>

### To run the Docker image
```shell 
docker run -p 8080:8080 pod_info_go
```
