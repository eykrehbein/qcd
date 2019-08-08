~/go/bin/gox -osarch="darwin/amd64 linux/386 linux/amd64" -output="releases/{{.Dir}}_{{.OS}}_{{.Arch}}" ./src

rm -rf releases/darwin
rm -rf releases/linux_386
rm -rf releases/linux_amd64

mkdir releases/darwin
mkdir releases/linux_386
mkdir releases/linux_amd64


cp -R src/ releases/darwin/src
cp -R src/ releases/linux_386/src
cp -R src/ releases/linux_amd64/src

mkdir releases/darwin/bin
mkdir releases/linux_386/bin
mkdir releases/linux_amd64/bin

cp releases/src_darwin_amd64 releases/darwin/bin/qcdhelper
cp releases/src_linux_386 releases/linux_386/bin/qcdhelper
cp releases/src_linux_amd64 releases/linux_amd64/bin/qcdhelper

cp releases/darwin/src/qcdscript.sh releases/darwin/bin/qcdscript
cp releases/linux_386/src/qcdscript.sh releases/linux_386/bin/qcdscript
cp releases/linux_amd64/src/qcdscript.sh releases/linux_amd64/bin/qcdscript

cp setup.sh releases/darwin
cp setup.sh releases/linux_386
cp setup.sh releases/linux_amd64

cp update_dev.sh releases/darwin
cp update_dev.sh releases/linux_386
cp update_dev.sh releases/linux_amd64

cp build_dev.sh releases/darwin
cp build_dev.sh releases/linux_386
cp build_dev.sh releases/linux_amd64

rm releases/src_darwin_amd64
rm releases/src_linux_386
rm releases/src_linux_amd64

cd releases

zip -r qcd_darwin_MAC_OS.zip darwin
zip -r qcd_linux_386.zip linux_386
zip -r qcd_linux_amd64.zip linux_amd64

rm -rf darwin
rm -rf linux_386
rm -rf linux_amd64