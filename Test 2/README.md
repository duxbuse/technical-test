# Technical-test - Test2

## Application

Simple Go webserver that presents a `/version` endpoint detailing the current version of the application as well as its description and the commit associated with this version.

## Pipeline

There are 2 main stages to the pipeline and its all handled using GITHUB ACTIONS. The description of the stages can be found [here](../.github/workflows/test2.yml)

### CI

Here we run the unit tests, perform linting and security analysis. Then if all these things pass we move on to CD.

### CD

Knowing that all the tests and requirements of good code have been passed its now time to bump the version, and deploy this image to [docker hub](https://hub.docker.com/repository/docker/duxbuse/test2) for any users to access. The current pipeline will tag the image under the new version the current commit sha and latest.

## Test Suite

There is currently only the single test which targets the `/version` endpoint to ensure it is returning the right values in the correct format.

## Risks

The gosec github action seems to be broken currently being unable to find the GOPATH, [issue](https://github.com/securego/gosec/issues/259) So it is currently pointing to an empty dir and hence passes. When this issue is resolved ideally this would then point correctly to the source code.

## Versioning

Versioning follows semver as described [here](https://github.com/marketplace/actions/github-tag-with-semantic-versioning)
