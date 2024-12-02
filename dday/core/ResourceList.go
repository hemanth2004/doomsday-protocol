package core

type ResourceList struct {
	DefaultResources []Resource
	CustomResources  []Resource
}

func (r ResourceList) FindResourceFromKey(nameOfRow string) (Resource, error) {

	for _, resource := range r.DefaultResources {
		if resource.Name == nameOfRow {
			return resource, nil
		}
	}

	for _, resource := range r.CustomResources {
		if resource.Name == nameOfRow {
			return resource, nil
		}
	}

	return Resource{}, nil
}
