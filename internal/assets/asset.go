package assets

import (
	"context"
	"errors"
	"fmt"

	"ashish.com/m/pkg/models"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	TestData = "test.data"
)

var isInitialized = false
var assetStore map[string]Asset

type Asset struct {
	Path    string
	Content []byte
}

// Init initializes static content store.
func Init(ctx context.Context, assets map[string][]byte) error {
	// skip if already initialized
	if isInitialized {
		return nil
	}

	// build temporary storage, so we can loop
	// it and store the content in asset store later
	tmpMap := map[string]string{
		TestData: "data/test.yaml",
	}

	// fill the asset store with path and content
	assetStore = make(map[string]Asset)
	for key, value := range tmpMap {
		content, ok := assets[value]
		if !ok {
			msg := fmt.Sprintf("asset file '%s' doesn't exist", value)
			fields := log.Fields{
				"asset.key":  key,
				"asset.path": value,
			}
			log.WithContext(ctx).WithFields(fields).Error(msg)
			return errors.New(msg)
		}
		assetStore[key] = Asset{
			Path:    value,
			Content: content,
		}
	}
	//log.WithContext(ctx).Info("Asset got initialized", assetStore)

	// Create a struct to hold the YAML data
	var pod models.Pod

	// Unmarshal the YAML data into the struct
	err := yaml.Unmarshal(assetStore[TestData].Content, &pod)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Print the data
	fmt.Println(pod)
	isInitialized = true
	return nil
}

// Get returns the specified asset.
func Get(key string) Asset {
	if asset, ok := assetStore[key]; ok {
		return asset
	}
	return Asset{}
}

// GetAll returns all assets.
func GetAll() map[string]Asset {
	return assetStore
}
