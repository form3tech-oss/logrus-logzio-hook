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

package hook

import (
	"time"

	"github.com/logzio/logzio-go"
	"github.com/sirupsen/logrus"
)

// LogzioHook is a Logz.io hook for Logrus.
type LogzioHook struct {
	f logrus.Formatter
	s *logzio.LogzioSender
}

// NewLogzioHook returns a new instance of the Logz.io hook for Logrus which uses the specified sender.
func NewLogzioHook(s *logzio.LogzioSender) *LogzioHook {
	return &LogzioHook{
		s: s,
		f: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "@timestamp",
				logrus.FieldKeyMsg:  "message",
			},
		},
	}
}

// Stop stops the Logz.io sender.
func (h *LogzioHook) Stop() {
	h.s.Stop()
}

// Fire writes the specified entry to Logz.io.
func (h *LogzioHook) Fire(e *logrus.Entry) error {
	v, err := h.f.Format(e)
	if err != nil {
		return err
	}
	return h.s.Send(v)
}

// Levels returns the set of log levels for which the Logz.io hook is active.
func (h *LogzioHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
