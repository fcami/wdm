// TODO: config is something else. rename config to resource or object or deps...

package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	v1 "github.com/openshift-psap/wdm/api/v1"

	"sigs.k8s.io/yaml"
)

func LoadYamlFile(yamlFile string) ([]v1.DependencyObj, error) {
	var err error
	var configYaml []byte
	var dependencyObjs []v1.DependencyObj
	fmt.Println("Reading from", yamlFile)
	if yamlFile == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			configYaml = append(configYaml, scanner.Bytes()...)
			configYaml = append(configYaml, '\n')
		}
	} else {
		configYaml, err = ioutil.ReadFile(yamlFile)
		if err != nil {
			return nil, fmt.Errorf("read error: %v", err)
		}
	}
	for _, doc := range bytes.Split(configYaml, []byte("---")) {
		var obj v1.DependencyObj
		err = yaml.Unmarshal(doc, &obj)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error: %v", err)
		}
		fmt.Println("--")
		str, err := yaml.Marshal(obj)
		fmt.Println(string(str))
		fmt.Println("--")
		fmt.Println(err)
		dependencyObjs = append(dependencyObjs, obj)
	}
	return dependencyObjs, nil
}
