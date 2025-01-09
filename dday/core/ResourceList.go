package core

import (
	"math"
)

type ResourceList struct {
	DefaultResources []Resource
	CustomResources  []Resource
}

func (r *ResourceList) PauseAllResources() {
	for i := range r.DefaultResources {
		r.DefaultResources[i].PauseResource()
	}
	for i := range r.CustomResources {
		r.CustomResources[i].PauseResource()
	}
}

func (r *ResourceList) ResumeAllResources() {
	for i := range r.DefaultResources {
		r.DefaultResources[i].ResumeResource()
	}
	for i := range r.CustomResources {
		r.CustomResources[i].ResumeResource()
	}
}

func (r *ResourceList) GetOverallProgress() float64 {
	var totalSize, downloadedSize float64

	// Process default resources
	for _, resource := range r.DefaultResources {
		if resource.Info.Size > 0 {
			totalSize += math.Abs(resource.Info.Size)
			downloadedSize += math.Abs(resource.Info.Done)
		}
	}

	// Process custom resources
	for _, resource := range r.CustomResources {
		if resource.Info.Size > 0 {
			totalSize += math.Abs(resource.Info.Size)
			downloadedSize += math.Abs(resource.Info.Done)
		}
	}

	if totalSize > 0 {
		return downloadedSize / totalSize
	}
	return 0.0
}

func (r *ResourceList) GetCoreProgress() float64 {
	var totalSize, downloadedSize float64

	// Process default resources
	for _, resource := range r.DefaultResources {
		if resource.Tier == 0 {
			if resource.Info.Size > 0 {
				totalSize += math.Abs(resource.Info.Size)
				downloadedSize += math.Abs(resource.Info.Done)
			}
		}
	}

	if totalSize > 0 {
		return downloadedSize / totalSize
	}
	return 0.0
}
