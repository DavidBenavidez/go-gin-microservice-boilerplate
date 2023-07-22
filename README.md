# Golang Microservice Gin Boilerplate
Sample Golang Gin Boilerplate 
For easy deployment as a microservice in a kubernetes cluster

# 
<pre>
├── api 
├── configs
├── internal # service/business logic
    ├── rest # routing abstraction / controller
    ├── service # service/business logic
    └── clients # external services
├── pkg
└── test
    └── mocks
</pre>

# Requirements
1. GO 1.20
1. Docker
1. kubectl

# Development
Setup local environment
(optional) Setup database connection
```
  go run .
```

# Unit Tests
### Update Mocks
```
 make mocks
```

### Run units test
```
  make tests
```
or 
```
  make tests_ui
```
# Deployment
### Deploy docker / set k8s image
```
  make deploy
```

# Build SDK
```
  make client
```

# Build Kubernetes resource templates
```
  make build_kustomize
```