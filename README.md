# kubernetes101
This repo contains a sample app that have some three components and can be deployed on kubernetes natively using some kubernetes resources.

# run-mysql
To run mysql you wil have to run below commands, they are going to create a PV, PVC and deployment that will actually host the mysql, `cd` into `run-mysql` and run 
```
kubectl create -f mysql-pv.yaml
kubectl create -f mysql-pv-claim.yaml -n <ns-name>
kubectl create -f mysql-deployment.yaml -n <ns-name>
kubectl create -f mysql-service.yaml -n <ns-name>
```

# restapi
This dir contains the Dockerfile for the app that talks to the mysql database(`testlocal`)'s `books` table and gets all the  details. This application requires 
some env variables like APIPATH, DBHOST and DBPASS and these env variables will be passed while creating deployment manifest of this application, that can be 
found inside restapi-k8s dir.

To build the docker images that will be used by the deployment manifest, cd into restapi dir and run 
```
docker build -t restapi:1.0 .
```


# restapi-k8s
This dir contains all manifests that are required deploy the restapi app on the kubernetes cluster. If you want to deploy all these apps on separate namespace 
you can create another namespace and then deploy all these manifest files in that namespace. To create a namespace you can run 
```
kubectl create ns <ns-name>
```
Once you have the namespace created you will have to create the deployment that will run the image that you have already built to create the deployment and expose it through a service and ingress run below command
```
kubectl create -f restapi-deployment.yaml -n <ns-name>
kubectl create -f restapi-service.yaml -n <ns-name>
kubectl create -f restapi-ingress.yaml -n <ns-name>
```
The above commands will deploy the restapi application on kubernetes cluster and expose them to the outside word using service and ingress.

# restapi-ui
To to create the build of restapi ui code `cd` to the `restapi-ui` dir and run below command 
```
docker build -t ui:1.0 .
```
above command will create the docker image of the UI component and that can be used to deploy the UI component on the kubernetes.

# restapi-ui-k8s
To deploy the UI image that you have already created in the previous step you will have to run below command in order to create deployment for UI, service and ingress
```
kubectl create -f restapi-ui-deployment.yaml -n <ns-name>
kubectl create -f restapi-ui-service.yaml -n <ns-name>
kubectl create -f restapi-ui-ingress.yaml -n <ns-name>
```

# Note
To actually get to know why are we creating deployment, services and ingress you can follow a post that is [here](https://medium.com/@viveksinghggits/hello-world-of-kubernetes-part-1-d1153fc2fc37).
