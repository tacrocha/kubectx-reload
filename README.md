# kubectx-reload
![Latest GitHub release](https://img.shields.io/github/release/tacrocha/kubectx-reload.svg)
![GitHub contributors](https://img.shields.io/github/contributors/tacrocha/kubectx-reload.svg)
![GitHub stars](https://img.shields.io/github/stars/tacrocha/kubectx-reload.svg)
![Written in Bash](https://img.shields.io/badge/written%20in-go-ff69b4.svg)

A tool that merges multiple kubeconfig files for the wonderful and essential [kubectx](https://github.com/ahmetb/kubectx).

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
