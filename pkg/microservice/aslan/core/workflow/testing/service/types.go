/*
Copyright 2022 The KodeRover Authors.

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

package service

import (
	commonmodels "github.com/koderover/zadig/pkg/microservice/aslan/core/common/repository/models"
	"github.com/koderover/zadig/pkg/types"
)

const DefaultScanningTimeout = 60 * 60

type Scanning struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	ProjectName string              `json:"project_name"`
	Description string              `json:"description"`
	ScannerType string              `json:"scanner_type"`
	ImageID     string              `json:"image_id"`
	SonarID     string              `json:"sonar_id"`
	Repos       []*types.Repository `json:"repos"`
	// Parameter is for sonarQube type only
	Parameter string `json:"parameter"`
	// Script is for other type only
	Script          string                         `json:"script"`
	AdvancedSetting *types.ScanningAdvancedSetting `json:"advanced_settings"`
}

type ListScanningRespItem struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Statistics *ScanningStatistic `json:"statistics"`
	CreatedAt  int64              `json:"created_at"`
	UpdatedAt  int64              `json:"updated_at"`
}

type ScanningRepoInfo struct {
	CodehostID int    `json:"codehost_id"`
	Source     string `json:"source"`
	RepoOwner  string `json:"repo_owner"`
	RepoName   string `json:"repo_name"`
	PR         int    `json:"pr"`
	Branch     string `json:"branch"`
}

type ScanningStatistic struct {
	TimesRun       int64 `json:"times_run"`
	AverageRuntime int64 `json:"run_time_average"`
}

type ListScanningTaskResp struct {
	ScanInfo   *ScanningInfo       `json:"scan_info"`
	ScanTasks  []*ScanningTaskResp `json:"scan_tasks"`
	TotalTasks int                 `json:"total_tasks"`
}

type ScanningInfo struct {
	Editor    string `json:"editor"`
	UpdatedAt int64  `json:"updated_at"`
}

type ScanningTaskResp struct {
	ScanID    int64  `json:"scan_id"`
	Status    string `json:"status"`
	RunTime   int64  `json:"run_time"`
	Creator   string `json:"creator"`
	CreatedAt int64  `json:"created_at"`
}

type ScanningTaskDetail struct {
	Creator    string              `json:"creator"`
	Status     string              `json:"status"`
	CreateTime int64               `json:"create_time"`
	EndTime    int64               `json:"end_time"`
	RepoInfo   []*types.Repository `json:"repo_info"`
	ResultLink string              `json:"result_link,omitempty"`
}

func ConvertToDBScanningModule(args *Scanning) *commonmodels.Scanning {
	// ID is omitted since they are of different type and there will be no use of it
	return &commonmodels.Scanning{
		Name:            args.Name,
		ProjectName:     args.ProjectName,
		Description:     args.Description,
		ScannerType:     args.ScannerType,
		ImageID:         args.ImageID,
		SonarID:         args.SonarID,
		Repos:           args.Repos,
		Parameter:       args.Parameter,
		Script:          args.Script,
		AdvancedSetting: args.AdvancedSetting,
	}
}

func ConvertDBScanningModule(scanning *commonmodels.Scanning) *Scanning {
	return &Scanning{
		ID:              scanning.ID.Hex(),
		Name:            scanning.Name,
		ProjectName:     scanning.ProjectName,
		Description:     scanning.Description,
		ScannerType:     scanning.ScannerType,
		ImageID:         scanning.ImageID,
		SonarID:         scanning.SonarID,
		Repos:           scanning.Repos,
		Parameter:       scanning.Parameter,
		Script:          scanning.Script,
		AdvancedSetting: scanning.AdvancedSetting,
	}
}
