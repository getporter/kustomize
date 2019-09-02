# Building the Porter-Kustomize Mixin

## Running Make

!!! note
    Please install the `porter` binaries first. Instructions for how to do so on different platforms can be found in
    the Porter.sh [documentation](https://porter.sh/install/).
    
In order to build the `porter-kustomize` mixin the following should be run from the command line in the root of the
git repository: -

```bash 
make xbuild-all build install
```

which triggers `make` to build a porter mixin for each supported client platform as the first step. Secondly this
command builds a `porter-runtime` binary which is deployed into the Docker `innvocationImage` that is run by the main
`porter` client executable. Thirdly, the `make` installs the newly build binaries into that of the `porter` client.

For example if the user is creating a CNAB bundle on their Mac then the porter client will be for `Darwin` but the
`innvocationImage` is based on a Linux based container then `porter-runtime` will be a Linux binary.
