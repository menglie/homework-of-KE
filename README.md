先安装go 的 编译环境
Windows/Centos安装GO语言环境
方法一：Centos下使用epel源安装：
yum install golang

方法二：Centos/Linux下源码安装golang：
wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
tar-C /usr/local-xzf go1.5.1.linux-amd64.tar.gz
mkdir$HOME/go
echo'export GOROOT=/usr/local/go' >> ~/.bashrc 
echo'export GOPATH=$HOME/go' >> ~/.bashrc 
echo'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc 
source$HOME/.bashrc 
安装go get工具：
yum install mercurial git bzr subversion

Windows下安装：
https://storage.googleapis.com/golang/go1.5.1.windows-386.zip

设置环境变量：
setx GOOS windows
setx GOARCH 386
setx GOROOT "D:\Program Files\go"
setx GOBIN "%GOROOT%\bin"
setxPATH%PATH%;D:\Program Files\go\bin"
