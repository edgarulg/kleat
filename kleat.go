package main

import (
	"errors"
	"github.com/ezimanyi/kleat/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

func main() {
	fn := os.Args[1]
	h := parseHalConfig(fn)
	if err := validateHalConfig(h); err != nil {
		panic(err)
	}
	if err := printHalConfig(*h); err != nil {
		panic(err)
	}
}

type HalConfigValidator func(*proto.HalConfig) []ValidationFailure

func getValidators() []HalConfigValidator {
	return []HalConfigValidator{
		validateKindsAndOmitKinds,
	}
}

func validateHalConfig(h *proto.HalConfig) error {
	messages := getValidationMessages(h, getValidators())
	if len(messages) > 0 {
		msg := strings.Join(messages, "\n")
		return errors.New(msg)
	}
	return nil
}

func getValidationMessages(h *proto.HalConfig, fa []HalConfigValidator) []string {
	var messages []string
	for _, f := range fa {
		rs := f(h)
		for _, r := range rs {
			messages = append(messages, r.msg)
		}
	}
	return messages
}

func validateKindsAndOmitKinds(h *proto.HalConfig) []ValidationFailure {
	var messages []ValidationFailure
	for _, a := range h.Providers.Kubernetes.Accounts {
		if !(len(a.Kinds) == 0) && !(len(a.OmitKinds) == 0) {
			messages = append(messages, fatalResult("Cannot specify both kinds and omitKinds."))
		}
	}
	return messages
}

type ValidationFailure struct {
	msg string
}

func fatalResult(msg string) ValidationFailure {
	return ValidationFailure{
		msg: msg,
	}
}

func parseHalConfig(fn string) *proto.HalConfig {
	dat, err := ioutil.ReadFile(fn)

	h := proto.HalConfig{}
	err = yaml.Unmarshal([]byte(dat), &h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &h
}

func printHalConfig(h proto.HalConfig) error {
	c, _ := halToClouddriver(h)
	if err := printObject(c, os.Stdout); err != nil {
		return err
	}
	return nil
	//todo: test whether enum providerVersion marshals correctly
}

func printObject(i interface{}, w io.Writer) error {
	bytes, err := yaml.Marshal(i)
	if err != nil {
		return err
	}
	if _, err = w.Write(bytes); err != nil {
		return err
	}
	return nil
}

func getTestHalConfig() proto.HalConfig {
	h := proto.HalConfig{
		Providers: &proto.HalConfig_Providers{
			Kubernetes: &proto.Kubernetes{
				Enabled: false,
				Accounts: []*proto.Kubernetes_Account{
					{
						Name:            "hal",
						ProviderVersion: proto.Kubernetes_V2,
					},
				},
			},
		},
	}
	return h
}

func halToFront50(h proto.HalConfig) (proto.Front50Config, error) {
	f := proto.Front50Config{
		Spinnaker: &proto.Front50Config_Spinnaker{
			PersistentStoreType: h.PersistentStorage.PersistentStoreType,
			Gcs:                 h.PersistentStorage.Gcs,
		},
	}
	return f, nil
}

func extractPersistentStoreType(h proto.HalConfig) *string {
	if h.PersistentStorage == nil {
		return nil
	}
	return &h.PersistentStorage.PersistentStoreType
}

func halToClouddriver(h proto.HalConfig) (proto.ClouddriverConfig, error) {
	c := proto.ClouddriverConfig{
		Kubernetes: h.Providers.Kubernetes,
		Google:     h.Providers.Google,
	}
	return c, nil
}
