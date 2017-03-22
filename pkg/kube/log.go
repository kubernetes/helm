/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kube

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

var logger *log.Entry

func init() {
	if level := os.Getenv("KUBE_LOG_LEVEL"); level != "" {
		flag.Set("vmodule", fmt.Sprintf("loader=%s,round_trippers=%s,request=%s", level, level, level))
		flag.Set("logtostderr", "true")
	}
	logger = log.WithFields(log.Fields{
		"_package": "kube",
	})
}
