package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const tamInicial int = 19

const factorDeCargaMax float64 = 2.7

type parClaveValor[K any, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K any, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
	esIgual  func(K, K) bool
}

type iteradorHash[K any, V any] struct {
	hash      *hashAbierto[K, V]
	posActual int
	iterLista TDALista.IteradorLista[parClaveValor[K, V]]
}

func CrearHash[K any, V any](comparar func(K, K) bool) Diccionario[K, V] {
	tabla := crearListas[K, V](tamInicial)
	hash := hashAbierto[K, V]{tabla, tamInicial, 0, comparar}
	return &hash
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	if float64(h.cantidad)/float64(h.tam) > factorDeCargaMax {
		redimensionar(h, h.tam*2)
	}
	iterLista := obtenerIterador(h, clave)
	if iterLista.HaySiguiente() {
		iterLista.Borrar()
	} else {
		h.cantidad++
	}
	iterLista.Insertar(parClaveValor[K, V]{clave, dato})
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	iterLista := obtenerIterador(h, clave)
	return iterLista.HaySiguiente()
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	iterLista := obtenerIterador(h, clave)
	if !iterLista.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}
	return iterLista.VerActual().dato
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	iterLista := obtenerIterador(h, clave)
	if !iterLista.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}
	dato := iterLista.VerActual().dato
	iterLista.Borrar()
	h.cantidad--
	if h.tam > tamInicial && float64(h.cantidad)/float64(h.tam) < factorDeCargaMax/4 {
		redimensionar(h, h.tam/2)
	}

	return dato
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, lista := range h.tabla {
		if lista == nil {
			continue
		}
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			if !visitar(par.clave, par.dato) {
				return
			}
			iter.Siguiente()
		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorHash[K, V]{hash: h, posActual: 0}

	// Buscar la primera lista no vacía
	for iter.posActual < len(h.tabla) {
		lista := h.tabla[iter.posActual]
		if !lista.EstaVacia() {
			iter.iterLista = lista.Iterador()
			break
		}
		iter.posActual++
	}
	return iter
}

func (it *iteradorHash[K, V]) HaySiguiente() bool {
	if it.posActual >= it.hash.tam || it.iterLista == nil {
		return false
	}

	if it.iterLista.HaySiguiente() {
		return true
	}

	// Buscar siguiente lista
	for i := it.posActual + 1; i < it.hash.tam; i++ {
		lista := it.hash.tabla[i]
		if !lista.EstaVacia() {
			return true
		}
	}
	return false
}

func (it *iteradorHash[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	par := it.iterLista.VerActual()
	return par.clave, par.dato
}

func (it *iteradorHash[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	it.iterLista.Siguiente()

	// Si la lista actual terminó, avanzar hasta otra con elementos o hasta el fin
	for !it.iterLista.HaySiguiente() && it.posActual < it.hash.tam {
		it.posActual++
		if it.posActual >= it.hash.tam {
			return // terminó el hash
		}
		lista := it.hash.tabla[it.posActual]
		if !lista.EstaVacia() {
			it.iterLista = lista.Iterador()
			break
		}
	}
}

func obtenerIterador[K, V any](
	h *hashAbierto[K, V],
	clave K,
) TDALista.IteradorLista[parClaveValor[K, V]] {

	posHash := obtenerHash(clave) % uint32(h.tam)
	lista := h.tabla[posHash]

	return buscarEnLaLista(clave, lista, h.esIgual)
}

func buscarEnLaLista[K, V any](
	clave K,
	lista TDALista.Lista[parClaveValor[K, V]],
	esIgual func(K, K) bool,
) TDALista.IteradorLista[parClaveValor[K, V]] {

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if esIgual(iter.VerActual().clave, clave) {
			return iter
		}
		iter.Siguiente()
	}
	return iter
}

func obtenerHash[K any](clave K) uint32 {
	h := fnv.New32a()
	h.Write(convertirABytes(clave))
	return h.Sum32()
}

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func redimensionar[K, V any](h *hashAbierto[K, V], nuevoTam int) {
	nuevaTabla := crearListas[K, V](nuevoTam)
	for _, lista := range h.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			posHash := obtenerHash(par.clave) % uint32(nuevoTam)
			nuevaTabla[posHash].InsertarUltimo(par)
			iter.Siguiente()
		}
	}
	h.tabla = nuevaTabla
	h.tam = nuevoTam
}

func crearListas[K, V any](tam int) []TDALista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
}
