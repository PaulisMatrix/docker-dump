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
    ```
    ❯ kubectl delete -f flask-deployment.yaml
        deployment.apps "flask-demo" deleted
        service "flask-entrypoint" deleted

9.  Adding ingress to the app. 

    Suppose this is the IP resolution for our website `https://flask-demo.com/` (For example purposes only.)
    ```
        Non-authoritative answer:
        Name:	flask-demo.com
        Address: 130.211.24.50
    ```
    Now how is the request actually routed? 

    a.  When you hit `https://flask-demo.com/`, whatever DNS provider you are using, resolves this hostname to its public IP address
        which is `130.211.24.50` in this case.<br>
    b.  The DNS server then routes this request to its appropriate destination which is the ingress controller service inside your
        kubernetes cluster.<br>
    c.  Kubernetes ensures that requests from this IP address is passed to your cluster by setting up a network route between the
        public IP address and the ingress controller service.<br>
    d.  This is typically done by setting up a load balancer or reverse proxy(like nginx) that listens on the public IP address and
        forwards the incoming request to the ingress controller service inside the k8s cluster.
        (After setting it up, you link this proxying under ingress class in your ingress resource yaml file). <br>
    e.  When the ingress controller receives a request from the public IP address, it uses the rules defined in the ingress resource
        yaml file to determine to which service backend and endpoint should handle the request.<br>
    f.  The request is then forwarded to the appropriate service and endpoint within the cluster, which is basically your flask demo app.

