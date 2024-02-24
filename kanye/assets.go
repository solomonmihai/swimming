package kanye

import (
	"errors"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AssetType string

const (
	ASSET_TEXTURE AssetType = "texture"
	ASSET_MODEL AssetType = "model"
	ASSET_SOUND AssetType = "sound"
)

type AssetDesc struct {
	Name, Path string
	AssetType  AssetType
}

type AssetsManager struct {
	textures map[string]rl.Texture2D
	models map[string]rl.Model
	sounds map[string]rl.Sound
}

func NewAssetsManager() *AssetsManager {
	return &AssetsManager{
		textures: make(map[string]rl.Texture2D),
		models: make(map[string]rl.Model),
		sounds: make(map[string]rl.Sound),
	}
}

func (am *AssetsManager) Texture(name string) *rl.Texture2D {
	img, found := am.textures[name]

	if !found {
		panic(fmt.Sprintf("[assets] texture not found: %s", name))
	}

	return &img
}

func (am *AssetsManager) Model(name string) *rl.Model {
	model, found := am.models[name]
	if !found {
		panic(fmt.Sprintf("[assets] model not found: %s", name))
	}

	return &model
}

func (am *AssetsManager) LoadAssets(descriptions []AssetDesc) {
  fmt.Printf("\n--- loading assets ---\n\n")

	for _, desc := range descriptions {
		checkAssetExists(desc)

		switch desc.AssetType {
		case ASSET_TEXTURE:
			am.textures[desc.Name] = am.loadTexture(desc)
			break
		case ASSET_MODEL:
			am.models[desc.Name] = am.loadModel(desc)
			break
		case ASSET_SOUND:
			am.sounds[desc.Name] = am.loadSound(desc)
			break
		}

    fmt.Println(fmt.Sprintf("[assets] loaded %ss[ \"%s\" ]: %s", desc.AssetType, desc.Name, desc.Path))
	}

  fmt.Printf("\n ----------------------\n")
}

func checkAssetExists(desc AssetDesc) {
	if _, err := os.Stat(desc.Path); errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("[assets] %s %s does not exist", desc.AssetType, desc.Path))
	}
}

func (am *AssetsManager) loadTexture(desc AssetDesc) rl.Texture2D {
	if _, found := am.textures[desc.Name]; found {
		panic(fmt.Sprintf("[assets] texture already exists: %s", desc.Name))
	}

  return rl.LoadTexture(desc.Path)
}

func (am *AssetsManager) loadModel(desc AssetDesc) rl.Model {
	if _, found := am.models[desc.Name]; found {
		panic(fmt.Sprintf("[assets] model already exists: %s", desc.Name))
	}

  return rl.LoadModel(desc.Path)
}

func (am *AssetsManager) loadSound(desc AssetDesc) rl.Sound {
	if _, found := am.sounds[desc.Name]; found {
		panic(fmt.Sprintf("[assets] sound already exists: %s", desc.Name))
	}

  return rl.LoadSound(desc.Path)
}

func (am *AssetsManager) Close() {
	for _, texture := range am.textures {
		rl.UnloadTexture(texture)
	}

	for _, model := range am.models {
		rl.UnloadModel(model)
	}

	for _, sound := range am.sounds {
		rl.UnloadSound(sound)
	}
}
