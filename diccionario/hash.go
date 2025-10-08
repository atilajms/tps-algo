package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const tamInicial int = 20

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[*parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iteradorHash[K comparable, V any] struct {
	hash      *hashAbierto[K, V]
	posActual int
	iterLista TDALista.IteradorLista[*parClaveValor[K, V]]
}

func CrearHash[K comparable, T any]() Diccionario[K, T] {
	tabla := make([]TDALista.Lista[*parClaveValor[K, T]], tamInicial)
	// acá no sé si creamos listas vacias para cada posicion o si dejamos nil.
	// cambiarlo tiene implicaciones en los metodos y funciones auxiliares
	// for i := range tamInicial {
	// 	tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K,T]]()
	// }
	hash := hashAbierto[K, T]{tabla, tamInicial, 0}
	return &hash
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	posHash := hash(clave) % uint32(h.tam)

	// Si creamos para cada posicion una lista, este if no es necesario
	if h.tabla[posHash] == nil {
		h.tabla[posHash] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
		h.tabla[posHash].InsertarUltimo(&parClaveValor[K, V]{clave, dato})
		h.cantidad++
	} else {
		iterador := buscarEnLaLista(clave, h.tabla[posHash])

		// si el elemento está, se actualiza el valor
		if iterador.HaySiguiente() {
			iterador.VerActual().dato = dato
		} else { // si no está
			iterador.Insertar(&parClaveValor[K, V]{clave, dato})
			h.cantidad++
		}
	}
}

func buscarEnLaLista[K comparable, V any](
	clave K,
	lista TDALista.Lista[*parClaveValor[K, V]]) TDALista.IteradorLista[*parClaveValor[K, V]] {

	iter := lista.Iterador()

	for iter.HaySiguiente() {
		if iter.VerActual().clave == clave {
			return iter
		}
		iter.Siguiente()
	}
	return iter
}

func hash[K comparable](clave K) uint32 {
	h := fnv.New32a()
	h.Write(convertirABytes(clave))
	return h.Sum32()
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	//Si en el hash inician todas las listas creadas esto no va
	pos := hash(clave) % uint32(h.tam)
	lista := h.tabla[pos]
	if lista == nil {
		return false
	}

	iter := buscarEnLaLista(clave, lista)
	return iter.HaySiguiente() // solo hay siguiente si se encotro la clave
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	pos := hash(clave) % uint32(h.tam)
	lista := h.tabla[pos]

	iter := buscarEnLaLista(clave, lista)
	if !iter.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}
	return iter.VerActual().dato
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	pos := hash(clave) % uint32(h.tam)
	lista := h.tabla[pos]

	if lista == nil {
		panic("La clave no pertenece al diccionario")
	}

	iter := buscarEnLaLista(clave, lista)
	if !iter.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}

	dato := iter.VerActual().dato
	iter.Borrar()
	h.cantidad--

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
			//Ejecutamos la funcion
			if !visitar(par.clave, par.dato) {
				return // si devuelve false, se corta la iteración
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
		if lista != nil && !lista.EstaVacia() {
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
		if lista != nil && !lista.EstaVacia() {
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

	// Si la lista actual terminó, avanzar al próximo bucket válido
	for !it.iterLista.HaySiguiente() && it.posActual < it.hash.tam {
		it.posActual++
		if it.posActual >= it.hash.tam {
			return // terminó el hash
		}
		lista := it.hash.tabla[it.posActual]
		if lista != nil && !lista.EstaVacia() {
			it.iterLista = lista.Iterador()
			break
		}
	}
}
