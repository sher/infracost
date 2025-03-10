package main_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/infracost/infracost/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestOutputHelp(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--help"}, nil)
}

func TestOutputFormatHTML(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "html", "--path", "./testdata/example_out.json", "--path", "./testdata/azure_firewall_out.json"}, nil)
}

func TestOutputFormatJSON(t *testing.T) {
	opts := DefaultOptions()
	opts.IsJSON = true
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "json", "--path", "./testdata/example_out.json", "--path", "./testdata/azure_firewall_out.json"}, opts)
}

func TestOutputFormatGitHubComment(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "github-comment", "--path", "./testdata/example_out.json", "--path", "./testdata/terraform_v0.14_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatGitHubCommentMultipleSkipped(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "github-comment", "--path", "./testdata/example_out.json", "--path", "./testdata/terraform_v0.14_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatGitHubCommentNoChange(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "github-comment", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatGitLabComment(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "gitlab-comment", "--path", "./testdata/example_out.json", "--path", "./testdata/terraform_v0.14_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatGitLabCommentMultipleSkipped(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "gitlab-comment", "--path", "./testdata/example_out.json", "--path", "./testdata/terraform_v0.14_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatGitLabCommentNoChange(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "gitlab-comment", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatSlackMessage(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "slack-message", "--path", "./testdata/example_out.json", "--path", "./testdata/terraform_v0.14_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatSlackMessageMultipleSkipped(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "slack-message", "--path", "./testdata/example_out.json", "--path", "./testdata/terraform_v0.14_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatSlackMessageNoChange(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "slack-message", "--path", "./testdata/terraform_v0.14_nochange_breakdown.json"}, nil)
}

func TestOutputFormatTable(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--format", "table", "--path", "./testdata/example_out.json", "--path", "./testdata/azure_firewall_out.json"}, nil)
}

func TestOutputTerraformFieldsAll(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"output", "--path", "./testdata/example_out.json", "--path", "./testdata/azure_firewall_out.json", "--fields", "all"}, nil)
}

func TestOutputTerraformOutFileHTML(t *testing.T) {
	testdataName := testutil.CalcGoldenFileTestdataDirName()
	goldenFilePath := "./testdata/" + testdataName + "/infracost_output.golden"
	outputPath := filepath.Join(t.TempDir(), "infracost_output.html")

	GoldenFileCommandTest(t, testdataName, []string{"output", "--path", "./testdata/example_out.json", "--format", "html", "--out-file", outputPath}, nil)

	actual, err := ioutil.ReadFile(outputPath)
	require.Nil(t, err)
	actual = stripDynamicValues(actual)

	testutil.AssertGoldenFile(t, goldenFilePath, actual)
}

func TestOutputTerraformOutFileJSON(t *testing.T) {
	testdataName := testutil.CalcGoldenFileTestdataDirName()
	goldenFilePath := "./testdata/" + testdataName + "/infracost_output.golden"
	outputPath := filepath.Join(t.TempDir(), "infracost_output.json")

	GoldenFileCommandTest(t, testdataName, []string{"output", "--path", "./testdata/example_out.json", "--format", "json", "--out-file", outputPath}, nil)

	actual, err := ioutil.ReadFile(outputPath)
	require.Nil(t, err)
	actual = stripDynamicValues(actual)

	testutil.AssertGoldenFile(t, goldenFilePath, actual)
}

func TestOutputTerraformOutFileTable(t *testing.T) {
	testdataName := testutil.CalcGoldenFileTestdataDirName()
	goldenFilePath := "./testdata/" + testdataName + "/infracost_output.golden"
	outputPath := filepath.Join(t.TempDir(), "infracost_output.txt")

	GoldenFileCommandTest(t, testdataName, []string{"output", "--path", "./testdata/example_out.json", "--out-file", outputPath}, nil)

	actual, err := ioutil.ReadFile(outputPath)
	require.Nil(t, err)
	actual = stripDynamicValues(actual)

	testutil.AssertGoldenFile(t, goldenFilePath, actual)
}
