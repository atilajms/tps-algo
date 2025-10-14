package diccionario

import TDAPila "tdas/pila"

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		clave:     clave,
		dato:      dato,
		izquierdo: nil,
		derecho:   nil,
	}
}

func CrearABB[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcionCmp}
}

func (a *abb[K, V]) buscar(clave K) **nodoAbb[K, V] {
	actual := &a.raiz

	for *actual != nil {
		cmp := a.cmp(clave, (*actual).clave)
		if cmp == 0 {
			// Si es la clave
			return actual
		} else if cmp < 0 {
			// Si es mas chica
			actual = &(*actual).izquierdo
		} else {
			// Si es mas grande
			actual = &(*actual).derecho
		}
	}
	//Si no esta devuelve el puntero a nil
	return actual
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	puntero := a.buscar(clave)
	// Si ya existe lo remplazamos
	if *puntero != nil { // No uso el pertenece para no tener q hacer la busqueda otra vez
		(*puntero).dato = dato
		return
	}
	//Si no lo creamos+
	*puntero = crearNodoAbb(clave, dato)
	a.cantidad++
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	puntero := a.buscar(clave)
	return *puntero != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	puntero := a.buscar(clave)
	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*puntero).dato
}

func (a abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a abb[K, V]) Borrar(clave K) V {
	puntero := a.buscar(clave)

	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := (*puntero).dato

	// Caso sin hijos
	if (*puntero).derecho == nil && (*puntero).izquierdo == nil {
		*puntero = nil

		// Caso un hijo
	} else if (*puntero).derecho == nil {
		*puntero = (*puntero).izquierdo
	} else if (*puntero).izquierdo == nil {
		*puntero = (*puntero).derecho

		//Dos hijos
	} else {
		// BUsco un sucesor
		sucesor := &(*puntero).izquierdo
		for (*sucesor).derecho != nil {
			sucesor = &(*sucesor).derecho
		}
		(*puntero).clave = (*sucesor).clave
		(*puntero).dato = (*sucesor).dato

		// Borro alsucesor
		*sucesor = (*sucesor).izquierdo
	}
	a.cantidad--
	return dato
}

func (a abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	if a.raiz == nil {
		return
	}

	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	actual := a.raiz

	for actual != nil || !pila.EstaVacia() {
		// Bajo a la izquierda
		for actual != nil {
			pila.Apilar(actual)
			actual = actual.izquierdo
		}

		// sao el nodo de la pila
		n := pila.Desapilar()
		actual = n

		// aplico la func
		if !f(actual.clave, actual.dato) {
			return // si devuelve false
		}

		actual = actual.derecho
	}
}
