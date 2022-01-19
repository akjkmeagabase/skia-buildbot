package frontend

// targetDirectories is the set of directories for which this Gazelle extension will generate or
// update BUILD files.
//
// The value of this map indicates whether to recurse into the directory.
//
// TODO(lovisolo): Delete once we are targeting the entire repository.
var targetDirectories = map[string]bool{
	"am/modules":           true,
	"am/pages":             true,
	"bugs-central/modules": true,
	"bugs-central/pages":   true,
	"codesize/modules":     true,
	"codesize/pages":       true,
	"ct/modules":           true,
	"ct/pages":             true,
	"debugger-app/modules": true,
	"debugger-app/pages":   true,
	"demos/modules":        true,
	"demos/pages":          true,
	"fiddlek/modules":      true,
	"fiddlek/pages":        true,
	"golden/modules":       true,
	"golden/pages":         true,
	"hashtag/modules":      true,
	"hashtag/pages":        true,
	"infra-sk/modules":     true,
	"jsfiddle/modules":     true,
	"jsfiddle/pages":       true,
	"leasing/modules":      true,
	"leasing/pages":        true,
	"machine/modules":      true,
	"machine/pages":        true,
	"modules/devices":      true,
	"new_element/modules":  true,
	"particles/modules":    true,
	"particles/pages":      true,
	"perf/modules":         true,
	"perf/pages":           true,
	"puppeteer-tests":      true,
	"scrap/modules":        true,
	"scrap/pages":          true,
	"shaders/modules":      true,
	"shaders/pages":        true,
	"skcq/modules":         true,
	"skcq/pages":           true,
	"skottie/modules":      true,
	"skottie/pages":        true,
	"task_driver/modules":  true,
	"task_driver/pages":    true,
	"tree_status/modules":  true,
	"tree_status/pages":    true,
}
