// Copyright (c) 2023 - for information on the respective copyright owner
// see the NOTICE file or the repository https://github.com/boschresearch/go-env-tools.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: goconst
package envtools

import (
	"bytes"
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

const envVarName = "SOME_ARBITRARY_TEST_ENV_VAR_NAME"
const expectedValue = "Not Empty"

// setupLoggingAndTearDown sets up a temporary buffer with the specified log level and a teardown
// function.
func setupLoggingAndTearDown(level logrus.Level) (*bytes.Buffer, func()) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	originalLevel := logrus.GetLevel()
	logrus.SetLevel(level)
	return &buf, func() {
		logrus.SetLevel(originalLevel)
		logrus.SetOutput(os.Stderr)
	}
}

func TestGetEnvOrWarn_SucceedsIfSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	t.Setenv(envVarName, expectedValue)

	actualValue := GetEnvOrWarn(envVarName)

	const expectedOutput = "using configured value 'Not Empty' for " +
		"'" + envVarName + "'"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvOrWarn_ReturnsEmptyStringIfEnvNotSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	expectedValue := ""
	t.Setenv(envVarName, expectedValue)
	err := os.Unsetenv(envVarName)
	assert.NoError(t, err)

	actualValue := GetEnvOrWarn(envVarName)

	const expectedOutput = "environment variable '" + envVarName + "' is not set"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvOrWarn_WarnsIfEmptyEnvVarSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	expectedValue := ""
	t.Setenv(envVarName, expectedValue)

	actualValue := GetEnvOrWarn(envVarName)

	const expectedOutput = "environment variable '" + envVarName + "' is not set"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvSecretOrWarn_SucceedsIfSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	t.Setenv(envVarName, expectedValue)

	actualValue := GetEnvSecretOrWarn(envVarName)

	const expectedOutput = "using configured secret '**********' for " +
		"'" + envVarName + "'"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvSecretOrWarn_ReturnsEmptyStringIfEnvNotSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	expectedValue := ""
	t.Setenv(envVarName, expectedValue)
	err := os.Unsetenv(envVarName)
	assert.NoError(t, err)

	actualValue := GetEnvSecretOrWarn(envVarName)

	const expectedOutput = "environment variable '" + envVarName + "' is not set"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvSecretOrWarn_ReturnsEmptyStringIfEmptyEnvVarSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	expectedValue := ""
	t.Setenv(envVarName, expectedValue)

	actualValue := GetEnvSecretOrWarn(envVarName)

	const expectedOutput = "environment variable '" + envVarName + "' is not set"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvOrDefault_SucceedsIfSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	t.Setenv(envVarName, expectedValue)

	actualValue := GetEnvOrDefault(envVarName, "Default Value")

	const expectedOutput = "using configured value 'Not Empty' for " +
		"'" + envVarName + "'"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvOrDefault_ReturnsDefaultIfSetToEmptyString(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	expectedValue := "Default Value"
	t.Setenv(envVarName, "")

	actualValue := GetEnvOrDefault(envVarName, expectedValue)

	const expectedOutput = "environment variable '" + envVarName + "' is not set," +
		" defaulting to Default Value"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvOrDefault_ReturnsDefaultIfEnvNotSet(t *testing.T) {
	buf, tearDownLogging := setupLoggingAndTearDown(logrus.InfoLevel)
	defer tearDownLogging()

	expectedValue := "Default Value"
	t.Setenv(envVarName, "")
	err := os.Unsetenv(envVarName)
	assert.NoError(t, err)

	actualValue := GetEnvOrDefault(envVarName, expectedValue)

	const expectedOutput = "environment variable '" + envVarName + "' is not set," +
		" defaulting to Default Value"
	assert.Equal(t, expectedValue, actualValue)
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestGetEnvOrFail_SucceedsIfEnvSet(t *testing.T) {
	t.Setenv(envVarName, expectedValue)

	actualValue, err := GetEnvOrFail(envVarName)

	assert.NoError(t, err)
	assert.Equal(t, expectedValue, actualValue)
}

func TestGetEnvOrFail_IndeedFailsIfEnvNotSet(t *testing.T) {
	t.Setenv(envVarName, "")
	err := os.Unsetenv(envVarName)
	assert.NoError(t, err)

	_, err = GetEnvOrFail(envVarName)
	assert.ErrorContains(t, err, "please set the environment variable '"+envVarName+"'")
}

func TestGetEnvOrFail_IndeedFailsIfEmptyEnvVarSet(t *testing.T) {
	t.Setenv(envVarName, "")
	_, err := GetEnvOrFail(envVarName)
	assert.ErrorContains(t, err, "please set the environment variable '"+envVarName+"'")
}

func TestGetEnvSecretOrFail_SucceedsIfEnvSet(t *testing.T) {
	t.Setenv(envVarName, expectedValue)

	actualValue, err := GetEnvSecretOrFail(envVarName)

	assert.NoError(t, err)
	assert.Equal(t, expectedValue, actualValue)
}

func TestGetEnvSecretOrFail_IndeedFailsIfEnvNotSet(t *testing.T) {
	t.Setenv(envVarName, "")
	err := os.Unsetenv(envVarName)
	assert.NoError(t, err)

	_, err = GetEnvSecretOrFail(envVarName)
	assert.ErrorContains(t, err, "please set the environment variable '"+envVarName+"'")
}

func TestGetEnvSecretOrFail_IndeedFailsIfEmptyEnvVarSet(t *testing.T) {
	t.Setenv(envVarName, "")
	_, err := GetEnvSecretOrFail(envVarName)
	assert.ErrorContains(t, err, "please set the environment variable '"+envVarName+"'")
}

func TestGetEnvOrPanic_PanicsIfUnset(t *testing.T) {
	t.Setenv(envVarName, "")
	err := os.Unsetenv(envVarName)
	assert.NoError(t, err)

	assert.Panics(t, func() {
		GetEnvOrPanic(envVarName)
	})
}

func TestGetEnvOrPanic_PanicsIfEmpty(t *testing.T) {
	t.Setenv(envVarName, "")

	assert.Panics(t, func() {
		GetEnvOrPanic(envVarName)
	})
}

func TestGetEnvOrPanic_SucceedsIfSet(t *testing.T) {
	t.Setenv(envVarName, expectedValue)

	actualValue := GetEnvOrPanic(envVarName)

	assert.Equal(t, expectedValue, actualValue)
}

func TestGetEnvSecretOrPanic_PanicsIfUnset(t *testing.T) {
	t.Setenv(envVarName, "")
	err := os.Unsetenv(envVarName)
	assert.NoError(t, err)

	assert.Panics(t, func() {
		GetEnvSecretOrPanic(envVarName)
	})
}

func TestGetEnvSecretOrPanic_PanicsIfEmpty(t *testing.T) {
	t.Setenv(envVarName, "")

	assert.Panics(t, func() {
		GetEnvSecretOrPanic(envVarName)
	})
}

func TestGetEnvSecretOrPanic_SucceedsIfSet(t *testing.T) {
	t.Setenv(envVarName, expectedValue)

	actualValue := GetEnvSecretOrPanic(envVarName)

	assert.Equal(t, expectedValue, actualValue)
}
