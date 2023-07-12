1. Kubernetes cheatsheet: https://kubernetes.io/docs/reference/kubectl/cheatsheet/

2. Kubernetes tools for development:

   a. devspace: https://devspace.sh/

   b. telepresence: https://www.telepresence.io/

   c. squash: https://squash.solo.io/overview/

   d. okteto: https://www.okteto.com/

3. Kuberenetes IDE: https://k8slens.dev/

4. Tool to manage contexts and namspaces: https://github.com/ahmetb/kubectx

5. How deployment, service and ingress in a normal kubernetes setup are related to eachother:<br>

   a. https://dwdraju.medium.com/how-deployment-service-ingress-are-related-in-their-manifest-a2e553cf0ffb<br>
   b. Basically you have your deployment pods, on top of that you define a service which acts like a common entry
      point for all your pods. Then define ingress on top service to define your incoming traffic.

6. Help discover deprecated api versions for k8s:
 
   a. https://github.com/FairwindsOps/pluto<br>
   b. https://github.com/doitintl/kube-no-trouble

7. K8s cluster autoscaler: https://github.com/aws/karpenter (only aws exclusive)

8. etcd which is the k8s distributed kv store maintaining all the relevant state of k8s resources uses Raft algorithm for
   managing this data across the cluster.

   Raft algorithm: <br>
         1. https://raft.github.io/<br>
         2. https://kubernetes.io/blog/2019/08/30/announcing-etcd-3-4/<br>
         3. https://developer.hashicorp.com/consul/docs/architecture/consensus<br>
         4. Raft demonstration: https://thesecretlivesofdata.com/raft/

9. Walks through kubernetes internals: https://book.kubebuilder.io/introduction.html

10. Awesome k8s: https://github.com/ramitsurana/awesome-kubernetes

11. K8s networking guide: https://www.tkng.io/

12. Kubernetes talks:

    a. Overview of k8s: https://www.youtube.com/watch?v=WwBdNXt6wO4 
    
    b. Containers from scratch:<br>
          1. https://youtu.be/_TsSmSu57Zo<br>
          2. https://youtu.be/8fi7uSYlOdc
    
    c. Kelsey Hightower talks:<br>
           1. Reflections: https://youtu.be/jiaLsxjBeOQ<br>
           2. Schedular from scratch: https://youtu.be/IYcL0Un1io0

    d. Kubernetes Deconstructed: https://youtu.be/90kZRyPcRZw
  
    e. Envoy Internals: https://youtu.be/gQF23Vw0keg
   
    f. CoreDNS(Which handles DNS resolutions and all for k8s) : https://youtu.be/qRiLmLACYSY
    
    g. Kubernetes design principles: https://youtu.be/ZuIQurh_kDk

    h. Kubernetes networking: https://youtu.be/0Omvgd7Hg1I
    
    i. About Kubernetes Storage:<br>
          1. https://youtu.be/uSxlgK1bCuA<br>
          2. https://youtu.be/hqE5c5pyfrk
