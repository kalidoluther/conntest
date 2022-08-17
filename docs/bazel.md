# Bazel

``` bash
bazel run servers/client
bazel run servers/server
bazel run servers/sim7600g

gazelle update -go_prefix kalidoluther.com/conntest

gazelle -go_prefix kalidoluther.com/conntest

bazel run //...


bazel test //servers/server:server_test
bazel test //servers/client:client_test
```