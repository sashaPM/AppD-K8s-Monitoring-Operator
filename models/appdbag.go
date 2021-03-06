package models

type AppDBag struct {
	AgentNamespace              string
	AppName                     string
	TierName                    string
	NodeName                    string
	AppID                       int
	TierID                      int
	NodeID                      int
	Account                     string
	GlobalAccount               string
	AccessKey                   string
	ControllerUrl               string
	ControllerPort              uint16
	RestAPIUrl                  string
	SSLEnabled                  bool
	SystemSSLCert               string
	AgentSSLCert                string
	EventKey                    string
	EventServiceUrl             string
	RestAPICred                 string
	EventAPILimit               int
	PodSchemaName               string
	NodeSchemaName              string
	DeploySchemaName            string
	RSSchemaName                string
	DaemonSchemaName            string
	EventSchemaName             string
	ContainerSchemaName         string
	EpSchemaName                string
	NsSchemaName                string
	RqSchemaName                string
	JobSchemaName               string
	LogSchemaName               string
	DashboardTemplatePath       string
	DashboardSuffix             string
	DashboardDelayMin           int
	AgentEnvVar                 string
	AgentLabel                  string
	AppDAppLabel                string
	AppDTierLabel               string
	AppDAnalyticsLabel          string
	AgentMountName              string
	AgentMountPath              string
	AppLogMountName             string
	AppLogMountPath             string
	JDKMountName                string
	JDKMountPath                string
	NodeNamePrefix              string
	AnalyticsAgentUrl           string
	AnalyticsAgentImage         string
	AnalyticsAgentContainerName string
	AppDInitContainerName       string
	AppDJavaAttachImage         string
	AppDDotNetAttachImage       string
	AppDNodeJSAttachImage       string
	ProxyUrl                    string
	ProxyHost                   string
	ProxyPort                   string
	ProxyUser                   string
	ProxyPass                   string
	InitContainerDir            string
	MetricsSyncInterval         int // Frequency of metrics pushes to the controller, sec
	SnapshotSyncInterval        int // Frequency of snapshot pushes to events api, sec
	AgentServerPort             int
	NsToMonitor                 []string
	NsToMonitorExclude          []string
	DeploysToDashboard          []string
	NodesToMonitor              []string
	NodesToMonitorExclude       []string
	NsToInstrument              []string
	NsToInstrumentExclude       []string
	NSInstrumentRule            []AgentRequest
	InstrumentationMethod       InstrumentationMethod
	DefaultInstrumentationTech  TechnologyName
	BiqService                  string
	InstrumentContainer         string //all, first, name
	InstrumentMatchString       []string
	InitRequestMem              string
	InitRequestCpu              string
	BiqRequestMem               string
	BiqRequestCpu               string
	LogLines                    int //0 - no logging
	PodEventNumber              int
	RemoteBiqProtocol           string
	RemoteBiqHost               string
	RemoteBiqPort               int
	SchemaUpdateCache           []string
	LogLevel                    string
	OverconsumptionThreshold    int //percent
}

func IsUpdatable(fieldName string) bool {
	arr := []string{"AgentNamespace", "AppName", "TierName", "NodeName", "AppID", "TierID", "NodeID", "Account", "GlobalAccount", "AccessKey", "ControllerUrl",
		"ControllerPort", "RestAPIUrl", "SSLEnabled", "SystemSSLCert", "AgentSSLCert", "EventKey", "EventServiceUrl", "RestAPICred"}
	for _, s := range arr {
		if s == fieldName {
			return false
		}
	}
	return true
}
