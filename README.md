## Install air for development
### go install github.com/air-verse/air@latest

## Install tailwindcss globally
### bun install -g tailwindcss@latest

## Start the hot reload server part
### air

## Start the hot reload html templ
### templ generate --watch --proxy=http://localhost:8000
#### change 8000 with the port of the app in .env

## Develop with tailwindcss
### Install tailwindcss: bun install -D tailwindcss
### To watch the css change use cmd: tailwindcss -i views/css/app.css -o public/styles.css --watch
