# Helm- the package manager for kubernetes.

1.  Helm features:<br>
    a.  As a package manager. Install charts which are basically a collection of yaml files/manifests which you usually define.<br>
    b.  A templating engine. Define common template files for your different projects/envs and _fill_ in by only having a single values yaml file thereby avoiding duplication.
    c.  Easier to upgrade/rollback application versions through simple commands so basically release managment as well.

2.  [Helm commands](https://helm.sh/docs/helm/)

3.  Get started by firing `helm create foo` which will create a chart directory along with the common files and directories used in a   chart.<br>
    The directory structure will something like this:<br>
    ```
    foo/
    ├── .helmignore   # Contains patterns to ignore when packaging Helm charts.
    ├── Chart.yaml    # Information about your chart
    ├── values.yaml   # The default values for your templates
    ├── charts/       # Charts that this chart depends on
    └── templates/    # The template files
        └── tests/    # The test files

4.  
