// TODO: config is something else. rename config to resource or object or deps...

package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	// v1 "github.com/openshift-psap/ci-dashboard/api/matrix/v1"

	"sigs.k8s.io/yaml"
)

func LoadYamlFile(yamlFile string) ([]v1.DependencyObj, error) {
	var err error
	var configYaml []byte
	fmt.Println("Reading from", configFile)
	if configFile == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			configYaml = append(configYaml, scanner.Bytes()...)
			configYaml = append(configYaml, '\n')
		}
	} else {
		configYaml, err = ioutil.ReadFile(configFile)
		if err != nil {
			return nil, fmt.Errorf("read error: %v", err)
		}
	}
	var objs []v1.DependencyObj
	err = yaml.Unmarshal(configYaml, &objs)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}
	fmt.Println("--")
	str, err := yaml.Marshal(objs)
	fmt.Println(string(str))
	fmt.Println("--")
	return &objs, nil
}
