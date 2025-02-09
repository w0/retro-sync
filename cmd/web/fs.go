package main

import "path"

func (app *application) PathBuilder(resource string, systemId string, filename string) string {
	return path.Join(app.rootPath, resource, systemId, filename)
}
