# Blueprint/Boilerplate For Golang Projects

[![Build Status](https://travis-ci.com/MartinHeinz/go-project-blueprint.svg?branch=master)](https://travis-ci.com/MartinHeinz/go-project-blueprint)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=MartinHeinz_go-project-blueprint&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=MartinHeinz_go-project-blueprint)
[![Test Coverage](https://api.codeclimate.com/v1/badges/ec7ebefe63609984cb5c/test_coverage)](https://codeclimate.com/github/MartinHeinz/go-project-blueprint/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/trung85-go/go-project-blueprint)](https://goreportcard.com/report/github.com/trung85-go/go-project-blueprint)

## Blog Posts - More Information About This Repo

You can find more information about this project/repository and how to use it in following blog posts:

- [Ultimate Setup for Your Next Golang Project](https://towardsdatascience.com/ultimate-setup-for-your-next-golang-project-1cc989ad2a96)
- [Setting up GitHub Package Registry with Docker and Golang](https://towardsdatascience.com/setting-up-github-package-registry-with-docker-and-golang-7a75a2533139)
- [Building RESTful APIs in Golang](https://towardsdatascience.com/building-restful-apis-in-golang-e3fe6e3f8f95)
- [Setting Up Swagger Docs for Golang API](https://towardsdatascience.com/setting-up-swagger-docs-for-golang-api-8d0442263641)

### Setting Up
- Replace All Occurrences of `trung85-go/go-project-blueprint` with your username repository name
- Replace All Occurrences of `blueprint` with your desired image name


### Adding New Libraries/Dependencies
```bash
go mod vendor
```

### Using GitHub Registry

Create and Push:

```bash
docker login docker.pkg.github.com -u <USERNAME> -p <GITHUB_TOKEN>
docker build -t  docker.pkg.github.com/trung85-go/go-project-blueprint/blueprint:latest .
# make container
docker push docker.pkg.github.com/trung85-go/go-project-blueprint/blueprint:latest
# make push
```

List all docker packages on github: https://github.com/trung85-go/go-project-blueprint/packages

Pull and Run:

```bash
docker pull docker.pkg.github.com/trung85-go/go-project-blueprint/blueprint:latest
docker run -p 1234:1234 \
           -e BLUEPRINT_DSN="postgres://<USER>:<PASSWORD>@<DOCKER_SERVICE/URL>:<PORT>/<DB>?sslmode=disable" \
           -e BLUEPRINT_API_KEY="api_key" \
           docker.pkg.github.com/trung85-go/go-project-blueprint/blueprint:latest
```

### Setup new SonarCloud Project

- On _SonarCloud_:
    - Click _Plus_ Sign in Upper Right Corner
    - _Analyze New Project_
    - Click _GitHub app configuration_ link
    - Configure SonarCloud
    - Select Repository and Save
    - Go Back to Analyze Project
    - Tick Newly Added Repository
    - Click Set Up
    - Click Configure with Travis
    - Copy the Command to Encrypt the Travis Token
    - Run `travis encrypt --com <TOKEN_YOU_COPPIED>` (generate token in SonarCloud)
    - Populate the `secure` Field in `.travis.yml` with outputted string
    - Follow steps to populate your `sonar-project.properties`
    - Push
- On Travis CI:
    - Set `DOCKER_USERNAME`.  (organization github trung85-go)
    - Set `DOCKER_PASSWORD` to Your GitHub Registry Token (develop github token)

### Setup CodeClimate
- Go to <https://codeclimate.com/github/repos/new>
- Add Repository
- Go to Test Coverage Tab
- Copy Test Reporter ID
- Go to Travis and Open Settings for Your Repository
- Add Environment Variable: name: `CC_TEST_REPORTER_ID`, value: _Copied from CodeClimate_ (https://codeclimate.com/repos/5f114fcb20e8de018b00b776/settings/test_reporter)

## Swagger

- Application uses [gin-swagger](https://github.com/swaggo/gin-swagger).
- To generate/update docs use `swag init` (from `/go-project-blueprint/cmd/blueprint`)
- You can find generated docs in `docs` package

To view docs, navigate to <http://localhost:1234/swagger/index.html> or to <http://localhost:1234/swagger/doc.json> for raw _JSON_
