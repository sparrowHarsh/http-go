package src

type HandlerFunc func(req *HttpRequest) *HttpResponse

/* ROuter stores handler paht for a method and path */
type Router struct {
	// method -> path -> handler
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) Handle(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.Handle("GET", path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.Handle("POST", path, handler)
}

func (r *Router) FindHandler(method, path string) HandlerFunc {
	if methods, ok := r.routes[method]; ok {
		if handler, ok := methods[path]; ok {
			return handler
		}
	}
	return nil
}
