# Router & HandlerFunc Explained

## The Problem We're Solving

You are building an HTTP server.

At some point the server must do this:

> **"When a request comes for `/users`, run THIS function."**

That "THIS function" is written by you, not the server.

So the server needs:
- A way to **store** that function
- A way to **call** that function later

**That's the real problem.**

---

## What is HandlerFunc in Plain English?

```go
type HandlerFunc func(req *HttpRequest) *HttpResponse
```

This just says:

> "A handler is a function that takes a request and returns a response."

Nothing more. Nothing less.

---

## Why Can't We Just Write Normal Functions?

You **do** write normal functions:

```go
func hello(req *HttpRequest) *HttpResponse {
    return NewTextResponse("hello")
}
```

But the router doesn't know this function by name.

The router only knows:
- Request method (`GET`)
- Request path (`/hello`)

So how does it connect `/hello` → `hello()`?

👉 **By storing the function in a variable.**

And for that, Go needs a **function type**.

---

## Why is HandlerFunc Defined at the Top?

Because the router literally **stores handlers as values**.

Look at this carefully:

```go
routes map[string]map[string]HandlerFunc
```

Read it slowly:

> For a method and path, store a function.

That function must have a **known shape**.

`HandlerFunc` defines that shape.

Without it, the router has no idea:
- What arguments the function takes
- What it returns

---

## What Happens When a Request Comes In?

Let's walk step-by-step.

### 1️⃣ You Register a Route

```go
router.GET("/hello", hello)
```

Here:
- `hello` is a function
- Go checks: does `hello` match `HandlerFunc`?
- If yes → store it

### 2️⃣ Router Stores It

```go
routes["GET"]["/hello"] = hello
```

So now the router literally holds a **function pointer**.

### 3️⃣ A Request Arrives

```
GET /hello
```

Server does:

```go
handler := router.FindHandler("GET", "/hello")
```

What is `handler`?

👉 **A function.**

### 4️⃣ Server Calls the Function

```go
response := handler(req)
```

That's it.

The server didn't:
- Parse JSON
- Write business logic
- Know what `/hello` does

It just **called a function**.

---

## Why is This IMPORTANT?

Because without `HandlerFunc`:
- Router can't store handlers
- Server can't call user logic
- Everything must be hardcoded
- Server becomes useless

---

## Think of It Like a Remote Control 📺

| Concept | Analogy |
|---------|---------|
| Button | URL (`/hello`) |
| Remote logic | Router |
| What happens when you press | `HandlerFunc` |

The remote doesn't care what the TV does.  
It just sends the signal.

---

## One Sentence Summary

> **HandlerFunc exists so the server can store and later execute user-written functions in a uniform way.**

---

## Why Every Framework Does This

| Language | Concept |
|----------|---------|
| Go | `HandlerFunc` |
| Java | Controller / Servlet |
| Node | `(req, res) => {}` |
| C++ | Function pointer / functor |

**Same idea everywhere.**

---

## Final Check ✅

Answer this in your head:

> "When a request comes, how does my server know which function to call?"

👉 **HandlerFunc is the answer.**
