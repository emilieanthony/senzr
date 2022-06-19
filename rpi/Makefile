include ../tools/golangci-lint/rules.mk

package = github.com/emilieanthony/senzr/rpi
out = build/rpi

pico_software_version = 2.1.21
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
	curl -X POST -F 'somdn_download_key=MTYzMTQyOTQxOA==' -F 'action=somdn_download_single_variation' -F 'somdn_product=2747' -F 'somdn_variation=5645' https://www.dracal.com/en/product/qtenki/ --output ${pico_software_name}.tar.gz
	tar xvzf ${pico_software_name}.tar.gz -C tools/
	rm ${pico_software_name}.tar.gz
	cd tools/${pico_software_name}-${pico_software_version}/client && make
	sudo cp tools/${pico_software_name}-${pico_software_version}/client/usbtenkiget /usr/local/bin/
	sudo chmod 777 /usr/local/bin/usbtenkiget
	sudo cp tools/${pico_software_name}-${pico_software_version}/scripts/99-usbtenki.rules /etc/udev/rules.d/
	sudo udevadm control --reload-rules

install:
	go install ${package}

test:
	go test ./...

run:
	nohup build/rpi &

lint: golangci-lint

build: build-rpi install
