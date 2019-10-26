# kubernetes101
This repo contains a sample app that have some three components and can be deployed on kubernetes natively using some kubernetes resources. This repo is a part of a tutorial that I have written on [medium](https://medium.com/@viveksinghggits/hello-world-of-kubernetes-part-1-d1153fc2fc37).

Please clone the repository to go through the commands that are given below.

# Create a secret that will be used as password in the myql deployment 
```
kubectl create secret generic mysql-password --from-literal password='vivek$ingh'
```

# run-mysql
To run mysql you wil have to run below commands, they are going to create a PV, PVC and deployment that will actually host the mysql, `cd` into `run-mysql` and run 
```
# Create the persistent volume resource
kubectl create -f mysql-pv.yaml
# Create the persistent volumeclaim resource 
kubectl create -f mysql-pv-claim.yaml -n <ns-name>
# Create the deployment that will actually run the mysql pod
kubectl create -f mysql-deployment.yaml -n <ns-name>
# Create the mysql service that will expose the mysql deployment, so that it can be access by other components.
kubectl create -f mysql-service.yaml -n <ns-name>
```
Once you have the PersistentVolume, and the claim created you can just list the PersistentVolume and Claim to make sure the claim is satisfied by the
PersistentVolume that we have already created. Below command can be used to list the persistentvolumes and claims.
```
kubectl get persistentvolume -n <ns-name>
kubectl get persistentvolumeclaim -n <ns-name>
```

# restapi
This dir contains the Dockerfile for the app that talks to the mysql database(`testlocal`)'s `books` table and gets all the  details. This application requires 
some env variables like APIPATH, DBHOST and DBPASS and these env variables will be passed while creating deployment manifest of this application, that can be 
found inside restapi-k8s dir.

To build the docker images that will be used by the deployment manifest, cd into restapi dir and run 
```
docker build -t restapi:1.0 .
```
# Create the configmap that will be used by the restapi app to talk to talk to the mysql database
Once we have the docker image for restapi app crated we will have to create configmaps so that we will not have to hard code the config values in the manifest itself, we can read these configuration values from the configmaps itself.
```
kubectl create configmap restapi-config --fron-literal APIPATH=/api/v1/books --from-literal DBHOST=mysql:3306
```

# restapi-k8s
This dir contains all manifests that are required deploy the restapi app on the kubernetes cluster. If you want to deploy all these apps on separate namespace 
you can create another namespace and then deploy all these manifest files in that namespace. To create a namespace you can run 
```
kubectl create ns <ns-name>
```
Note:
In the first post we have discussed how a pod  communicated to the service that is in the different namespace, so make sure to change [this line](https://github.com/viveksinghggits/kubernetes101/blob/master/restapi-k8s/restapi-deployment.yaml#L25) with the correct namespace where your mysql database is deployed. So the the restapi application can talk to the mysql service.

Once you have the namespace created you will have to create the deployment that will run the image that you have already built to create the deployment and expose it through a service and ingress run below command
```
kubectl create -f restapi-deployment.yaml -n <ns-name>
kubectl create -f restapi-service.yaml -n <ns-name>
kubectl create -f restapi-ingress.yaml -n <ns-name>
```
The above commands will deploy the restapi application on kubernetes cluster and expose them to the outside word using service and ingress.
Please take a note that we are not just creating the service but we are actually creating an ingress as well to expose that service outside of the cluster. This is the same ingress that the UI component will be using to talk to the restapi.

To make the ingress work in minikube you might have to run below command 
```
minikube addons enable ingress
```
In this branch we are not going to deploy the UI component because the main motive was to get to know how secrets and configmaps can be used to store env variables that we are going to provide to the applications. 


# Note
To actually get to know why are we creating deployment, services and ingress you can follow a post that is [here](https://medium.com/@viveksinghggits/hello-world-of-kubernetes-part-1-d1153fc2fc37).
