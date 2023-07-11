## Here lies all the usefull links for docker or kubernetes, all devops stuff, you name it.

* DOCKER:

    [Play with](https://labs.play-with-docker.com/) Docker :') Its fun, I tell you.

1. Best practices to write a Dockerfile: 

    https://kapeli.com/cheat_sheets/Dockerfile.docset/Contents/Resources/Documents/index

    https://docs.docker.com/develop/develop-images/dockerfile_best-practices/

    https://dockerlabs.collabnix.com/docker/cheatsheet/

2. Docker commands:

    https://gist.github.com/bradtraversy/89fad226dc058a41b596d586022a9bd3

3. Compose best practice - How to write compose files:

    https://github.com/compose-spec/compose-spec/blob/master/spec.md

4. Multistage builds:

    https://blog.alexellis.io/mutli-stage-docker-builds/

5. Docker Caching - How docker layers work and how can you take maximum advantage of it:

    https://medium.com/swlh/docker-caching-introduction-to-docker-layers-84f20c48060a

6. In general docker articles:

    https://pythonspeed.com/docker/

7. Dive into your docker layers, some wizardry here:

    https://github.com/wagoodman/dive

8. Tools for analysing your containers for potential vulnerabilities, fixes, improvements, etc :

    1.  https://snyk.io/docker 

    2.  https://github.com/snyk/cli

    3.  https://github.com/anchore/grype

9. Dokter - Analyzer for Dockerfiles in hopes to build faster, more secure, smaller images.
            Snyk and most container scanning tools happen after the build, and usually in the registry (unless you locally tar), that could potentially expose credentials or an vulnerable image. Dokter aims to prevent building an insecure image (saving CI minutes) and preventing leaking credentials by not building if its known to produce a faulty image.
    
    https://gitlab.com/gitlab-org/incubation-engineering/ai-assist/dokter


10. Podman - Alternative for docker by RedHat: https://www.redhat.com/en/topics/containers/what-is-podman 

11. Kraken - P2P docker registry, an alternative to Docker Hub by Uber: 

    1.  https://github.com/uber/kraken 
    2.  https://uber-kraken.readthedocs.io/en/latest/ 

12. Colima - Container runtime, an alternative for Docker Desktop: 

    1.  https://github.com/abiosoft/colima
    2.  https://medium.com/javarevisited/free-docker-desktop-alternative-for-mac-c3845d8a2345
