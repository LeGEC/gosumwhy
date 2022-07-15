# gosumwhy
search dependencies between go modules from the output of go mod graph

### Installation

```
go install github.com/LeGEC/gosumwhy@latest
```

### Usage

from within a go module (it internally runs `go mod graph` and reads the graph from the output of that command) :

```
gosumwhy list                      # list all modules present in the graph
gosumwhy path module/name@version  # print a dependency path to that module@version
gosumwhy -allv path module/name    # print a dependency path for each version of that module
```

See `gosumwhy -h` for more details on ways to provide a dependency graph:

```
# from stdin:
go mod graph | gosumwhy ...
# from a file (e.g: generated using 'go mod graph > graph.txt'):
gosumwhy -f graph.txt ...
# running 'go mod graph' in a module located somewhere else on disk:
gosumwhy -modpath path/to/module ...
```
