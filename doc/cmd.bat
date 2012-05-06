@set git_install_root=N:\Git
@set GOROOT=N:\Go
@set GOBIN=N:\Go\bin
@set PATH=%git_install_root%\bin;%git_install_root%\mingw\bin;%git_install_root%\cmd;%GOBIN%;%PATH%
@set GOPATH=N:\c--2llvm

@set PLINK_PROTOCOL=ssh
@set HOME=N:\
declare -x GIT_SSH="N:\.ssh"
@cd %HOME%
@start %COMPSEC%