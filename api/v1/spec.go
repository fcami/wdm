package v1

const Version = "v1"

type DependencyObjSpec struct {
	//Version     string `json:"version"`
	//Description string `json:"description,omitempty"`
	//TestHistory int    `json:"test_history"`
	//Matrices    map[string]MatrixSpec `json:"matrices,omitempty"`
	Req []string `json:"requirements,omitempty"`
}

type DependencyObj struct {
	Name string            `json:"name"`
	Spec DependencyObjSpec `json:"spec,omitempty"`
}
