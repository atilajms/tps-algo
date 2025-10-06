package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	c := colaEnlazada[T]{nil, nil}
	return &c
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil && c.ultimo == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := nodo[T]{elem, nil}

	if c.EstaVacia() {
		c.primero = &nuevoNodo
		c.ultimo = &nuevoNodo
	} else {
		c.ultimo.siguiente = &nuevoNodo
		c.ultimo = &nuevoNodo
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	primero := c.primero.dato
	c.primero = c.primero.siguiente
	if c.primero == nil {
		c.ultimo = nil
	}
	return primero
}
