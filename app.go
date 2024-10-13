package main

import (
	"context"
	"log"
)

// App struct
type App struct {
	ctx    context.Context
	config *Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) LoadState() (bool, error) {
	var err error
	a.config, err = CreateOrReadConfigFile()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
		return false, err
	}

	isConfigEmpty := a.config.GitHubToken == "" || a.config.TableName == "" || a.config.Owner == "" || a.config.Repo == ""
	return isConfigEmpty, nil
}

// GetConfig returns the current config to the frontend
func (a *App) GetConfig() (*Config, error) {
	return a.config, nil
}

// SaveConfig saves the updated config from the frontend
func (a *App) SaveConfig(newConfig *Config) (*Config, error) {
	err := WriteConfig(newConfig)
	if err != nil {
		return nil, err
	}
	a.config = newConfig
	return a.config, nil
}

type VersionedData struct {
	Releases []Release
	Version  string
}

func stripVersion(version string) string {
	return version[1:]
}

func prepareAllData(releases []Release, version string) VersionedData {
	for i, release := range releases {
		if stripVersion(release.TagName) == version {
			releases[i].Current = true
		}
	}
	return VersionedData{Releases: releases, Version: version}
}

func (a *App) LoadData() (VersionedData, error) {
	prepareClient(a.ctx)
	version, err := a.GetVersion()
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
	}
	releases, err := getReleases(a.config)
	if err != nil {
		log.Fatalf("Error getting releases: %v", err)
	}

	return prepareAllData(releases, version), nil
}

func (a *App) DeployVersion(version string) (string, error) {
	err := a.UpdateVersion(a.ctx, version)
	if err != nil {
		log.Fatalf("Error updating version: %v", err)
	}

	newVersion, err := a.GetVersion()
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
		return "", err
	}

	return newVersion, nil
}
