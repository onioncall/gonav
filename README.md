# GONAV 
---

This is a lightweight tool meant to speed up navigating directories in the terminal. 

## Commands
---

```
     s                       - focus in search bar
     spacebar, right arrow   - go into directory
     b, left arrow           - exit directory
     enter                   - cd into selected directory
     esc                     - if focused in search, focuses in directory box
     tab                     - toggles between search and directory box, can be used in place of s and esc depending where cursor focus is
     esc                     - if not focused in search, exits application without changing directory

     when searching for a directory, if there is only one result 'enter' will cd you to that directory and exit the application, and 'spacebar' will navigate you into that directory
```

https://github.com/user-attachments/assets/b809ae57-237a-4e68-b8f7-3b52dd36f2e4

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
   local dir=$(command gonav)
   if [ ! -d "$dir" ]; then
       echo "gonav: not a valid directory: $dir" >&2
       return 1
   fi
   
   cd "$dir"
}
```
