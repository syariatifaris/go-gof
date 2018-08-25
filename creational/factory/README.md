### Factory

Create an instances of derived object.We create an object without exposing the creation logic. This pattern fully use an interface  

**Sample Case**
1. Strategic approach: Algorithm selection
2. Switching implementation (i.e: Get data from db or redis)

**When to use it**
1. Subclasses figure out what objects should be created.
2. Parent instance allows later instantiation to subclasses means the creation of object is done when it is required.
3. The process of objects creation is required to centralize within the application.
4. A  (creator) will not know what classes it will be required to create.