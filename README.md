# GOJO - starting your go journey

Welcome to `go`!

Start your journey by reading Rob Pike's [_Go at Google: Language Design in the Service of Software Engineering_](https://talks.golang.org/2012/splash.article) which outlines the design principles behind the language.

## Let's code!

### On a mac:

0. Install go:

        brew install go
    
0. `go` works best when all your `go` code is in a single tree, so make the place where you're going to keep your code:

        mkdir -p ~/src/go/src
        
0. Set your `$GOPATH` so you can `go get` without trouble:

        export GOPATH=$HOME/src/go
        
0. For extra awesome, add your go bin path to your `$PATH`:

        export PATH=$PATH:$GOPATH/bin

0. Clone this repo:

        cd $GOPATH/src
        git clone git@github.com:akjones/gojo.git 

0. Install the test framework:

        go get github.com/smartystreets/goconvey
        
0. Start the test runner:

        cd ~/src/go/src/gojo
        # If you put your go bins on your $PATH
        goconvey

        # Othewise
        $GOPATH/bin/goconvey
        
0. Visit http://localhost:8080 and enable html notifications.
        
0. Fix the broken test.
