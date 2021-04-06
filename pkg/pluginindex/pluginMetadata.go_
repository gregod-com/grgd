package pluginindex

// PluginMetadataImpl ...
type PluginMetadataImpl struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	URL      string `yaml:"url"`
	Category string `yaml:"category"`
	Active   bool   `yaml:"active"`
	Loaded   bool   `yaml:"loaded"`
	Path     string `yaml:"path"`
}

// GetFullname ...
func (plmeta PluginMetadataImpl) GetFullname() string {
	return plmeta.Category + "-" + plmeta.Name + "-" + plmeta.Version
}

// GetIdentifier ...
func (plmeta PluginMetadataImpl) GetIdentifier() string {
	return plmeta.Category + "-" + plmeta.Name
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

// GetLoaded ...
func (plmeta PluginMetadataImpl) GetLoaded() bool {
	return plmeta.Loaded
}

// SetLoaded ...
func (plmeta *PluginMetadataImpl) SetLoaded(l bool) {
	plmeta.Loaded = l
}

// SetActive ...
func (plmeta *PluginMetadataImpl) SetActive(a bool) {
	plmeta.Active = a
}

// ToggleActive ...
func (plmeta *PluginMetadataImpl) ToggleActive() bool {
	plmeta.Active = !plmeta.Active
	return plmeta.Active
}

// SetPath ...
func (plmeta *PluginMetadataImpl) SetPath(path string) {
	plmeta.Path = path
}
