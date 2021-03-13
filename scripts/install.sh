#!/usr/bin/env bash

set -e

VERSION=$(curl -s https://api.github.com/repos/avakarev/dotfiles-cli/releases/latest | grep -oP '"tag_name": "\K(.*)(?=")')

OS="$(uname -s)"
ARCH="$(uname -m)"

case $OS in
    "Linux")
        case $ARCH in
        "x86_64")
            ARCH=amd64
            ;;
        "aarch64")
            ARCH=arm64
            ;;
        "armv6l")
            ARCH=armv6
            ;;
        "armv7l")
            ARCH=armv7
            ;;
        "armv8")
            ARCH=arm64
            ;;
        .*386.*)
            echo "Error: 386 arch is not supported"
            exit 1
            ;;
        esac
        PLATFORM="linux_$ARCH"
    ;;
    "Darwin")
        echo "Macos detected. Use Homebrew on macos instead: \"brew tap avakarev/tap && brew install dotfiles\""
        exit 1
    ;;
esac

if [ -z "$PLATFORM" ]; then
    echo "Error: operating system $OS is not supported"
    exit 1
fi

if [ -f /etc/os-release ]; then
    DISTRO=$(awk -F= '$1=="ID_LIKE" { print $2 ;}' /etc/os-release)
fi

if [ -z "$DISTRO" ]; then
    echo "Error: distro is not detected"
    exit 1
fi

TEMP_DIRECTORY=$(mktemp -d)

download() {
    PACKAGE="dotfiles_${VERSION}_${PLATFORM}.${1}"
    URL="https://github.com/avakarev/dotfiles-cli/releases/download/$VERSION/$PACKAGE"
    curl -Ls -o $TEMP_DIRECTORY/$PACKAGE $URL
    if [ $? -ne 0 ]; then
        echo "Download failed! Exiting."
        exit 1
    fi
    echo $TEMP_DIRECTORY/$PACKAGE
}

install_deb() {
    echo "Installing dotfiles $VERSION..."
    echo "Downloading deb package..."
    _pkg=$(download "deb")
    echo "Installing deb package..."
    echo "$_pkg"
    sudo apt install "$_pkg"
    rm -f $_pkg
}

install_rpm() {
    echo "Installing dotfiles $VERSION..."
    echo "Downloading rpm package..."
    _pkg=$(download "rpm")
    echo "Installing rpm package..."
    echo "$_pkg"
    sudo yum localinstall "$_pkg"
    rm -f $_pkg
}

case $DISTRO in
    *debian*|*ubuntu*)
        install_deb
    ;;
    *rhel*|*centos*|*fedora*)
        install_rpm
    ;;
    *)
    echo "Error: \"$DISTRO\" distro is not supported"
    exit 1
  ;;
esac

echo "Done!"
