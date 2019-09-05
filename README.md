# TODO: Update the README

# kubernetes101
This repo contains a sample app that have some three components and can be deployed on kubernetes natively using some kubernetes resources.

# run-mysql


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

# restapi-ui

# restapi-ui-k8s

