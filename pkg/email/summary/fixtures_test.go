package summary_test

import (
	"go.uber.org/zap"

	"github.com/skbki/policy-reporter/pkg/crd/client/clientset/versioned/fake"
	v1alpha2client "github.com/skbki/policy-reporter/pkg/crd/client/clientset/versioned/typed/policyreport/v1alpha2"
	"github.com/skbki/policy-reporter/pkg/email"
	"github.com/skbki/policy-reporter/pkg/validate"
)

var (
	filter = email.NewFilter(validate.RuleSets{}, validate.RuleSets{})
	logger = zap.NewNop()
)

func NewFakeClient() (v1alpha2client.Wgpolicyk8sV1alpha2Interface, v1alpha2client.PolicyReportInterface, v1alpha2client.ClusterPolicyReportInterface) {
	client := fake.NewSimpleClientset().Wgpolicyk8sV1alpha2()

	return client, client.PolicyReports("test"), client.ClusterPolicyReports()
}
