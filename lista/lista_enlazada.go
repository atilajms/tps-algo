package lista

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iterador[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodo[T]
	anterior *nodo[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	l := listaEnlazada[T]{nil, nil, 0}
	return &l
}

func crearNodo[T any](elem T) *nodo[T] {
	return &nodo[T]{elem, nil}
}

func crearIterador[T any](l *listaEnlazada[T]) *iterador[T] {
	return &iterador[T]{
		lista:    l,
		actual:   l.primero,
		anterior: nil,
	}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNodo(elem)
	if l.EstaVacia() {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = l.primero
		l.primero = nuevoNodo
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevoNodo := crearNodo(elem)

	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.siguiente = nuevoNodo
	}
	l.ultimo = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := l.primero.dato
	l.primero = l.primero.siguiente
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return crearIterador(l)
}

func (it *iterador[T]) VerActual() T {
	if it.actual == nil {
		panic("El iterador termino de iterar")
	}
	return it.actual.dato
}

func (it *iterador[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *iterador[T]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	it.anterior = it.actual
	it.actual = it.actual.siguiente
}

func (it *iterador[T]) Insertar(dato T) {
	nuevo := crearNodo(dato)

	// Si no hay primero
	if it.anterior == nil {
		it.lista.primero = nuevo
	} else {
		it.anterior.siguiente = nuevo
	}
	if it.actual == nil {
		it.lista.ultimo = nuevo
	}
	it.lista.largo++
	nuevo.siguiente = it.actual
	it.actual = nuevo
}

func (it *iterador[T]) Borrar() T {
	if it.actual == nil {
		panic("El iterador termino de iterar")
	}

	dato := it.actual.dato

	// Primer nodo
	if it.anterior == nil {
		it.lista.primero = it.actual.siguiente
	} else {
		it.anterior.siguiente = it.actual.siguiente
	}

	// ultimo
	if it.actual.siguiente == nil {
		it.lista.ultimo = it.anterior
	}

	it.actual = it.actual.siguiente
	it.lista.largo--

	return dato
}
