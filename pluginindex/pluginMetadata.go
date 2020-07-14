package pluginindex

// PluginMetadataImpl ...
type PluginMetadataImpl struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	Size     uint64 `yaml:"a,omitempty"`
	URL      string `yaml:"a,omitempty"`
	Category string `yaml:"a,omitempty"`
	Active   bool   `yaml:"a,omitempty"`
}

// GetName ...
func (plmeta PluginMetadataImpl) GetName() string {
	return plmeta.Category
}

// GetVersion ...
func (plmeta PluginMetadataImpl) GetVersion() string {
	return plmeta.Version
}

// GetSize ...
func (plmeta PluginMetadataImpl) GetSize() uint64 {
	return plmeta.Size
}

// GetURL ...
func (plmeta PluginMetadataImpl) GetURL() string {
	return plmeta.URL
}

// GetCategory ...
func (plmeta PluginMetadataImpl) GetCategory() string {
	return plmeta.Category
}

// SetName ...
func (plmeta PluginMetadataImpl) SetName(name string) {
	plmeta.Name = name
}

// SetVersion ...
func (plmeta PluginMetadataImpl) SetVersion(version string) {
	plmeta.Version = version
}

// SetSize ...
func (plmeta PluginMetadataImpl) SetSize(size uint64) {
	plmeta.Size = size
}

// SetURL ...
func (plmeta PluginMetadataImpl) SetURL(url string) {
	plmeta.URL = url
}

// SetCategory ...
func (plmeta PluginMetadataImpl) SetCategory(category string) {
	plmeta.Category = category
}
