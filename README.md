# go-hello-scratch
A sample project for creating a tiny Docker image using Go.

## Build
Open Terminal and run `make docker`. It will build a container about 6MB.

## Run
Run the container `docker run -it -p 3000:3000 hello-scratch` and open brower, visit http://localhost:3000/ (**Note**: for Mac user, make sure you are running `boot2docker ssh -vnNTL 3000:localhost:3000` in another terminal window.)

## References
* http://blog.codeship.com/building-minimal-docker-containers-for-go-applications/?utm_source=golangweekly&utm_medium=email
* https://github.com/boot2docker/boot2docker/blob/master/doc/WORKAROUNDS.md

