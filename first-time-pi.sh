#!/bin/bash

echo "Executing first time pi"

echo 'export PS1="\[\e[31m\]\u\[\e[m\]\[\e[36m\]@\[\e[m\]\[\e[32m\]\h\[\e[m\]: \[\e[35m\]\w\[\e[m\] \$ "' >>~/.bash_profile

cat >>~/.bash_profile <<'EOF'
export GOPATH=$HOME/go
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
EOF

echo 'alias ls="ls --color=auto"' >>~/.bash_profile
echo 'alias grep="grep --color=auto"' >>~/.bash_profile
echo 'alias fgrep="fgrep --color=auto"' >>~/.bash_profile
echo 'alias egrep="egrep --color=auto"' >>~/.bash_profile

echo "Done"
