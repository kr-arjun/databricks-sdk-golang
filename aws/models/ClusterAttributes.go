package models

type ClusterAttributes struct {
	ClusterName            string           `json:"cluster_name,omitempty" url:"cluster_name,omitempty"`
	SparkVersion           string           `json:"spark_version,omitempty" url:"spark_version,omitempty"`
	SparkConf              SparkConfPair    `json:"spark_conf,omitempty" url:"spark_conf,omitempty"`
	AwsAttributes          AwsAttributes    `json:"aws_attributes,omitempty" url:"aws_attributes,omitempty"`
	NodeTypeID             string           `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	DriverNodeTypeID       string           `json:"driver_node_type_id,omitempty" url:"driver_node_type_id,omitempty"`
	SSHPublicKeys          []string         `json:"ssh_public_keys,omitempty" url:"ssh_public_keys,omitempty"`
	CustomTags             []ClusterTag     `json:"custom_tags,omitempty" url:"custom_tags,omitempty"`
	ClusterLogConf         ClusterLogConf   `json:"cluster_log_conf,omitempty" url:"cluster_log_conf,omitempty"`
	InitScripts            []InitScriptInfo `json:"init_scripts,omitempty" url:"init_scripts,omitempty"`
	SparkEnvVars           SparkEnvPair     `json:"spark_env_vars,omitempty" url:"spark_env_vars,omitempty"`
	AutoterminationMinutes int32            `json:"autotermination_minutes,omitempty" url:"autotermination_minutes,omitempty"`
	EnableElasticDisk      bool             `json:"enable_elastic_disk,omitempty" url:"enable_elastic_disk,omitempty"`
	ClusterSource          AwsAvailability  `json:"cluster_source,omitempty" url:"cluster_source,omitempty"`
}
