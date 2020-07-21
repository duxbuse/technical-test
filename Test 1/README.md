# Technical-test - Test1

## Issues with sample dockerfile

By adding all the files in before the `apk` update it meant that every time the code changes you will need to re download and install all the packages. So changing the order will mean the code can change but all the installed dependencies will remain cached.

Use of the `ADD` command, `COPY` is now preferred.

All the build dependencies are left in the container. A better solution is to have only whats required in the final output image. Hence extracting the binary output and putting that into a scratch container, you seriously reduce the size of the image as well as preventing your attack surface.

## Issue with sample go code

https://github.com/xUnholy/technical-tests/blob/a206ed5537465a9eb75d9bc9fc7509ee0d5aab3f/main.go#L23

This line sets the web server to listen for incomming requests, however inside the docker container this will mean that it will never be able to service request from outside the container due to the namespaces. This address needs to be changed to `0.0.0.0:8000` to allow traffic from the docker network bridge to still be correctly handled.
