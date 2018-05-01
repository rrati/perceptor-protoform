package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/koki/short/types"
)

// WARNING: If you add a config value, make sure to
// add it to the parameterize function as well !
type ProtoformConfig struct {
	// Dry run wont actually install, but will print the objects definitions out.
	DryRun bool

	// CONTAINER CONFIGS
	// These are sed replaced into the config maps for the containers.
	PerceptorPort                    int
	ScannerPort                      int
	PerceiverPort                    int
	ImageFacadePort                  int
	InternalDockerRegistries         []string
	AnnotationIntervalSeconds        int
	DumpIntervalMinutes              int
	HubHost                          string
	HubUser                          string
	HubUserPassword                  string
	HubPort                          int
	HubClientTimeoutPerceptorSeconds int
	HubClientTimeoutScannerSeconds   int
	ConcurrentScanLimit              int
	Namespace                        string
	Defaultversion                   string

	// CONTAINER PULL CONFIG
	// These are for defining docker registry and image location and versions
	Registry  string
	ImagePath string

	PerceptorImageName      string
	ScannerImageName        string
	PodPerceiverImageName   string
	ImagePerceiverImageName string
	ImageFacadeImageName    string

	PerceptorContainerVersion   string
	ScannerContainerVersion     string
	PerceiverContainerVersion   string
	ImageFacadeContainerVersion string

	// AUTH CONFIGS
	// These are given to containers through secrets or other mechanisms.
	// Not necessarily a one-to-one text replacement.
	// TODO Lets try to have this injected on serviceaccount
	// at pod startup, eventually Service accounts.
	DockerPasswordOrToken string
	DockerUsername        string

	ServiceAccounts map[string]string
	Openshift       bool

	// CPU and memory configurations
	DefaultCPU string // Should be passed like: e.g. "300m"
	DefaultMem string // Should be passed like: e.g "1300Mi"

	// Log level
	LogLevel string

	// Viper secrets
	ViperSecret string

	// Environment Variables
	HubUserPasswordEnvVar string

	// Automate test
	PerceptorSkyfire bool
}

func (p *ProtoformConfig) setDefaultValues() {
	if p.PerceptorPort == 0 {
		p.PerceptorPort = 3001
	}
	if p.PerceiverPort == 0 {
		p.PerceiverPort = 3002
	}
	if p.ScannerPort == 0 {
		p.ScannerPort = 3003
	}
	if p.ImageFacadePort == 0 {
		p.ImageFacadePort = 3004
	}
	if p.AnnotationIntervalSeconds == 0 {
		p.AnnotationIntervalSeconds = 30
	}
	if p.DumpIntervalMinutes == 0 {
		p.DumpIntervalMinutes = 30
	}
	if p.HubClientTimeoutPerceptorSeconds == 0 {
		p.HubClientTimeoutPerceptorSeconds = 5
	}
	if p.HubClientTimeoutScannerSeconds == 0 {
		p.HubClientTimeoutScannerSeconds = 30
	}
	if p.HubHost == "" {
		// meaningless default unless your in same namespace as hub.
		p.HubHost = "nginx-webapp-logstash"
	}
	if p.HubUser == "" {
		p.HubUser = "sysadmin"
	}
	if p.HubUserPassword == "" {
		panic("config failing: cannot continue without a hub password!!!")
	}
	if p.HubPort == 0 {
		p.HubPort = 8443
	}
	if p.DockerUsername == "" {
		p.DockerUsername = "admin"
	}
	if p.DockerPasswordOrToken == "" {
		log.Printf("config ERROR : cannot continue without a Docker password!!!")
	}
	if p.ConcurrentScanLimit == 0 {
		p.ConcurrentScanLimit = 7
	}
	if p.InternalDockerRegistries == nil {
		p.InternalDockerRegistries = []string{"docker-registry.default.svc:5000", "172.1.1.0:5000"}
	}
	if p.Defaultversion == "" {
		p.Defaultversion = "master"
	}
	if p.Registry == "" {
		p.Registry = "gcr.io"
	}
	if p.ImagePath == "" {
		p.ImagePath = "gke-verification/blackducksoftware"
	}
	if p.PerceptorImageName == "" {
		p.PerceptorImageName = "perceptor"
	}
	if p.ScannerImageName == "" {
		p.ScannerImageName = "perceptor-scanner"
	}
	if p.ImagePerceiverImageName == "" {
		p.ImagePerceiverImageName = "image-perceiver"
	}
	if p.PodPerceiverImageName == "" {
		p.PodPerceiverImageName = "pod-perceiver"
	}
	if p.ImageFacadeImageName == "" {
		p.ImageFacadeImageName = "perceptor-imagefacade"
	}
	if p.PerceptorContainerVersion == "" {
		p.PerceptorContainerVersion = p.Defaultversion
	}
	if p.ScannerContainerVersion == "" {
		p.ScannerContainerVersion = p.Defaultversion
	}
	if p.PerceiverContainerVersion == "" {
		p.PerceiverContainerVersion = p.Defaultversion
	}
	if p.ImageFacadeContainerVersion == "" {
		p.ImageFacadeContainerVersion = p.Defaultversion
	}
	if p.LogLevel == "" {
		p.LogLevel = "debug"
	}
}

func generateStringFromStringArr(strArr []string) string {
	str, _ := json.Marshal(strArr)
	return string(str)
}

// prometheus.yml
func (p *ProtoformConfig) ToConfigMap() []*types.ConfigMap {
	p.setDefaultValues()
	// TODO, parameterize prometheus
	// strings.Replace(prometheus_t,
	configs := map[string]string{
		"prometheus":                   "prometheus.yml",
		"perceptor-scanner-config":     "perceptor_scanner_conf.yaml",
		"perceiver":                    "perceiver.yaml",
		"perceptor-config":             "perceptor_conf.yaml",
		"perceptor-imagefacade-config": "perceptor_imagefacade_conf.yaml",
		"skyfire-config":               "skyfire_conf.yaml",
	}

	// Sed replace these.  Fine grained control over the json default format
	// makes this easier then actually modelling / mutating nested json in golang.
	// (I think)? Due to the fct that nested anonymous []string's seem to not be
	// "a thing".
	defaults := map[string]string{
		"prometheus":                   fmt.Sprint(`{"global":{"scrape_interval":"5s"},"scrape_configs":[{"job_name":"perceptor-scrape","scrape_interval":"5s","static_configs":[{"targets":["`, p.PerceptorImageName, `:`, p.PerceptorPort, `","`, p.ScannerImageName, `:`, p.ScannerPort, `","`, p.ImagePerceiverImageName, `:`, p.PerceiverPort, `","`, p.PodPerceiverImageName, `:`, p.PerceiverPort, `","`, p.ImageFacadeImageName, `:`, p.ImageFacadePort, `","skyfire:`, "3005", `"]}]}]}`),
		"perceptor-config":             fmt.Sprint(`{"HubHost": "`, p.HubHost, `","HubPort": "`, p.HubPort, `","HubUser": "`, p.HubUser, `","HubUserPasswordEnvVar": "`, p.HubUserPasswordEnvVar, `","HubClientTimeoutSeconds": "`, p.HubClientTimeoutPerceptorSeconds, `","ConcurrentScanLimit": "`, p.ConcurrentScanLimit, `","Port": "`, p.PerceptorPort, `","LogLevel": "`, p.LogLevel, `"}`),
		"perceptor-scanner-config":     fmt.Sprint(`{"HubHost": "`, p.HubHost, `","HubPort": "`, p.HubPort, `","HubUser": "`, p.HubUser, `","HubUserPasswordEnvVar": "`, p.HubUserPasswordEnvVar, `","HubClientTimeoutSeconds": "`, p.HubClientTimeoutScannerSeconds, `","Port": "`, p.ScannerPort, `","PerceptorHost": "`, p.PerceptorImageName, `","PerceptorPort": "`, p.PerceptorPort, `","ImageFacadePort": "`, p.ImageFacadePort, `","LogLevel": "`, p.LogLevel, `"}`),
		"perceiver":                    fmt.Sprint(`{"PerceptorHost": "`, p.PerceptorImageName, `","PerceptorPort": "`, p.PerceptorPort, `","AnnotationIntervalSeconds": "`, p.AnnotationIntervalSeconds, `","DumpIntervalMinutes": "`, p.DumpIntervalMinutes, `","Port": "`, p.PerceiverPort, `","LogLevel": "`, p.LogLevel, `"}`),
		"perceptor-imagefacade-config": fmt.Sprint(`{"DockerUser": "`, p.DockerUsername, `","DockerPassword": "`, p.DockerPasswordOrToken, `","Port": "`, p.ImageFacadePort, `","InternalDockerRegistries": `, generateStringFromStringArr(p.InternalDockerRegistries), `,"LogLevel": "`, p.LogLevel, `"}`),
	}

	if p.PerceptorSkyfire {
		defaults["skyfire-config"] = fmt.Sprint(`{"UseInClusterConfig": "`, "true", `","Port": "`, "3005", `","HubHost": "`, p.HubHost, `","HubPort": "`, p.HubPort, `","HubUser": "`, p.HubUser, `","HubUserPasswordEnvVar": "`, p.HubUserPasswordEnvVar, `","HubClientTimeoutSeconds": "`, p.HubClientTimeoutScannerSeconds, `","PerceptorHost": "`, p.PerceptorImageName, `","PerceptorPort": "`, p.PerceptorPort, `","ImageFacadePort": "`, p.ImageFacadePort,
			`","LogLevel": "`, p.LogLevel, `"}`)
	}

	maps := make([]*types.ConfigMap, len(configs))
	x := 0
	for config, filename := range configs {
		contents := defaults[config]
		maps[x] = &types.ConfigMap{
			Name: config,
			Data: map[string]string{
				filename: contents,
			},
		}
		x = x + 1
	}
	return maps
}
