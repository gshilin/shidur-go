package config

import (
  "github.com/gshilin/train"
  "html/template"
  "os"
  "net/http"
  "strings"
)

// The AssetHeaders struct is used to set Cache-Control headers to all GET and HEAD
// requests to /assets in production. Because these assets have digested names, we
// can set the cache time really high, and use this app as origin for a CDN.
type AssetHeaders struct {
}

func NewAssetHeaders() *AssetHeaders {
  return &AssetHeaders{}
}

func (s *AssetHeaders) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  // Ignore everything not in /assets
  path := r.URL.Path

  if strings.HasPrefix(path, "/assets") {
    // Set asset caching to one year
    rw.Header().Set("Cache-Control", "public, max-age=31536000")
    train.ServeRequest(rw, r)
    return
  }

  // Ignore if not production
  if os.Getenv("GO_ENV") != "production" {
    next(rw, r)
    return
  }

  // Ignore all but GET and HEAD
  if r.Method != "GET" && r.Method != "HEAD" {
    next(rw, r)
    return
  }

  next(rw, r)
}


// These helper functions are passed to the Go templates.
// You can add more easily by adding more functions to the FuncMap.
// The asset_path function reads from the manifest.json file to return
// the path to the digested assets in production or non-digested assets in
// development.
func AssetHelpers(root string) template.FuncMap {

  // Return digested asset paths in production
  if train.IsInProduction() {
    return template.FuncMap{
      "asset_path": func(asset string) string {
        return "/assets/" + train.ManifestInfo[asset]
      },
      "javascript_tag":            train.JavascriptTag,
      "stylesheet_tag":            train.StylesheetTag,
      "stylesheet_tag_with_param": train.StylesheetTagWithParam,
    }

    // Return non-digested asset paths in other envs
  } else {
    return template.FuncMap{
      "asset_path": func(asset string) string {
        return "/assets/" + asset
      },
      "javascript_tag":            train.JavascriptTag,
      "stylesheet_tag":            train.StylesheetTag,
      "stylesheet_tag_with_param": train.StylesheetTagWithParam,
    }
  }
}
