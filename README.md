# logrus-logzio-hook

[![Build Status](https://travis-ci.com/form3tech-oss/logrus-logzio-hook.svg?branch=master)](https://travis-ci.com/form3tech-oss/logrus-logzio-hook)

A Logrus hook for shipping logs to Logz.io.

## Usage

```go
package main

import (
    "github.com/logzio/logzio-go"
    "github.com/sirupsen/logrus"
    
    "github.com/form3tech-oss/logrus-logzio-hook/pkg/hook"
)

func main() { 
    // Create the Logz.io client.
    s, err := logzio.New("<TOKEN>")
    if err != nil {
        logrus.Fatal(err)
    }
    // Create the Logz.io hook for Logrus.
    h := hook.NewLogzioHook(s)
    defer h.Stop()
    // Add the Logz.io hook to the standard logger.
    logrus.AddHook(h)
    // Happy logging.
    logrus.Info("I'll show up in Logz.io.")
}
```

## License

Copyright 2019 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
