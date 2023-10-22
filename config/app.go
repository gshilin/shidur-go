package config

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gshilin/train"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// Struct to hold main variables for this application.
// Routes all have access to an instance of this struct.
type App struct {
	Negroni *negroni.Negroni
	Router  *mux.Router
	Render  *render.Render
	QRender *render.Render
	DB      *gorm.DB
}

// This function is called from main.go and from the tests
// to create a new application.
func NewApp(root string) *App {

	CheckEnv()

	// Use negroni for middleware
	ne := negroni.New()

	// Use gorilla/mux for routing
	ro := mux.NewRouter()

	// Use Render for template. Pass in path to templates folder
	// as well as asset helper functions.
	re := render.New(render.Options{
		Directory:  filepath.Join(root, "templates"),
		Layout:     "layouts/layout",
		Extensions: []string{".html"},
		Funcs: []template.FuncMap{
			AssetHelpers(root),
		},
	})
	qre := render.New(render.Options{
		Directory:  filepath.Join(root, "templates"),
		Layout:     "layouts/message",
		Extensions: []string{".html"},
		Funcs: []template.FuncMap{
			AssetHelpers(root),
		},
	})

	// Establish connection to DB as specificed in database.go
	db := NewDB()

	// Add middleware to the stack
	ne.Use(negroni.NewRecovery())
	ne.Use(negroni.NewLogger())
	ne.Use(NewAssetHeaders())
	ne.Use(negroni.NewStatic(http.Dir("public")))
	ne.UseHandler(ro)

	train.Config.SASS.DebugInfo = true
	train.Config.SASS.LineNumbers = true
	train.Config.Verbose = true
	train.Config.BundleAssets = true
	//ZZZtrain.ConfigureHttpHandler(ro)

	// Return a new App struct with all these things.
	return &App{ne, ro, re, qre, db}
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Setup(ro *mux.Router, routes Routes) {
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		ro.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
