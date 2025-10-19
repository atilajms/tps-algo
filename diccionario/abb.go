package diccionario

import TDAPila "tdas/pila"

type funcCmp[K any] func(K, K) int

type abb[K any, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type nodoAbb[K any, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type iteradorArbol[K any, V any] struct {
	pila  TDAPila.Pila[nodoAbb[K, V]]
	cmp   funcCmp[K]
	desde *K
	hasta *K
}

func crearNodoAbb[K any, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		clave:     clave,
		dato:      dato,
		izquierdo: nil,
		derecho:   nil,
	}
}

func crearIteradorAbb[K, V any](a *abb[K, V], desde *K, hasta *K, cmp funcCmp[K]) *iteradorArbol[K, V] {
	pila := TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	apilarHastaElMinimo(a.raiz, pila, desde, hasta, cmp)
	return &iteradorArbol[K, V]{
		pila:  pila,
		cmp:   cmp,
		desde: desde,
		hasta: hasta,
	}
}

func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
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

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) Borrar(clave K) V {
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

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.raiz.iterar(visitar, a.cmp, desde, hasta) // llamo a la funcion interna con el nodo en vez del arbol
}

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	a.raiz.iterar(visitar, a.cmp, nil, nil)
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorAbb(a, nil, nil, a.cmp)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return crearIteradorAbb(a, desde, hasta, a.cmp)
}

func (it *iteradorArbol[K, V]) HaySiguiente() bool {
	return !it.pila.EstaVacia()
}

func (it *iteradorArbol[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (it *iteradorArbol[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.Desapilar()
	apilarHastaElMinimo(nodo.derecho, it.pila, it.desde, it.hasta, it.cmp)
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(K, V) bool, cmp funcCmp[K], desde *K, hasta *K) {
	if nodo == nil {
		return
	}

	if desde == nil || cmp(nodo.clave, *desde) >= 0 {
		nodo.izquierdo.iterar(visitar, cmp, desde, hasta)
	}

	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return
		}
	}
	if hasta == nil || cmp(nodo.clave, *hasta) <= 0 {
		nodo.derecho.iterar(visitar, cmp, desde, hasta)
	}

}

func apilarHastaElMinimo[K, V any](nodo *nodoAbb[K, V], pila TDAPila.Pila[nodoAbb[K, V]], desde, hasta *K, cmp funcCmp[K]) {
	for nodo != nil {
		if desde != nil && cmp(nodo.clave, *desde) < 0 { // empezar por un elemento mayor a desde
			nodo = nodo.derecho
		} else if hasta != nil && cmp(nodo.clave, *hasta) > 0 {
			nodo = nodo.izquierdo
		} else {
			pila.Apilar(*nodo)
			nodo = nodo.izquierdo
		}
	}
}
