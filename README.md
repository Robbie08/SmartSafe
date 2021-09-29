# SmartSafe

### Getting Started
* To run this application you will need to have Go set up on your machine. 
  * You can [download Go by clicking on this link](https://golang.org/doc/tutorial/getting-started) and following the instruction. 
* Use the terminal to cd into `/go/src`.
  * (WARNING! This project needs to be in your GOPATH; If it isn't `/go/src` then use the appropriate destination.) 
* Next clone this project by using `git clone git@github.com:Robbie08/SmartSafe.git`

### To Build and Run:
The application contains a makefile so all the commands are taken care of. 
* To Build use the command `make build`
* To run you must use `make run` and the server should spin up on port 8080.
* If running on your machine you can type in `127.0.0.1:8080` into your browser but for the raspberry pi, ask me for the ip address.

### How to contribute :
1. Create a feature branch on your clone (This can be named anything)
   - `$ git checkout -b my-awesome-feature`
2. Do work and commit your changes
3. Ensure your code is bug free by testing it!
4. Make sure your master branch is up to date with master by rebasing
   - `$ git checkout master`
   - `$ git fetch `
   - `$ git rebase origin/master`
5. Rebase your feature branch
   - `$ git checkout my-awesome-feature`
   - `$ git rebase master`
6. Squash commits through interactive rebasing, use rebasing tutorials if needed: [article](https://thoughtbot.com/blog/git-interactive-rebase-squash-amend-rewriting-history)  or  [youtube video](https://www.youtube.com/watch?v=V5KrD7CmO4o)
7. Push your feature to your origin
    - `$ git push origin my-awesome-feature`
8. Make a pull request from feature branch to original master
