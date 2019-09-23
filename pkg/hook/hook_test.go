// Copyright 2019 Form3 Financial Cloud
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

package hook_test

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/logzio/logzio-go"
	"github.com/sirupsen/logrus"

	"github.com/form3tech-oss/logrus-logzio-hook/pkg/hook"
)

const (
	drainInterval = 5 * time.Second
)

var (
	token = os.Getenv("LOGZIO_TOKEN")
	url   = os.Getenv("LOGZIO_URL")
)

func TestHook(t *testing.T) {
	// Create a buffer to hold the debugging output for the Logz.io client.
	b := bytes.Buffer{}
	// Create the Logz.io client.
	s, err := logzio.New(token, logzio.SetDebug(&b), logzio.SetDrainDuration(drainInterval), logzio.SetUrl(url))
	if err != nil {
		logrus.Fatal(err)
	}
	// Create the Logz.io hook for Logrus.
	h := hook.NewLogzioHook(s)
	defer h.Stop()

	// Add the Logz.io hook to the standard logger.
	logrus.AddHook(h)

	// Log something.
	logrus.Info("logruz-logzio-hook")

	// Wait for twice the drain interval and make sure we've got no errors in the debug buffer.
	time.Sleep(2 * drainInterval)
	d := b.String()
	if strings.Contains(d, "error") || strings.Contains(d, "failed") {
		t.Fatalf("Failed to ship logs to Logz.io:\n%s", safeDump(d))
	} else {
		t.Logf("Logs successfully shipped to Logz.io:\n%s", safeDump(d))
	}
}

func safeDump(v string) string {
	return strings.ReplaceAll(v, token, "*****")
}
