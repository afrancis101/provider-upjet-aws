// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eks

import (
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"

	"github.com/upbound/provider-aws/apis/cluster/fis/v1beta1"
	"github.com/upbound/provider-aws/apis/cluster/fis/v1beta2"
)

// Configure adds configurations for the eks group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_fis_experiment_template", func(r *config.Resource) {
		r.Conversions = append(r.Conversions,
			conversion.NewCustomConverter("v1beta1", "v1beta2", templateConverterFromv1beta1Tov1beta2),
			conversion.NewCustomConverter("v1beta2", "v1beta1", templateConverterFromv1beta2Tov1beta1),
		)
	})
}

func templateConverterFromv1beta1Tov1beta2(src, target xpresource.Managed) error {
	srcTyped := src.(*v1beta1.ExperimentTemplate)
	targetTyped := target.(*v1beta2.ExperimentTemplate)

	if len(srcTyped.Spec.ForProvider.LogConfiguration) > 0 {
		if targetTyped.Spec.ForProvider.LogConfiguration == nil {
			targetTyped.Spec.ForProvider.LogConfiguration = &v1beta2.LogConfigurationParameters{}

			if len(srcTyped.Spec.ForProvider.LogConfiguration[0].CloudwatchLogsConfiguration) > 0 {
				if targetTyped.Spec.ForProvider.LogConfiguration.CloudwatchLogsConfiguration == nil {
					targetTyped.Spec.ForProvider.LogConfiguration.CloudwatchLogsConfiguration = &v1beta2.CloudwatchLogsConfigurationParameters{}
				}
				if srcTyped.Spec.ForProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn != nil {
					targetTyped.Spec.ForProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn = srcTyped.Spec.ForProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn
				}
			}

			if srcTyped.Spec.ForProvider.LogConfiguration[0].LogSchemaVersion != nil {
				targetTyped.Spec.ForProvider.LogConfiguration.LogSchemaVersion = srcTyped.Spec.ForProvider.LogConfiguration[0].LogSchemaVersion
			}

			if len(srcTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration) > 0 {
				if targetTyped.Spec.ForProvider.LogConfiguration.S3Configuration == nil {
					targetTyped.Spec.ForProvider.LogConfiguration.S3Configuration = &v1beta2.LogConfigurationS3ConfigurationParameters{}
				}

				if srcTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration[0].BucketName != nil {
					targetTyped.Spec.ForProvider.LogConfiguration.S3Configuration.BucketName = srcTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration[0].BucketName
				}

				if srcTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration[0].Prefix != nil {
					targetTyped.Spec.ForProvider.LogConfiguration.S3Configuration.Prefix = srcTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration[0].Prefix
				}
			}

		}
	}

	if len(srcTyped.Spec.InitProvider.LogConfiguration) > 0 {
		if targetTyped.Spec.InitProvider.LogConfiguration == nil {
			targetTyped.Spec.InitProvider.LogConfiguration = &v1beta2.LogConfigurationInitParameters{}

			if len(srcTyped.Spec.InitProvider.LogConfiguration[0].CloudwatchLogsConfiguration) > 0 {
				if targetTyped.Spec.InitProvider.LogConfiguration.CloudwatchLogsConfiguration == nil {
					targetTyped.Spec.InitProvider.LogConfiguration.CloudwatchLogsConfiguration = &v1beta2.CloudwatchLogsConfigurationInitParameters{}
				}
				if srcTyped.Spec.InitProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn != nil {
					targetTyped.Spec.InitProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn = srcTyped.Spec.InitProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn
				}
				if srcTyped.Spec.InitProvider.LogConfiguration[0].LogSchemaVersion != nil {
					targetTyped.Spec.InitProvider.LogConfiguration.LogSchemaVersion = srcTyped.Spec.InitProvider.LogConfiguration[0].LogSchemaVersion
				}

				if len(srcTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration) > 0 {
					if targetTyped.Spec.InitProvider.LogConfiguration.S3Configuration == nil {
						targetTyped.Spec.InitProvider.LogConfiguration.S3Configuration = &v1beta2.LogConfigurationS3ConfigurationInitParameters{}
					}

					if srcTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration[0].BucketName != nil {
						targetTyped.Spec.InitProvider.LogConfiguration.S3Configuration.BucketName = srcTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration[0].BucketName
					}

					if srcTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration[0].Prefix != nil {
						targetTyped.Spec.InitProvider.LogConfiguration.S3Configuration.Prefix = srcTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration[0].Prefix
					}
				}
			}
		}
	}
	if len(srcTyped.Status.AtProvider.LogConfiguration) > 0 {
		if targetTyped.Status.AtProvider.LogConfiguration == nil {
			targetTyped.Status.AtProvider.LogConfiguration = &v1beta2.LogConfigurationObservation{}

			if len(srcTyped.Status.AtProvider.LogConfiguration[0].CloudwatchLogsConfiguration) > 0 {
				if targetTyped.Status.AtProvider.LogConfiguration.CloudwatchLogsConfiguration == nil {
					targetTyped.Status.AtProvider.LogConfiguration.CloudwatchLogsConfiguration = &v1beta2.CloudwatchLogsConfigurationObservation{}
				}
				if srcTyped.Status.AtProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn != nil {
					targetTyped.Status.AtProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn = srcTyped.Status.AtProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn
				}
			}

			if srcTyped.Status.AtProvider.LogConfiguration[0].LogSchemaVersion != nil {
				targetTyped.Status.AtProvider.LogConfiguration.LogSchemaVersion = srcTyped.Status.AtProvider.LogConfiguration[0].LogSchemaVersion
			}

			if len(srcTyped.Status.AtProvider.LogConfiguration[0].S3Configuration) > 0 {
				if targetTyped.Status.AtProvider.LogConfiguration.S3Configuration == nil {
					targetTyped.Status.AtProvider.LogConfiguration.S3Configuration = &v1beta2.LogConfigurationS3ConfigurationObservation{}
				}

				if srcTyped.Status.AtProvider.LogConfiguration[0].S3Configuration[0].BucketName != nil {
					targetTyped.Status.AtProvider.LogConfiguration.S3Configuration.BucketName = srcTyped.Status.AtProvider.LogConfiguration[0].S3Configuration[0].BucketName
				}
				if srcTyped.Status.AtProvider.LogConfiguration[0].S3Configuration[0].Prefix != nil {
					targetTyped.Status.AtProvider.LogConfiguration.S3Configuration.Prefix = srcTyped.Status.AtProvider.LogConfiguration[0].S3Configuration[0].Prefix
				}
			}

		}
	}

	return nil
}

func templateConverterFromv1beta2Tov1beta1(src, target xpresource.Managed) error {
	srcTyped := src.(*v1beta2.ExperimentTemplate)
	targetTyped := target.(*v1beta1.ExperimentTemplate)

	if srcTyped.Spec.ForProvider.LogConfiguration != nil {
		if len(targetTyped.Spec.ForProvider.LogConfiguration) == 0 {
			targetTyped.Spec.ForProvider.LogConfiguration = []v1beta1.LogConfigurationParameters{{}}
		}

		if srcTyped.Spec.ForProvider.LogConfiguration.CloudwatchLogsConfiguration != nil {
			if len(targetTyped.Spec.ForProvider.LogConfiguration[0].CloudwatchLogsConfiguration) == 0 {
				targetTyped.Spec.ForProvider.LogConfiguration[0].CloudwatchLogsConfiguration = []v1beta1.CloudwatchLogsConfigurationParameters{{}}
			}
			if srcTyped.Spec.ForProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn != nil {
				targetTyped.Spec.ForProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn = srcTyped.Spec.ForProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn
			}
		}

		if srcTyped.Spec.ForProvider.LogConfiguration.LogSchemaVersion != nil {
			targetTyped.Spec.ForProvider.LogConfiguration[0].LogSchemaVersion = srcTyped.Spec.ForProvider.LogConfiguration.LogSchemaVersion
		}

		if srcTyped.Spec.ForProvider.LogConfiguration.S3Configuration != nil {
			if len(targetTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration) == 0 {
				targetTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration = []v1beta1.LogConfigurationS3ConfigurationParameters{{}}
			}

			if srcTyped.Spec.ForProvider.LogConfiguration.S3Configuration.BucketName != nil {
				targetTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration[0].BucketName = srcTyped.Spec.ForProvider.LogConfiguration.S3Configuration.BucketName
			}

			if srcTyped.Spec.ForProvider.LogConfiguration.S3Configuration.Prefix != nil {
				targetTyped.Spec.ForProvider.LogConfiguration[0].S3Configuration[0].Prefix = srcTyped.Spec.ForProvider.LogConfiguration.S3Configuration.Prefix
			}
		}
	}

	if srcTyped.Spec.InitProvider.LogConfiguration != nil {
		if len(targetTyped.Spec.InitProvider.LogConfiguration) == 0 {
			targetTyped.Spec.InitProvider.LogConfiguration = []v1beta1.LogConfigurationInitParameters{{}}
		}

		if srcTyped.Spec.InitProvider.LogConfiguration.CloudwatchLogsConfiguration != nil {
			if len(targetTyped.Spec.InitProvider.LogConfiguration[0].CloudwatchLogsConfiguration) == 0 {
				targetTyped.Spec.InitProvider.LogConfiguration[0].CloudwatchLogsConfiguration = []v1beta1.CloudwatchLogsConfigurationInitParameters{{}}
			}
			if srcTyped.Spec.InitProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn != nil {
				targetTyped.Spec.InitProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn = srcTyped.Spec.InitProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn
			}
		}

		if srcTyped.Spec.InitProvider.LogConfiguration.LogSchemaVersion != nil {
			targetTyped.Spec.InitProvider.LogConfiguration[0].LogSchemaVersion = srcTyped.Spec.InitProvider.LogConfiguration.LogSchemaVersion
		}

		if srcTyped.Spec.InitProvider.LogConfiguration.S3Configuration != nil {
			if len(targetTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration) == 0 {
				targetTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration = []v1beta1.LogConfigurationS3ConfigurationInitParameters{{}}
			}

			if srcTyped.Spec.InitProvider.LogConfiguration.S3Configuration.BucketName != nil {
				targetTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration[0].BucketName = srcTyped.Spec.InitProvider.LogConfiguration.S3Configuration.BucketName
			}

			if srcTyped.Spec.InitProvider.LogConfiguration.S3Configuration.Prefix != nil {
				targetTyped.Spec.InitProvider.LogConfiguration[0].S3Configuration[0].Prefix = srcTyped.Spec.InitProvider.LogConfiguration.S3Configuration.Prefix
			}
		}
	}

	if srcTyped.Status.AtProvider.LogConfiguration != nil {
		if len(targetTyped.Status.AtProvider.LogConfiguration) == 0 {
			targetTyped.Status.AtProvider.LogConfiguration = []v1beta1.LogConfigurationObservation{{}}
		}

		if srcTyped.Status.AtProvider.LogConfiguration.CloudwatchLogsConfiguration != nil {
			if len(targetTyped.Status.AtProvider.LogConfiguration[0].CloudwatchLogsConfiguration) == 0 {
				targetTyped.Status.AtProvider.LogConfiguration[0].CloudwatchLogsConfiguration = []v1beta1.CloudwatchLogsConfigurationObservation{{}}
			}
			if srcTyped.Status.AtProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn != nil {
				targetTyped.Status.AtProvider.LogConfiguration[0].CloudwatchLogsConfiguration[0].LogGroupArn = srcTyped.Status.AtProvider.LogConfiguration.CloudwatchLogsConfiguration.LogGroupArn
			}
		}

		if srcTyped.Status.AtProvider.LogConfiguration.LogSchemaVersion != nil {
			targetTyped.Status.AtProvider.LogConfiguration[0].LogSchemaVersion = srcTyped.Status.AtProvider.LogConfiguration.LogSchemaVersion
		}

		if srcTyped.Status.AtProvider.LogConfiguration.S3Configuration != nil {
			if len(targetTyped.Status.AtProvider.LogConfiguration[0].S3Configuration) == 0 {
				targetTyped.Status.AtProvider.LogConfiguration[0].S3Configuration = []v1beta1.LogConfigurationS3ConfigurationObservation{{}}
			}

			if srcTyped.Status.AtProvider.LogConfiguration.S3Configuration.BucketName != nil {
				targetTyped.Status.AtProvider.LogConfiguration[0].S3Configuration[0].BucketName = srcTyped.Status.AtProvider.LogConfiguration.S3Configuration.BucketName
			}
			if srcTyped.Status.AtProvider.LogConfiguration.S3Configuration.Prefix != nil {
				targetTyped.Status.AtProvider.LogConfiguration[0].S3Configuration[0].Prefix = srcTyped.Status.AtProvider.LogConfiguration.S3Configuration.Prefix
			}
		}
	}

	return nil
}
