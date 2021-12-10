package sarif

import "time"

// Invocation describes the runtime environment of the analysis tool run.
type Invocation struct {
	PropertyBag
	Account                            *string                  `json:"account,omitempty"`
	Arguments                          []string                 `json:"arguments,omitempty"`
	CommandLine                        *string                  `json:"commandLine,omitempty"`
	EndTimeUTC                         *time.Time               `json:"endTimeUtc,omitempty"`
	EnvironmentVariables               map[string]string        `json:"environmentVariables,omitempty"`
	ExecutableLocation                 *ArtifactLocation        `json:"executableLocation,omitempty"`
	ExecutionSuccessful                *bool                    `json:"executionSuccessful"`
	ExitCode                           *int                     `json:"exitCode,omitempty"`
	ExitCodeDescription                *string                  `json:"exitCodeDescription,omitempty"`
	ExitSignalName                     *string                  `json:"exitSignalName,omitempty"`
	ExitSignalNumber                   *int                     `json:"exitSignalNumber,omitempty"`
	Machine                            *string                  `json:"machine,omitempty"`
	NotificationConfigurationOverrides []*ConfigurationOverride `json:"notificationConfigurationOverrides,omitempty"`
	ProcessID                          *int                     `json:"processId,omitempty"`
	ProcessStartFailureMessage         *string                  `json:"processStartFailureMessage,omitempty"`
	ResponseFiles                      []*ArtifactLocation      `json:"responseFiles,omitempty"`
	RuleConfigurationOverrides         []*ConfigurationOverride `json:"ruleConfigurationOverrides,omitempty"`
	StartTimeUTC                       *time.Time               `json:"startTimeUtc,omitempty"`
	Stderr                             *ArtifactLocation        `json:"stderr,omitempty"`
	Stdin                              *ArtifactLocation        `json:"stdin,omitempty"`
	Stdout                             *ArtifactLocation        `json:"stdout,omitempty"`
	StdoutStderr                       *ArtifactLocation        `json:"stdoutStderr,omitempty"`
	ToolConfigurationNotifications     []*Notification          `json:"toolConfigurationNotifications,omitempty"`
	ToolExecutionNotifications         []*Notification          `json:"toolExecutionNotifications,omitempty"`
	WorkingDirectory                   *ArtifactLocation        `json:"workingDirectory,omitempty"`
}

func NewInvocation() *Invocation {
	return &Invocation{}
}

func (invocation *Invocation) WithAccount(account string) *Invocation {
	invocation.Account = &account
	return invocation
}

func (invocation *Invocation) WithArguments(arguments []string) *Invocation {
	invocation.Arguments = arguments
	return invocation
}

func (invocation *Invocation) AddArgument(argument string) {
	invocation.Arguments = append(invocation.Arguments, argument)
}

func (invocation *Invocation) WithCommanLine(commandLine string) *Invocation {
	invocation.CommandLine = &commandLine
	return invocation
}

// WithEndTimeUTC sets the instant when the invocation ended and returns the same Invocation.
func (invocation *Invocation) WithEndTimeUTC(endTime time.Time) *Invocation {
	endTimeUTC := endTime.UTC()
	invocation.EndTimeUTC = &endTimeUTC
	return invocation
}

func (invocation *Invocation) WithEnvironmentVariables(environmentVariables map[string]string) *Invocation {
	invocation.EnvironmentVariables = environmentVariables
	return invocation
}

func (invocation *Invocation) SetEnvironmentVariable(name, value string) {
	invocation.EnvironmentVariables[name] = value
}

func (invocation *Invocation) WithExecutableLocation(executableLocation *ArtifactLocation) *Invocation {
	invocation.ExecutableLocation = executableLocation
	return invocation
}
func (invocation *Invocation) WithExecutionSuccess(executionSuccessful bool) *Invocation {
	invocation.ExecutionSuccessful = &executionSuccessful
	return invocation
}
func (invocation *Invocation) WithExitCode(exitCode int) *Invocation {
	invocation.ExitCode = &exitCode
	return invocation
}
func (invocation *Invocation) WithExitCodeDescription(exitCodeDescription string) *Invocation {
	invocation.ExitCodeDescription = &exitCodeDescription
	return invocation
}
func (invocation *Invocation) WithExitSignalNumber(exitSignalNumber int) *Invocation {
	invocation.ExitSignalNumber = &exitSignalNumber
	return invocation
}
func (invocation *Invocation) WithExitSignalName(exitSignalName string) *Invocation {
	invocation.ExitSignalName = &exitSignalName
	return invocation
}
func (invocation *Invocation) WithMachine(machine string) *Invocation {
	invocation.Machine = &machine
	return invocation
}

func (invocation *Invocation) WithNotificationConfigurationOverrides(overrides []*ConfigurationOverride) *Invocation {
	invocation.NotificationConfigurationOverrides = overrides
	return invocation
}

func (invocation *Invocation) AddNotificationConfigurationOverride(override *ConfigurationOverride) {
	invocation.NotificationConfigurationOverrides = append(invocation.NotificationConfigurationOverrides, override)
}

func (invocation *Invocation) WithProcessID(processID int) *Invocation {
	invocation.ProcessID = &processID

	return invocation
}
func (invocation *Invocation) WithProcessStartFailureMessage(failureMessage string) *Invocation {
	invocation.ProcessStartFailureMessage = &failureMessage
	return invocation
}

func (invocation *Invocation) WithResponseFiles(responseFiles []*ArtifactLocation) *Invocation {
	invocation.ResponseFiles = responseFiles
	return invocation
}

func (invocation *Invocation) AddResponseFile(responseFile *ArtifactLocation) {
	invocation.ResponseFiles = append(invocation.ResponseFiles, responseFile)
}

func (invocation *Invocation) WithRuleConfigurationOverrides(overrides []*ConfigurationOverride) *Invocation {
	invocation.RuleConfigurationOverrides = overrides
	return invocation
}

func (invocation *Invocation) AddRuleConfigurationOverride(override *ConfigurationOverride) {
	invocation.RuleConfigurationOverrides = append(invocation.RuleConfigurationOverrides, override)
}

// WithStartTimeUTC sets the instant when the invocation started and returns the same Invocation.
func (invocation *Invocation) WithStartTimeUTC(startTime time.Time) *Invocation {
	startTimeUTC := startTime.UTC()
	invocation.StartTimeUTC = &startTimeUTC
	return invocation
}

func (invocation *Invocation) WithStdErr(stdErr *ArtifactLocation) *Invocation {
	invocation.Stderr = stdErr
	return invocation
}

func (invocation *Invocation) WithStdIn(stdIn *ArtifactLocation) *Invocation {
	invocation.Stdin = stdIn
	return invocation
}

func (invocation *Invocation) WithStdout(stdOut *ArtifactLocation) *Invocation {
	invocation.Stdout = stdOut
	return invocation
}
func (invocation *Invocation) WithStdoutStderr(stdoutStderr *ArtifactLocation) *Invocation {
	invocation.StdoutStderr = stdoutStderr
	return invocation
}

func (invocation *Invocation) WithToolConfigurationNotifications(toolConfigNotifications []*Notification) *Invocation {
	invocation.ToolConfigurationNotifications = toolConfigNotifications
	return invocation
}

func (invocation *Invocation) AddToolConfigurationNotification(toolConfigNotification *Notification) {
	invocation.ToolConfigurationNotifications = append(invocation.ToolConfigurationNotifications, toolConfigNotification)
}

func (invocation *Invocation) WithToolExecutionNotifications(toolExecutionNotification []*Notification) *Invocation {
	invocation.ToolExecutionNotifications = toolExecutionNotification
	return invocation
}

func (invocation *Invocation) AddTToolExecutionNotification(toolExecutionNotification *Notification) {
	invocation.ToolExecutionNotifications = append(invocation.ToolExecutionNotifications, toolExecutionNotification)
}

// WithWorkingDirectory sets the current working directory of the invocation and returns the same Invocation.
func (invocation *Invocation) WithWorkingDirectory(workingDirectory *ArtifactLocation) *Invocation {
	invocation.WorkingDirectory = workingDirectory
	return invocation
}
