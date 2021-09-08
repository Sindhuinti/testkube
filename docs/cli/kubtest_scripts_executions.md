## kubtest scripts executions

List scripts executions

### Synopsis

Getting list of execution for given script name or recent executions if there is no script name passed

```
kubtest scripts executions [flags]
```

### Options

```
  -h, --help   help for executions
```

### Options inherited from parent commands

```
  -c, --client string        Client used for connecting to kubtest API one of proxy|direct (default "proxy")
      --go-template string   in case of choosing output==go pass golang template (default "{{ . | printf \"%+v\"  }}")
  -s, --namespace string     kubernetes namespace (default "default")
  -o, --output string        output typoe one of raw|json|go  (default "raw")
  -v, --verbose              should I show additional debug messages
```

### SEE ALSO

* [kubtest scripts](kubtest_scripts.md)	 - Scripts management commands
