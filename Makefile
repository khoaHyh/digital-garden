live/templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v

live/server:
	go run github.com/air-verse/air@v1.61.7 \
		--build.cmd "go build -o tmp/main ./cmd/" \
		--build.bin "./tmp/main" \
		--build.delay "100" \
		--build.stop_on_error "false" \
		--color.build "yellow" \
		--color.main "magenta" \
		--color.runner "green" \
		--color.watcher "cyan" \
		--log.main_only "false" \
		--log.time "false" \
		--misc.clean_on_exit "true"

# TODO: minify in prod
live/tailwind:
	./tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

live: 
	make -j3 live/templ live/server live/tailwind
