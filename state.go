/*
   Copyright 2020 Takahiro Yamashita

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

package main

import (
	"fmt"
)

type State struct {
	OciVersion string
	Id         string
	Status     string
	Pid        int
	Bundle     string
	//	Annotations map[interface()]interface()
}

const (
	StatusCreating = "creating"
	StatusCreated  = "created"
	StatusRunning  = "running"
	StatusStopped  = "stopped"
)

func (s State) String() string {
	return fmt.Sprintf("ociversion:\"%s\", id:\"%s\", status:\"%s\", pid:%d, bundle:\"%s\"", s.OciVersion, s.Id, s.Status, s.Pid, s.Bundle)
}
