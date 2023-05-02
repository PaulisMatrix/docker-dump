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
