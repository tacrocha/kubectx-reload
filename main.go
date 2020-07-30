package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {

	const filePrefix string = "kubeconfig-"

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	var kubeDir string = filepath.Join(user.HomeDir, ".kube")

	files, err := ioutil.ReadDir(kubeDir)
	if err != nil {
		log.Fatal(err)
	}

	var kubeFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), filePrefix) {
			kubeFiles = append(kubeFiles, filepath.Join(kubeDir, file.Name()))
		}
	}

	kubectl, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatal(err)
	}

	configViewCmd := exec.Command(kubectl, "config", "view", "--flatten")
	configViewCmd.Env = append(os.Environ(), "KUBECONFIG="+strings.Join(kubeFiles, ":"))

	kubectlOut, err := configViewCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(filepath.Join(kubeDir, "config"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(kubectlOut)

	kubectx, err := exec.LookPath("kubectx")
	if err != nil {
		log.Fatal(err)
	}

	kubectxCmd := exec.Command(kubectx)
	if kubectxOut, err := kubectxCmd.Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(kubectxOut))
	}

}
