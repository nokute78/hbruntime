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

import ()

// https://github.com/opencontainers/runtime-spec/blob/master/config.md

type Root struct {
	Path     string `json:"path"`
	ReadOnly bool   `json:"readonly,omitempty"`
}

type Mount struct {
	Destination string   `json:"destination"`
	Source      string   `json:"source,omitempty"`
	Options     []string `json:"options,omitempty"`
	// POSIX-platform Mounts
	Type string `json:"type,omitempty"`
}

type ConsoleSize struct {
	Height uint `json:"height"`
	Width  uint `json:"width"`
}

// For POSIX process
type RLimit struct {
	Type string `json:"type"`
	Soft uint64 `json:"soft"`
	Hard uint64 `json:"hard"`
}

// For Linux
type Capabilities struct {
	Effective   []string `json:"effective,omitempty"`
	Bounding    []string `json:"bounding,omitempty"`
	Inheritable []string `json:"inheritable,omitempty"`
	Permitted   []string `json:"permitted,omitempty"`
	Ambient     []string `json:"ambient,omitempty"`
}

// TODO
type User struct {
	// Windows
	UserName string `json:"username,omitempty"`
	// POSIX-Platform
	Uid            int   `json:"uid"`
	Gid            int   `json:"gid"`
	Umask          int   `json:"umask,omitempty"`
	AdditionalGids []int `json:"additionalGids,omitempty"`
}

type Process struct {
	Terminal    bool        `json:"terminal,omitempty"`
	ConsoleSize ConsoleSize `json:"consolesize,omitempty"`
	Cwd         string      `json:"cwd"`
	Env         []string    `json:"env,omitempty"`
	Args        []string    `json:"args,omitempty"`
	CommandLine string      `json:"commandLine,omitempty"`
	// POSIX Process
	RLimits []RLimit `json:"rlimits,omitempty"`
	// Linux Process
	ApparmorProfile string       `json:"apparmorProfile,omitempty"`
	Capabilities    Capabilities `json:"capabilities,omitempty"`
	NoNewPrivileges bool         `json:"noNewPrivileges,omitempty"`
	OomScoreAdj     int          `json:"oomScoreAdj,omitempty"`
	SelinuxLabel    string       `json:"selinuxLabel,omitempty"`
	User            User         `json:"user,omitempty"`
}

type HookEntry struct {
	Path    string   `json:"path"`
	Args    []string `json:"args,omitempty"`
	Env     []string `json:"env,omitempty"`
	Timeout int      `json:"timeout,omitempty"`
}

type CreateRuntime struct {
	HookEntry
}

type CreateContainer struct {
	HookEntry
}

type StartContainer struct {
	HookEntry
}

type PostStart struct {
	HookEntry
}

type PostStop struct {
	HookEntry
}

type Hooks struct {
	PreStart        []interface{}     `json:"prestart,omitempty"` // Deprecated
	CreateRuntime   []CreateRuntime   `json:"createRuntime,omitempty"`
	CreateContainer []CreateContainer `json:"createContainer,omitempty"`
	StartContainer  []StartContainer  `json:"startContainer,omitempty"`
	PostStart       []PostStart       `json:"poststart,omitempty"`
	PostStop        []PostStop        `json:"poststop,omitempty"`
}

type RuntimeConfig struct {
	OciVersion  string                 `json:"ociVersion"`
	Root        Root                   `json:"root,omitempty"`
	Mounts      []Mount                `json:"mounts,omitempty"`
	Process     Process                `json:"process,omitempty"`
	Hooks       Hooks                  `json:"hooks,omitempty"`
	Annotations map[string]interface{} `json:"annotations,omitempty"`
}
