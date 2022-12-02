## kubectl-testkube run test

Starts new test

### Synopsis

Starts new test based on Test Custom Resource name, returns results to console

```
kubectl-testkube run test <testName> [flags]
```

### Options

```
      --args stringArray                           executor binary additional arguments
      --artifact-dir stringArray                   artifact dirs for container executor
      --artifact-storage-class-name string         artifact storage class name for container executor
      --artifact-volume-mount-path string          artifact volume mount path for container executor
      --concurrency int                            concurrency level for multiple test execution (default 10)
      --copy-files stringArray                     file path mappings from host to pod of form source:destination
  -d, --download-artifacts                         downlaod artifacts automatically
      --download-dir string                        download dir (default "artifacts")
      --env stringToString                         envs in a form of name1=val1 passed to executor (default [])
      --execution-label stringToString             execution-label key value pair: --execution-label key1=value1 (default [])
  -h, --help                                       help for test
      --http-proxy string                          http proxy for executor containers
      --https-proxy string                         https proxy for executor containers
      --image string                               execution variable passed to executor
      --iterations int                             how many times to run the test (default 1)
      --job-template string                        job template file path for extensions to job template
  -l, --label strings                              label key value pair: --label key1=value1
  -n, --name string                                execution name, if empty will be autogenerated
      --secret stringToString                      secret envs in a form of secret_key1=secret_name1 passed to executor (default [])
  -s, --secret-variable stringToString             execution secret variable passed to executor (default [])
      --secret-variable-reference stringToString   secret variable references in a form name1=secret_name1=secret_key1 (default [])
  -v, --variable stringToString                    execution variable passed to executor (default [])
      --variables-file string                      variables file path, e.g. postman env file - will be passed to executor if supported
  -f, --watch                                      watch for changes after start
```

### Options inherited from parent commands

```
  -a, --api-uri string     api uri, default value read from config if set (default "http://localhost:8088")
  -c, --client string      client used for connecting to Testkube API one of proxy|direct (default "proxy")
      --namespace string   Kubernetes namespace, default value read from config if set (default "testkube")
      --oauth-enabled      enable oauth (default true)
      --verbose            show additional debug messages
```

### SEE ALSO

* [kubectl-testkube run](kubectl-testkube_run.md)	 - Runs tests or test suites
