## Create an image and publish to docker hub

1. Dockerfile 

2. podman build . -f Dockerfile -t docker.io/jpalaparthi/user-service1:v1

3. podman login docker.io

4. podman push docker.io/jpalaparthi/user-service1:v1

5. podman rmi docker.io/jpalaparthi/user-service1:v1

6. podman pull docker.io/jpalaparthi/user-service1:v1

## Create containers and run 

1. create podman network 

2. create postgres, adminer and also app container on the same network.

3. If using nginx , then create nginx container in the same network.

4. User similar config file of nginx that is used for the demo

5. Test the application