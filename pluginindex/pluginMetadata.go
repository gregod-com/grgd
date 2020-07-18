package pluginindex

// PluginMetadataImpl ...
type PluginMetadataImpl struct {
	Name     string
	Version  string
	URL      string
	Category string
	Active   bool
	Path     string
}

// GetFullname ...
func (plmeta PluginMetadataImpl) GetFullname() string {
	return plmeta.Category + "-" + plmeta.Name + "-" + plmeta.Version
}

// GetName ...
func (plmeta PluginMetadataImpl) GetName() string {
	return plmeta.Name
}

// GetVersion ...
func (plmeta PluginMetadataImpl) GetVersion() string {
	return plmeta.Version
}

// GetURL ...
func (plmeta PluginMetadataImpl) GetURL() string {
	return plmeta.URL
}

// GetCategory ...
func (plmeta PluginMetadataImpl) GetCategory() string {
	return plmeta.Category
}

// GetActive ...
func (plmeta PluginMetadataImpl) GetActive() bool {
	return plmeta.Active
}

// GetPath ...
func (plmeta PluginMetadataImpl) GetPath() string {
	return plmeta.Path
}

// SetActive ...
func (plmeta *PluginMetadataImpl) SetActive(a bool) {
	plmeta.Active = a
}

// ToggleActive ...
func (plmeta *PluginMetadataImpl) ToggleActive() {
	plmeta.Active = !plmeta.Active
}

// SetPath ...
func (plmeta *PluginMetadataImpl) SetPath(path string) {
	plmeta.Path = path
}
