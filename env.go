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

// Package envtools provides helper functions for getting vars or secrets
// from environment variables.
// It uses logrus as logger.
package envtools

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.StandardLogger()

func SetLogger(newLogger *logrus.Logger) {
	logger = newLogger
}

// GetEnvOrWarn looks up the environment variable with the provided name.
// If the variable is set, its value is returned.
// Otherwise, a warning message will be logged.
func GetEnvOrWarn(envName string) string {
	val := os.Getenv(envName)
	if len(val) == 0 {
		logger.Warnf("environment variable '%v' is not set", envName)
	} else {
		logger.Infof("using configured value '%v' for '%v'", val, envName)
	}
	return val
}

// GetEnvSecretOrWarn looks up the environment variable with the provided name.
// If the variable is set, its value is returned.
// Otherwise, a warning message will be logged.
// The difference to GetEnvOrWarn is that the extracted value is masked by "*".
func GetEnvSecretOrWarn(envName string) string {
	val := os.Getenv(envName)
	if len(val) == 0 {
		logger.Warnf("environment variable '%v' is not set", envName)
	} else {
		logger.Infof("using configured secret '**********' for '%v'", envName)
	}
	return val
}

// GetEnvOrDefault looks up the environment variable with the provided name.
// If the variable is set, its value is returned.
// Otherwise, the provided defaultValue will be returned.
func GetEnvOrDefault(envName string, defaultValue string) string {
	val := os.Getenv(envName)
	if len(val) == 0 {
		logger.Infof(
			"environment variable '%v' is not set, defaulting to %v",
			envName,
			defaultValue,
		)
		return defaultValue
	}
	logger.Infof("using configured value '%v' for '%v'", val, envName)
	return val
}

// GetEnvOrFail looks up an environment variable. If the environment
// variable is not set or empty, an error is returned.
func GetEnvOrFail(envName string) (string, error) {
	val := os.Getenv(envName)
	if len(val) == 0 {
		msg := fmt.Sprintf(
			"please set the environment variable '%s'",
			envName,
		)
		logger.Errorln(msg)
		return "", fmt.Errorf(msg)
	}
	logger.Infof("using configured value '%v' for '%v'", val, envName)

	return val, nil
}

// GetEnvSecretOrFail looks up an environment variable. If the environment
// variable is not set or empty, an error is returned.
func GetEnvSecretOrFail(envName string) (string, error) {
	val := os.Getenv(envName)
	if len(val) == 0 {
		msg := fmt.Sprintf(
			"please set the environment variable '%s'",
			envName,
		)
		logger.Errorln(msg)
		return "", fmt.Errorf(msg)
	}
	logger.Infof("using configured secret '**********' for '%v'", envName)

	return val, nil
}

// GetEnvOrPanic looks up an environment variable. If the environment
// variable is not set, it panics.
func GetEnvOrPanic(envName string) string {
	value := os.Getenv(envName)
	if len(value) == 0 {
		msg := fmt.Sprintf("please set the environment variable '%s'", envName)
		logger.Panicln(msg)
		panic(msg)
	}
	logger.Infof("using configured value '%v' for '%v'", value, envName)

	return value
}

// GetEnvSecretOrPanic looks up an environment variable. If the environment
// variable is not set, it panics.
// The difference to GetEnvOrPanic is that the extracted value is masked by "*".
func GetEnvSecretOrPanic(envName string) string {
	value := os.Getenv(envName)
	if len(value) == 0 {
		msg := fmt.Sprintf("please set the environment variable '%s'", envName)
		logger.Panicln(msg)
		panic(msg)
	}
	logger.Infof("using configured secret '**********' for '%v'", envName)

	return value
}
