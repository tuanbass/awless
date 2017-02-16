/* Copyright 2017 WALLIX

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

// DO NOT EDIT
// This file was automatically generated with go generate
package aws

import "github.com/wallix/awless/template"

var AWSTemplatesDefinitions = map[string]template.TemplateDefinition{
	"createvpc": template.TemplateDefinition{
		Action:         "create",
		Entity:         "vpc",
		Api:            "ec2",
		RequiredParams: []string{"cidr"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletevpc": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "vpc",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createsubnet": template.TemplateDefinition{
		Action:         "create",
		Entity:         "subnet",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "vpc"},
		ExtraParams:    []string{"zone"},
		TagsMapping:    []string{},
	},
	"updatesubnet": template.TemplateDefinition{
		Action:         "update",
		Entity:         "subnet",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{"public-vms"},
		TagsMapping:    []string{},
	},
	"deletesubnet": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "subnet",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createinstance": template.TemplateDefinition{
		Action:         "create",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"image", "type", "count", "count", "subnet"},
		ExtraParams:    []string{"lock", "key", "ip", "group", "userdata"},
		TagsMapping:    []string{"name"},
	},
	"updateinstance": template.TemplateDefinition{
		Action:         "update",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{"lock", "group", "type"},
		TagsMapping:    []string{},
	},
	"deleteinstance": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"startinstance": template.TemplateDefinition{
		Action:         "start",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"stopinstance": template.TemplateDefinition{
		Action:         "stop",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"checkinstance": template.TemplateDefinition{
		Action:         "check",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"id", "state", "timeout"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createsecuritygroup": template.TemplateDefinition{
		Action:         "create",
		Entity:         "securitygroup",
		Api:            "ec2",
		RequiredParams: []string{"description", "name", "vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"updatesecuritygroup": template.TemplateDefinition{
		Action:         "update",
		Entity:         "securitygroup",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "id", "protocol"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletesecuritygroup": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "securitygroup",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createvolume": template.TemplateDefinition{
		Action:         "create",
		Entity:         "volume",
		Api:            "ec2",
		RequiredParams: []string{"zone", "size"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletevolume": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "volume",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachvolume": template.TemplateDefinition{
		Action:         "attach",
		Entity:         "volume",
		Api:            "ec2",
		RequiredParams: []string{"device", "instance", "id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createinternetgateway": template.TemplateDefinition{
		Action:         "create",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteinternetgateway": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachinternetgateway": template.TemplateDefinition{
		Action:         "attach",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{"id", "vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"detachinternetgateway": template.TemplateDefinition{
		Action:         "detach",
		Entity:         "internetgateway",
		Api:            "ec2",
		RequiredParams: []string{"id", "vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createroutetable": template.TemplateDefinition{
		Action:         "create",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"vpc"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteroutetable": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachroutetable": template.TemplateDefinition{
		Action:         "attach",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"id", "subnet"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"detachroutetable": template.TemplateDefinition{
		Action:         "detach",
		Entity:         "routetable",
		Api:            "ec2",
		RequiredParams: []string{"association"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createroute": template.TemplateDefinition{
		Action:         "create",
		Entity:         "route",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "gateway", "table"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteroute": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "route",
		Api:            "ec2",
		RequiredParams: []string{"cidr", "table"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createtags": template.TemplateDefinition{
		Action:         "create",
		Entity:         "tags",
		Api:            "ec2",
		RequiredParams: []string{"resource"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createkeypair": template.TemplateDefinition{
		Action:         "create",
		Entity:         "keypair",
		Api:            "ec2",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletekeypair": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "keypair",
		Api:            "ec2",
		RequiredParams: []string{"id"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createuser": template.TemplateDefinition{
		Action:         "create",
		Entity:         "user",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deleteuser": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "user",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"creategroup": template.TemplateDefinition{
		Action:         "create",
		Entity:         "group",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletegroup": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "group",
		Api:            "iam",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"attachpolicy": template.TemplateDefinition{
		Action:         "attach",
		Entity:         "policy",
		Api:            "iam",
		RequiredParams: []string{"arn", "user"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"detachpolicy": template.TemplateDefinition{
		Action:         "detach",
		Entity:         "policy",
		Api:            "iam",
		RequiredParams: []string{"arn", "user"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createbucket": template.TemplateDefinition{
		Action:         "create",
		Entity:         "bucket",
		Api:            "s3",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"deletebucket": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "bucket",
		Api:            "s3",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
	"createstorageobject": template.TemplateDefinition{
		Action:         "create",
		Entity:         "storageobject",
		Api:            "s3",
		RequiredParams: []string{"file", "bucket"},
		ExtraParams:    []string{"name"},
		TagsMapping:    []string{},
	},
	"deletestorageobject": template.TemplateDefinition{
		Action:         "delete",
		Entity:         "storageobject",
		Api:            "s3",
		RequiredParams: []string{"bucket", "key"},
		ExtraParams:    []string{},
		TagsMapping:    []string{},
	},
}
