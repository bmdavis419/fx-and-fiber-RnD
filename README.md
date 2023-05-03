# Fiber and fx R&D

### IMPORTANT DOCUMENTATION LINKS

- [Fiber](https://docs.gofiber.io/)
- [fx **READ THIS FIRST**](https://uber-go.github.io/fx/get-started/)

### What is this?

This is basically just me playing around with Fiber and fx, working to understand what fx is, how to use it, and how it works with my favorite web framework, Fiber. This is not a great example but it illustrates how to use fx with Fiber, and a few key concepts.

### Concepts shown here

1. *The Application Lifecycle:* this project shows off how the application lifecycle works with fx. For more information go to the [fx lifecycle docs](https://uber-go.github.io/fx/lifecycle.html). The TLDR of it is that within the main function ```NewStorage``` and ```NewDemoHandler``` are provided. This means that their ```OnStart``` methods are called, then their returned value is available for use elsewhere (and we use their result within the app itself). After that the app is invoked (which means the function actually runs) and it uses the values from the previous functions.

2. *Provided Functions are Called in Order:* notice how when you run the app it print out that storage has been started before demo handler has been started. This is because they are passed in in that order. Switch their order in the main function and you'll see the opposite.

3. *We can pass stuff around like crazy:* idk how else to put this other than we can pass stuff around like crazy. Notice how the ```NewDemoHandler``` takes in a storage object then uses it. We can just use the stuff we "provide" in other functions. **This is the whole point of fx.**

### How to run this

```go run .```