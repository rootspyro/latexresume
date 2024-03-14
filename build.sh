#! /bin/bash
#
binName=latexresume
builds=dist
version=$(git describe --tags --abbrev=0)

BuildBinary() {
  # $1 = arch
  # $2 = os 

  GOARCH=$1 GOOS=$2 go build -ldflags "-X main.version=$version"  -o $binName .
  local tarName=$binName\_$2\_$1.tar.gz

  tar -czvf $tarName $binName
  mv $tarName $builds/$tarName && rm $binName
}

# BUILD'S DIR
if [ ! -d $builds ]; then
  echo "creating $builds directory"
  mkdir $builds
fi

# OS CONFIGURATION

osList=("linux" "windows" "darwin")

linuxArchs=("arm64" "386" "amd64")
windowsArchs=("386" "amd64")
macOsArchs=("amd64")

declare -A os_archs

os_archs["linux"]=${linuxArchs[@]}
os_archs["windows"]=${windowsArchs[@]}
os_archs["darwin"]=${macOsArchs[@]}

# BUILD THE BINARIES
for os in ${osList[@]}
do
  osArchs="${os_archs[$os]}"

  for arch in ${osArchs[@]} 
  do
    BuildBinary $arch $os
  done
done
