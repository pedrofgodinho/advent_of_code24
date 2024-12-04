set windows-powershell := true

default: build

fmt:
    go fmt ./...

vet: fmt
    go vet ./...

build: vet
    go build

run DAY='0': build
    ./advent_of_code24 --day {{DAY}}

clean:
    just _clean-{{os()}}

_clean-linux:
    rm -f advent_of_code24

_clean-macos:
    rm -f advent_of_code24

_clean-windows:
    if (Test-Path -Path 'advent_of_code24.exe') { \
      Remove-Item -Path 'advent_of_code24.exe' -Force \
    }