# kubectx-reload
A simple golang script that merges multiple kubeconfig files for the wonderful and essential [kubectx](https://github.com/ahmetb/kubectx) tool.

## Installation 

```shell
go get github.com/tacrocha/kubectx-reload
```

## Usage

```shell
kubectx-reload
```

## Assumptions

1. `kubectl` and `kubectx` are installed and in the PATH;
2. Kubeconfig files are in `~/.kube`;
3. File names start with `kubeconfig-`, e.g.: `kubeconfig-mycontext01`, `kubeconfig-othercontext.yaml`
