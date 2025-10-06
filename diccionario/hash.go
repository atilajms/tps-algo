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
	} else {
		iterador := buscarEnLaLista(clave, h.tabla[posHash])

		// si el elemento está, se actualiza el valor
		if iterador.VerActual() != nil {
			iterador.VerActual().dato = dato
		} else { // si no está
			iterador.Insertar(&parClaveValor[K, V]{clave, dato})
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
