/*
Copyright 2019 The Kubernetes Authors.

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

package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
)

//go test ./*.go -v util_test.go
/*
=== RUN   TestFindSuggestionByErrorMessage
--- PASS: TestFindSuggestionByErrorMessage (0.00s)
=== RUN   TestFormat
--- PASS: TestFormat (0.00s)
=== RUN   TestMount
--- PASS: TestMount (0.00s)
=== RUN   TestCmdValid
    util_test.go:100:
        	Error Trace:	util_test.go:100
        	Error:      	Expected nil, but got: &errors.errorString{s:"Command echo is not support."}
        	Test:       	TestCmdValid
    util_test.go:101:
        	Error Trace:	util_test.go:101
        	Error:      	Expected nil, but got: &errors.errorString{s:"Command echo abc >> /mnt/abc is regexp match, args:>>"}
        	Test:       	TestCmdValid
    util_test.go:104:
        	Error Trace:	util_test.go:104
        	Error:      	Expected nil, but got: &errors.errorString{s:"Command cd is not support."}
        	Test:       	TestCmdValid
    util_test.go:108:
        	Error Trace:	util_test.go:108
        	Error:      	Expected nil, but got: &errors.errorString{s:"Command ping is not support."}
        	Test:       	TestCmdValid
    util_test.go:109:
        	Error Trace:	util_test.go:109
        	Error:      	Expected nil, but got: &errors.errorString{s:"Command ping && echo abc is regexp match, args:&&"}
        	Test:       	TestCmdValid
    util_test.go:113:
        	Error Trace:	util_test.go:113
        	Error:      	Expected nil, but got: &errors.errorString{s:"Command chmod -R 755 /mnt/../abc || echo abc is regexp match, args:||"}
        	Test:       	TestCmdValid
    util_test.go:117:
        	Error Trace:	util_test.go:117
        	Error:      	Expected nil, but got: &errors.errorString{s:"Command chmod -R 755 /mnt/abc; echo abc is regexp match, args:/mnt/abc;"}
        	Test:       	TestCmdValid
--- FAIL: TestCmdValid (0.00s)
FAIL
FAIL	command-line-arguments	0.552s
FAIL
*/
func TestCmdValid(t *testing.T) {
	//cmdSet = hashset.New("mount", "lctl", "umount", "nsenter", "findmnt", "chmod", "dd", "mkfs.ext4")
	/*cmd := "mount -t nfs -o vers=3,nolock,tcp,noresvport 138df4a7e1-xgi36.cn-beijing.nas.aliyuncs.com:/ /mnt"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "umount /var/lib/kubelet/pods/1867bbd7-539d-49ff-a14b-3856a143b8e6/volume-subpaths/app-log/vinayaprocess/0"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	configCmd := [...]string{"lctl set_param osc.*.max_rpcs_in_flight=64", "lctl set_param osc.*.max_pages_per_rpc=256", "lctl set_param mdc.*.max_rpcs_in_flight=64", "lctl set_param mdc.*.max_mod_rpcs_in_flight=64"}
	for _, element := range configCmd {
		assert.Nil(t, CheckCmd(cmd, strings.Split(element, " ")[0]))
		assert.Nil(t, CheckCmdArgs(cmd, strings.Split(element, " ")[1:]...))
	}

	cmd = "nsenter --mount=/proc/1/ns/mnt yum localinstall -y /etc/csi-tool/aliyun-alinas-utils-1.1-3.al7.noarch.rpm "
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "findmnt -o SOURCE --noheadings /mnt"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "chmod -R 755 /mnt/abc"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "dd if=/dev/zero of=/mnt/abc bs=4k seek=4k count=0"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "mkfs.ext4 -E packed_meta_blocks=1,nodiscard,lazy_itable_init=0 -O ^has_journal -F /mnt/abc"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "dd if=/dev/zero of=../abc bs=4k seek=4k count=0"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "echo abc >> /mnt/abc"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "cd .."
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "ping && echo abc"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "chmod -R 755 /mnt/abc || echo abc"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))

	cmd = "chmod -R 755 /mnt/abc; echo abc"
	assert.Nil(t, CheckCmd(cmd, strings.Split(cmd, " ")[0]))
	assert.Nil(t, CheckCmdArgs(cmd, strings.Split(cmd, " ")[1:]...))*/
}

func TestGetServiceType(t *testing.T) {
	defer func() { klog.OsExit = os.Exit }()
	klog.OsExit = func(c int) { panic(c) }
	tests := []struct {
		name                 string
		runAsController      bool
		runControllerService bool
		runNodeService       bool
		serviceTypeEnv       string
		want                 ServiceType
		fatal                bool
	}{
		{
			name:                 "default",
			runControllerService: true,
			runNodeService:       true,
			want:                 Controller | Node,
		},
		{
			name:            "Run as controller",
			runAsController: true,
			want:            Controller,
		},
		{
			name:           "env provisioner",
			serviceTypeEnv: ProvisionerService,
			want:           Controller,
		},
		{
			name:           "env plugin",
			serviceTypeEnv: PluginService,
			want:           Node,
		},
		{
			name:                 "Run controller",
			runControllerService: true,
			runNodeService:       false,
			want:                 Controller,
		},
		{
			name:                 "Run node",
			runControllerService: false,
			runNodeService:       true,
			want:                 Node,
		},
		{
			name:                 "nothing",
			runControllerService: false,
			runNodeService:       false,
			want:                 0,
		},
		{
			name:           "invalid env",
			serviceTypeEnv: "invalid",
			fatal:          true,
		},
		{
			name:            "conflict env and run-as-controller",
			runAsController: true,
			serviceTypeEnv:  PluginService,
			fatal:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SERVICE_TYPE", tt.serviceTypeEnv)
			if tt.fatal {
				assert.Panics(t, func() {
					GetServiceType(tt.runAsController, tt.runControllerService, tt.runNodeService)
				})
				return
			}
			got := GetServiceType(tt.runAsController, tt.runControllerService, tt.runNodeService)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsDirTmpfs(t *testing.T) {
	mounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{
			Path: "/",
		},
		{
			Path: "/some/tmpfs",
			Type: "tmpfs",
		},
		{
			Path: "/some/other",
			Type: "ext4",
		},
	})
	isTmpfs, err := IsDirTmpfs(mounter, "/some/tmpfs")
	assert.Nil(t, err)
	assert.True(t, isTmpfs)
}

func TestIsDirTmpfsFalse(t *testing.T) {
	mounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{
			Path: "/",
		},
		{
			Path: "/some/other",
			Type: "ext4",
		},
	})
	isTmpfs, err := IsDirTmpfs(mounter, "/some/tmpfs")
	assert.Nil(t, err)
	assert.False(t, isTmpfs)
}
