package v1

const Version = "v1"

type DependencyObj struct {
}

type DependencyObjSpec struct {
	Version     string                `json:"version"`
	Description string                `json:"description,omitempty"`
	TestHistory int                   `json:"test_history"`
	Matrices    map[string]MatrixSpec `json:"matrices,omitempty"`
}

type ToolboxStepResult struct {
	Name string

	Ok       int
	Failures int
	Ignored  int

	ExpectedFailure string

	FlakeFailure string
}

type TestResult struct {
	BuildId    string
	Passed     bool
	Result     string
	FinishDate string

	StepExecuted bool
	StepPassed   bool
	StepResult   string

	Warnings map[string]string
	Flakes   map[string]string

	/* *** */

	OperatorVersion    string
	OpenShiftVersion   string
	CiArtifactsVersion string

	/* *** */
	TestSpec *TestSpec

	ToolboxSteps []string

	ToolboxStepsResults []ToolboxStepResult

	/* *** */

	Ok       int
	Failures int
	Ignored  int

	FlakeFailure bool
}

type TestSpec struct {
	TestName        string `json:"test_name,omitempty"`
	Branch          string `json:"branch,omitempty"`
	OperatorVersion string `json:"operator_version,omitempty"`
	Variant         string `json:"variant,omitempty"`
	ProwStep        string `json:"prow_step,omitempty"`

	/* *** */
	Matrix *MatrixSpec

	ProwName  string
	TestGroup string

	OldTests []*TestResult
}

type MatrixSpec struct {
	Description    string                `json:"description,omitempty"`
	ViewerURL      string                `json:"viewer_url,omitempty"`
	ArtifactsURL   string                `json:"artifacts_url,omitempty"`
	ArtifactsCache string                `json:"artifacts_cache,omitempty"`
	ProwConfig     string                `json:"prow_config,omitempty"`
	ProwStep       string                `json:"prow_step,omitempty"`
	OperatorName   string                `json:"operator_name,omitempty"`
	Tests          map[string][]TestSpec `json:"tests,omitempty"`

	/* *** */

	Name string
}
