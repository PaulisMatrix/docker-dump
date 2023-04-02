## Running the flask app from flask-docker inside a kubernetes pod

1.  Build the flask image and tag it:   `docker build -t flask-docker .`

2.  Define kubernetes deployment and service manifests under `flask-deployment.yaml` with whatever spec you want. Make sure app lables are mapped correctly.

3.  Apply the manifests:
    ```
    ❯ kubectl apply -f flask-deployment.yaml 
        deployment.apps/flask-demo created
        service/flask-entrypoint created

4.  Check your deployment and service status:
    ```
    ❯ kubectl get deployments
        NAME         READY   UP-TO-DATE   AVAILABLE   AGE
        flask-demo   1/1     1            1           4m20s
    
    ❯ kubectl get services
        NAME               TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
        flask-entrypoint   NodePort    10.96.91.42   <none>        3000:30001/TCP   4m42s

5.  Check if your pods are running fine or not:
    ```
    ❯ kubectl get pods
        NAME                         READY   STATUS    RESTARTS   AGE
        flask-demo-97bb68cdf-kjjgn   1/1     Running   0          5m49s
    
    ❯ kubectl logs flask-demo-97bb68cdf-kjjgn
        * Serving Flask app 'hello-world' (lazy loading)
        * Environment: production

6.  Go inside the pod: 
    ```
    ❯ kubectl exec -it flask-demo-97bb68cdf-kjjgn /bin/bash
        root@flask-demo-97bb68cdf-kjjgn:/usr/src# ls
        Dockerfile  app  docker-compose.yml  flask-docker.md  python.Dockerfile  requirements.txt

7.  Navigate to `localhost:30001/hello`. 30001 is the nodeport in the service manifest we have defined for our flask app pods.

8.  Teardown your application: 
    ```❯ kubectl delete -f flask-deployment.yaml
        deployment.apps "flask-demo" deleted
        service "flask-entrypoint" deleted