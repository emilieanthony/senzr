
package = github.com/emilieanthony/senzr/rpi
out = build/rpi

pico_software_version = 2.1.22
pico_software_name = usbtenki

develop:
	go run main.go

build-mac:
	env GOOS=darwin GOARCH=arm64 go build -o ${out} ${package}

build-rpi:
	env GOOS=linux GOARCH=arm go build -o ${out} ${package}

install-pico-deps:
	sudo apt-get install libusb-1.0-0.dev

install-pico: install-pico-deps
	mkdir -p tools
	curl https://www.dracal.com/store/downloads/${pico_software_name}-${pico_software_version}.tar.gz --output ${pico_software_name}.tar.gz
	tar xvzf ${pico_software_name}.tar.gz -C tools/
	rm ${pico_software_name}.tar.gz
	cd tools/${pico_software_name}-${pico_software_version}/client && make
	sudo cp tools/${pico_software_name}-${pico_software_version}/client/usbtenkiget /usr/local/bin/

install:
	go install ${package}
