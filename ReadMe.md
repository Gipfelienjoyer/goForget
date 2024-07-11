# GoForget

Hello there

GoForget is a ToDo app written in Go and is one of my first Go projects.

I tried to make the code as simple as possible for me but if you have any improvements fell free to contribute :)

I also made a documentation for everyone that also wants to make an ToDo app and needs inspiration.




### Dokumentation

#### Main.go
In Main.go we declared some important Variables:
1. ToDoStorage: This variable is used as a temporary storage that gets initially doesn't have a value when starting the application, but the gets loaded in later.
2. NextID: This is just used for the add command to know what's the last ID is to append is (The last item + 1 of ToDoStorage)

### Add.go
In Add.go is (who would have guessed) everything that has to do with adding a task to our app.
