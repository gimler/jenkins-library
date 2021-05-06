// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type integrationArtifactGetServiceEndpointOptions struct {
	Username              string `json:"username,omitempty"`
	Password              string `json:"password,omitempty"`
	IntegrationFlowID     string `json:"integrationFlowId,omitempty"`
	Platform              string `json:"platform,omitempty"`
	Host                  string `json:"host,omitempty"`
	OAuthTokenProviderURL string `json:"oAuthTokenProviderUrl,omitempty"`
}

type integrationArtifactGetServiceEndpointCommonPipelineEnvironment struct {
	custom struct {
		iFlowServiceEndpoint string
	}
}

func (p *integrationArtifactGetServiceEndpointCommonPipelineEnvironment) persist(path, resourceName string) {
	content := []struct {
		category string
		name     string
		value    interface{}
	}{
		{category: "custom", name: "iFlowServiceEndpoint", value: p.custom.iFlowServiceEndpoint},
	}

	errCount := 0
	for _, param := range content {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(param.category, param.name), param.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting piper environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Piper environment")
	}
}

// IntegrationArtifactGetServiceEndpointCommand Get an deployed CPI intgeration flow service endpoint
func IntegrationArtifactGetServiceEndpointCommand() *cobra.Command {
	const STEP_NAME = "integrationArtifactGetServiceEndpoint"

	metadata := integrationArtifactGetServiceEndpointMetadata()
	var stepConfig integrationArtifactGetServiceEndpointOptions
	var startTime time.Time
	var commonPipelineEnvironment integrationArtifactGetServiceEndpointCommonPipelineEnvironment

	var createIntegrationArtifactGetServiceEndpointCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Get an deployed CPI intgeration flow service endpoint",
		Long:  `With this step you can obtain information about the service endpoints exposed by SAP Cloud Platform Integration on a tenant using OData API. Learn more about the SAP Cloud Integration remote API for getting service endpoint of deployed integration artifact [here](https://help.sap.com/viewer/368c481cd6954bdfa5d0435479fd4eaf/Cloud/en-US/d1679a80543f46509a7329243b595bdb.html).`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.Username)
			log.RegisterSecret(stepConfig.Password)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				commonPipelineEnvironment.persist(GeneralConfig.EnvRootPath, "commonPipelineEnvironment")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			integrationArtifactGetServiceEndpoint(stepConfig, &telemetryData, &commonPipelineEnvironment)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addIntegrationArtifactGetServiceEndpointFlags(createIntegrationArtifactGetServiceEndpointCmd, &stepConfig)
	return createIntegrationArtifactGetServiceEndpointCmd
}

func addIntegrationArtifactGetServiceEndpointFlags(cmd *cobra.Command, stepConfig *integrationArtifactGetServiceEndpointOptions) {
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User to authenticate to the SAP Cloud Platform Integration Service")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "Password to authenticate to the SAP Cloud Platform Integration Service")
	cmd.Flags().StringVar(&stepConfig.IntegrationFlowID, "integrationFlowId", os.Getenv("PIPER_integrationFlowId"), "Specifies the ID of the Integration Flow artifact")
	cmd.Flags().StringVar(&stepConfig.Platform, "platform", os.Getenv("PIPER_platform"), "Specifies the running platform of the SAP Cloud platform integraion service")
	cmd.Flags().StringVar(&stepConfig.Host, "host", os.Getenv("PIPER_host"), "Specifies the protocol and host address, including the port. Please provide in the format `<protocol>://<host>:<port>`. Supported protocols are `http` and `https`.")
	cmd.Flags().StringVar(&stepConfig.OAuthTokenProviderURL, "oAuthTokenProviderUrl", os.Getenv("PIPER_oAuthTokenProviderUrl"), "Specifies the oAuth Provider protocol and host address, including the port. Please provide in the format `<protocol>://<host>:<port>`. Supported protocols are `http` and `https`.")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("integrationFlowId")
	cmd.MarkFlagRequired("host")
	cmd.MarkFlagRequired("oAuthTokenProviderUrl")
}

// retrieve step metadata
func integrationArtifactGetServiceEndpointMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "integrationArtifactGetServiceEndpoint",
			Aliases:     []config.Alias{},
			Description: "Get an deployed CPI intgeration flow service endpoint",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name: "username",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "cpiCredentialsId",
								Param: "username",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name: "password",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "cpiCredentialsId",
								Param: "password",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name:        "integrationFlowId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "platform",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "host",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "oAuthTokenProviderUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
				},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "commonPipelineEnvironment",
						Type: "piperEnvironment",
						Parameters: []map[string]interface{}{
							{"Name": "custom/iFlowServiceEndpoint"},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
