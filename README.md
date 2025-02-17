# pbkdf2Decode

 [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
 [![OS - Linux](https://img.shields.io/badge/OS-Linux-blue?logo=linux&logoColor=white)](https://www.linux.org/ "Go to Linux homepage")
 [![OS - FreeBSD](https://img.shields.io/badge/OS-FreeBSD-blue)](https://www.freebsd.org/ "Go to FreeBSD homepage")
 [![OS - MacOS](https://img.shields.io/badge/OS-macOS-blue?logo=Apple&logoColor=white)](https://apple.com/ "Go to Apple homepage")
 [![contributions - welcome](https://img.shields.io/badge/contributions-welcome-blue)](/CONTRIBUTING.md "Go to contributions doc")



## Roadmap:

- [ ] man pages
- [ ] DEB package
- [ ] RPM package
- [ ] MacOS PKG
- [ ] FreeBSD PKG

## Building from source

If you want to build yourself `pbkdf2Decode` from source, please verify to have already installed **go1.24.x** or higher.

Then run this command:

```bash
go build -v -ldflags="-X 'github.com/CyberPsyQuack/pbkdf2Decode/build.Version=$(cat VERSION)' -X 'github.com/CyberPsyQuack/pbkdf2Decode/build.BuildUser=Your Name' -X 'github.com/CyberPsyQuack/pbkdf2Decode/build.BuildTime=$(date)'"
```

## Manual Installation

### GNU/Linux

Remove any previous `pbkdf2Decode` installation by deleting the /usr/local/pbkdf2Decode folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh `pbkdf2Decode` tree in /usr/local/pbkdf2Decode

```bash
sudo tar -C /usr/local -zxf pbkdf2Decode-x86_64-linux.tar.gz
```
You can do this by adding the following line to your $HOME/.profile (or $HOME/.bashrc) or /etc/profile (for a system-wide installation)

```bash
export PATH=$PATH:/usr/local/pbkdf2Decode/bin
```
**Note**: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, run the following command (may differ if you are using .profile instead of .bashrc):

```bash
source $HOME/.bashrc
```