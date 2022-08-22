package aws

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/smithy-go/transport/http"
	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog/log"
	"go.mondoo.io/mondoo/resources/library/jobpool"
	aws_transport "go.mondoo.io/mondoo/motor/providers/aws"
	"go.mondoo.io/mondoo/resources/packs/aws/awspolicy"
	"go.mondoo.io/mondoo/resources/packs/core"
)

func (l *mqlAwsLambda) id() (string, error) {
	return "aws.lambda", nil
}

func (l *mqlAwsLambda) GetFunctions() ([]interface{}, error) {
	at, err := awstransport(l.MotorRuntime.Motor.Provider)
	if err != nil {
		return nil, err
	}
	res := []interface{}{}
	poolOfJobs := jobpool.CreatePool(l.getFunctions(at), 5)
	poolOfJobs.Run()

	// check for errors
	if poolOfJobs.HasErrors() {
		return nil, poolOfJobs.GetErrors()
	}
	// get all the results
	for i := range poolOfJobs.Jobs {
		res = append(res, poolOfJobs.Jobs[i].Result.([]interface{})...)
	}

	return res, nil
}

func (l *mqlAwsLambda) getFunctions(at *aws_transport.Provider) []*jobpool.Job {
	tasks := make([]*jobpool.Job, 0)
	regions, err := at.GetRegions()
	if err != nil {
		return []*jobpool.Job{{Err: err}}
	}

	for _, region := range regions {
		regionVal := region
		f := func() (jobpool.JobResult, error) {
			log.Debug().Msgf("calling aws with region %s", regionVal)

			svc := at.Lambda(regionVal)
			ctx := context.Background()
			res := []interface{}{}

			var marker *string
			for {
				functionsResp, err := svc.ListFunctions(ctx, &lambda.ListFunctionsInput{Marker: marker})
				if err != nil {
					return nil, errors.Wrap(err, "could not gather aws lambda functions")
				}
				for _, function := range functionsResp.Functions {
					vpcConfigJson, err := core.JsonToDict(function.VpcConfig)
					if err != nil {
						return nil, err
					}
					var dlqTarget string
					if function.DeadLetterConfig != nil {
						dlqTarget = core.ToString(function.DeadLetterConfig.TargetArn)
					}
					tags := make(map[string]interface{})
					tagsResp, err := svc.ListTags(ctx, &lambda.ListTagsInput{Resource: function.FunctionArn})
					if err == nil {
						for k, v := range tagsResp.Tags {
							tags[k] = v
						}
					}
					mqlFunc, err := l.MotorRuntime.CreateResource("aws.lambda.function",
						"arn", core.ToString(function.FunctionArn),
						"name", core.ToString(function.FunctionName),
						"dlqTargetArn", dlqTarget,
						"vpcConfig", vpcConfigJson,
						"region", regionVal,
						"tags", tags,
					)
					if err != nil {
						return nil, err
					}
					res = append(res, mqlFunc)
				}
				if functionsResp.NextMarker == nil {
					break
				}
				marker = functionsResp.NextMarker
			}
			return jobpool.JobResult(res), nil
		}
		tasks = append(tasks, jobpool.NewJob(f))
	}
	return tasks
}

func (l *mqlAwsLambdaFunction) GetConcurrency() (int64, error) {
	funcName, err := l.Name()
	if err != nil {
		return 0, err
	}
	region, err := l.Region()
	if err != nil {
		return 0, err
	}
	at, err := awstransport(l.MotorRuntime.Motor.Provider)
	if err != nil {
		return 0, err
	}
	svc := at.Lambda(region)
	ctx := context.Background()

	// no pagination required
	functionConcurrency, err := svc.GetFunctionConcurrency(ctx, &lambda.GetFunctionConcurrencyInput{FunctionName: &funcName})
	if err != nil {
		return 0, errors.Wrap(err, "could not gather aws lambda function concurrency")
	}
	if functionConcurrency.ReservedConcurrentExecutions != nil {
		return core.ToInt64From32(functionConcurrency.ReservedConcurrentExecutions), nil
	}

	return 0, nil
}

func (l *mqlAwsLambdaFunction) GetPolicy() (interface{}, error) {
	funcArn, err := l.Arn()
	if err != nil {
		return nil, err
	}
	region, err := l.Region()
	if err != nil {
		return 0, err
	}
	at, err := awstransport(l.MotorRuntime.Motor.Provider)
	if err != nil {
		return nil, err
	}
	svc := at.Lambda(region)
	ctx := context.Background()

	// no pagination required
	functionPolicy, err := svc.GetPolicy(ctx, &lambda.GetPolicyInput{FunctionName: &funcArn})
	var respErr *http.ResponseError
	if err != nil && errors.As(err, &respErr) {
		if respErr.HTTPStatusCode() == 404 {
			return nil, nil
		}
	} else if err != nil {
		return nil, err
	}
	if functionPolicy != nil {
		var policy lambdaPolicyDocument
		err = json.Unmarshal([]byte(*functionPolicy.Policy), &policy)
		if err != nil {
			return nil, err
		}
		return core.JsonToDict(policy)
	}

	return nil, nil
}

func (l *mqlAwsLambdaFunction) id() (string, error) {
	return l.Arn()
}

type lambdaPolicyDocument struct {
	Version   string                  `json:"Version,omitempty"`
	Statement []lambdaPolicyStatement `json:"Statement,omitempty"`
}

type lambdaPolicyStatement struct {
	Sid       string              `json:"Sid,omitempty"`
	Effect    string              `json:"Effect,omitempty"`
	Action    string              `json:"Action,omitempty"`
	Resource  string              `json:"Resource,omitempty"`
	Principal awspolicy.Principal `json:"Principal,omitempty"`
}
