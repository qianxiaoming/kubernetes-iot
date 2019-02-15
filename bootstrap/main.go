package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		help bool
		command string
		sourceDir string
		targetDir string
		valueFile string
	)
	flag.BoolVar(&help, "h", false, "Show this messages")
	flag.StringVar(&command, "c", "init", "Command to be executed")
	flag.StringVar(&sourceDir, "s", "", "Source directory containing template YAML files")
	flag.StringVar(&targetDir, "t", "", "Target directory to save generated kubernetes YAML files")
	flag.StringVar(&valueFile, "v", "", "Configuration values used to generate kubernetes YAML files")
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	if command == "init" && !GenerateYAMLs(sourceDir, targetDir, valueFile) {
		os.Exit(1)
	}
}

type VolumeInfo struct {
	Replicas int
	Size string
	Path string
}

type NodeLabels struct {
	Service string
	Storage string
	Gateway string
}

type RequestResources struct {
	CPU string
	Memory string
	Storage string
}

type ZookeeperConfig struct {
	AppName string `yaml:"appName"`
	InitImage string `yaml:"initImage"`
	Image string `yaml:"image"`
	Replicas int `yaml:"replicas"`
	TickTime int `yaml:"tickTime"`
	MaxClientCnxns int `yaml:"maxClientCnxns"`
	InitLimit int `yaml:"initLimit"`
	SyncLimit int `yaml:"syncLimit"`
	Resources RequestResources `yaml:"resources"`
}

type ConnectConfig struct {
	Image string `yaml:"image"`
	Resources RequestResources `yaml:"resources"`
}

type KafkaConfig struct {
	AppName string `yaml:"appName"`
	Image string `yaml:"image"`
	Replicas int `yaml:"replicas"`
	Zkchroot string `yaml:"zkchroot"`
	Resources RequestResources `yaml:"resources"`
	Connect ConnectConfig `yaml:"connect"`
}

type ConnectorConfig struct {
	Source string `yaml:"source"`
	Sink string `yaml:"sink"`
}

type ElasticsearchConfig struct {
	AppName string `yaml:"appName"`
	InitImage string `yaml:"initImage"`
	Image string `yaml:"image"`
	Replicas int `yaml:"replicas"`
	JavaOpts string `yaml:"javaOpts"`
	Resources RequestResources `yaml:"resources"`
	Connector ConnectorConfig `yaml:"connector"`
}

type DeploymentConfig struct {
	Namespace string `yaml:"namespace"`
	LabelConfig NodeLabels `yaml:"labelConfig"`
	ConfigVolume VolumeInfo `yaml:"configVolume"`
	StreamVolume VolumeInfo `yaml:"streamVolume"`
	AnalysisVolume VolumeInfo `yaml:"analysisVolume"`
	Zookeeper ZookeeperConfig `yaml:"zookeeper"`
	Kafka KafkaConfig `yaml:"kafka"`
	Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
}

func int2slice(v int) []int {
	slice := make([]int, v)
	for i := 1; i <= v; i++ {
		slice[i-1] = i
	}
	return slice
}

func GenerateYAMLs(sourceDir string, targetDir string, valueFile string) bool {
	info, err := os.Stat(sourceDir)
	if err != nil || !info.IsDir() {
		log.Printf("%s is not accessible or is not a directory", sourceDir)
		return false
	}
	info, err = os.Stat(targetDir)
	if err != nil {
		if !os.IsNotExist(err) {
			return false
		} else {
			err = os.MkdirAll(targetDir, os.ModePerm)
			if err != nil {
				log.Printf("Cannot create target directory %s: %v", targetDir, err)
				return false
			}
		}
	} else if !info.IsDir() {
		log.Printf("%s is not a directory", sourceDir)
		return false
	}
	info, err = os.Stat(valueFile)
	if err != nil {
		log.Printf("%s is not accessible or does not exist", valueFile)
		return false
	}

	config := new(DeploymentConfig)
	file, err := ioutil.ReadFile(valueFile)
	if err != nil {
		log.Printf("Unable to read value file: %v", err)
		return false
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("Parse value file failed: %v", err)
	}
	fmt.Printf("%v\n", config)

	err = filepath.Walk(sourceDir, func(path string, f os.FileInfo, err error) error {
		if f == nil { return err }
		if f.IsDir() { return nil }
		if !strings.HasSuffix(path, ".tmpl") { return nil }

		tmplFile, yamlFile := path, targetDir + string(os.PathSeparator) + path[len(sourceDir)+1:]
		yamlFile = yamlFile[0:len(yamlFile)-5]
		fmt.Printf("Generating %s...\n", yamlFile)
		os.MkdirAll(filepath.Dir(yamlFile), os.ModePerm)
		file, err := os.Create(yamlFile)
		if err != nil {
			log.Fatalf("Cannot create output yaml file %s: %v", yamlFile, err)
		}
		defer file.Close()

		t := template.New(filepath.Base(tmplFile))
		t.Funcs(template.FuncMap{"int2slice": int2slice})
		t, err = t.ParseFiles(tmplFile)
		if err != nil {
			return err
		}
		return t.Execute(file, config)
	})
	return err == nil
}
