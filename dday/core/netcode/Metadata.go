package netcode

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DownloadMeta represents a single download's metadata
type DownloadMeta struct {
	Filename        string    `json:"filename"`
	URL             string    `json:"url"`
	TotalSize       int64     `json:"total_size"`
	BytesDownloaded int64     `json:"bytes_downloaded"`
	SupportsResume  bool      `json:"supports_resume"` // Supports HTTP Range Requests
	LastModified    time.Time `json:"last_modified"`
	ETag            string    `json:"etag"`
	Status          string    `json:"status"`
	Location        string    `json:"location"`
	PartialPath     string    `json:"partial_path"` // path to .part file
}

// MetaStore represents the structure of our meta.json file
type MetaStore struct {
	Downloads map[string]*DownloadMeta `json:"downloads"` // key is a unique download ID
}

// MetaManager handles all metadata operations
type MetaManager struct {
	filepath string
	mu       sync.RWMutex
	store    MetaStore
}

// NewMetaManager creates a new metadata manager
func NewMetaManager(metaPath string) (*MetaManager, error) {
	mm := &MetaManager{
		filepath: metaPath,
		store: MetaStore{
			Downloads: make(map[string]*DownloadMeta),
		},
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(metaPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// Load existing metadata if it exists
	if err := mm.load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return mm, nil
}

// load reads the metadata file into memory
func (mm *MetaManager) load() error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	data, err := os.ReadFile(mm.filepath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &mm.store)
}

// save writes the current metadata to disk
func (mm *MetaManager) save() error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	data, err := json.MarshalIndent(mm.store, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(mm.filepath, data, 0644)
}

// GetDownload retrieves metadata for a specific download
func (mm *MetaManager) GetDownload(id string) (*DownloadMeta, bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	meta, exists := mm.store.Downloads[id]
	return meta, exists
}

// UpdateDownload updates or creates a download's metadata
func (mm *MetaManager) UpdateDownload(id string, meta *DownloadMeta) error {
	mm.mu.Lock()
	mm.store.Downloads[id] = meta
	mm.mu.Unlock()

	return mm.save()
}

// UpdateProgress updates just the progress of a download
func (mm *MetaManager) UpdateProgress(id string, bytesDownloaded int64) error {
	mm.mu.Lock()
	if meta, exists := mm.store.Downloads[id]; exists {
		meta.BytesDownloaded = bytesDownloaded
		meta.LastModified = time.Now()
	}
	mm.mu.Unlock()

	return mm.save()
}

// DeleteDownload removes a download's metadata
func (mm *MetaManager) DeleteDownload(id string) error {
	mm.mu.Lock()
	delete(mm.store.Downloads, id)
	mm.mu.Unlock()

	return mm.save()
}
