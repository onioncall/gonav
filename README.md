# GONAV 
---

This is a lightweight tool meant to speed up navigating directories in the terminal. 

## Commands
---

s                       - Focus in search Bar
spacebar, right arrow   - go into directory
b, left arrow           - exit directory
enter                   - cd into selected directory
esc                     - if focused in search, focuses in directory box
esc                     - if not focused in search, exits application without changing directory

https://github.com/user-attachments/assets/57319784-13fe-41f9-a527-882d7cfbabce


## Setup
--- 

move the binary (gonav file) and add it to your bin
if you'd like to build it yourself
```
sudo go build -o /usr/local/bin/gonav .
```

otherwise
```
sudo cp gonav /usr/local/bin/gonav 
```

add the following to your shell rc file (.bashrc, .zshrc, etc)
```
gonav() {
     cd $(command gonav)
}
```
There isn't error handling for this yet, I'm working on it. 
