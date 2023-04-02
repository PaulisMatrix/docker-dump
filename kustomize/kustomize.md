## Kustomize for kubernetes manifests.

1.  [kustomize](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/) lets you customize raw, template-free YAML files for multiple purposes, leaving the original YAML untouched and usable as is.

2.  Good thing is kustomize is native to kubectl. 
    ```
    ❯ kubectl version --short --client
        Flag --short has been deprecated, and will be removed in the future. The --short output will become the default.
        Client Version: v1.26.0
        Kustomize Version: v4.5.7

3.  Kustomize is based around the concept of a base and overlays. <br><br>
    a.  A **base** is a directory with a kustomization.yaml, which contains a set of resources and associated customization. A base could be either a local directory or a directory from a remote repo, as long as a `kustomization.yaml` is present inside.
    A base has no knowledge of an overlay and can be used in multiple overlays<br><br>
    b.  An **overlay** is a directory with a kustomization.yaml that refers to other kustomization directories as its bases.
    An overlay may have multiple bases and it composes all resources from bases and may also have customization on top of them.


4.  So basically **overlays** work on top of a common **base** manifests which helps in reusability.

5.  Most suitable for scenario wherein you have different variants like `dev`, `stage` and `production` envs with having its own
    **overlays** but referring to a common **base**.

6.  How to use it?<br>
    <br>a.  cd into any directory having `kustomization.yaml` file.
    <br>b.  In this case, I would cd into `overlays/dev` directory.
    <br>c.  Run the command: `kubectl apply -k .`
    <br>d.  You can see two deployments for your dev environment running.
    ```
    ❯ kubectl get deployment
        NAME         READY   UP-TO-DATE   AVAILABLE   AGE
        flask-demo   2/2     2            2           14s

7.  Delete or teardown your app: `kubectl delete -k .` in the same `overlays/dev` directory.

