package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const tamInicial int = 19

const factorDeCargaMax float64 = 3

type parClaveValor[K any, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K any, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
	esIgual func(K, K) bool
}

type iteradorHash[K any, V any] struct {
	hash      *hashAbierto[K, V]
	posActual int
	iterLista TDALista.IteradorLista[parClaveValor[K, V]]
}

func CrearHash[K any, V any](comparar func(K, K) bool) Diccionario[K, V] {
    tabla := make([]TDALista.Lista[parClaveValor[K, V]], tamInicial)
    // listas vacias (??)
    // for i := range tabla {
    //     tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
    // }

    hash := hashAbierto[K, V]{tabla, tamInicial, 0, comparar}
    return &hash
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
    if float64(h.cantidad)/float64(h.tam) > factorDeCargaMax {
        redimensionar(h, h.tam*2)
    }

    iterLista, pos := obtenerIterador(h, clave)

    // listas vacias (??)
    if iterLista == nil {
        lista := TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
        lista.InsertarUltimo(parClaveValor[K, V]{clave, dato})
        h.tabla[pos] = lista
        h.cantidad++
    } else {
        if iterLista.HaySiguiente() {
            iterLista.Borrar()
        } else {   
            h.cantidad++
        }
	iterLista.Insertar(parClaveValor[K, V]{clave, dato})
    }
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
    // listas vacias (??) etc
    iterLista, _ := obtenerIterador(h, clave)
    if iterLista == nil {
        return false
    }
    return iterLista.HaySiguiente() // solo hay siguiente si se encotro la clave
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
    iterLista, _ := obtenerIterador(h, clave)
    if iterLista == nil || !iterLista.HaySiguiente() {
        panic("La clave no pertenece al diccionario")
    }
    return iterLista.VerActual().dato
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
    iterLista, _ := obtenerIterador(h, clave)
    if iterLista == nil || !iterLista.HaySiguiente() {
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

func obtenerIterador[K, V any](
    h *hashAbierto[K, V],
    clave K,
) (TDALista.IteradorLista[parClaveValor[K, V]], uint32) {

    posHash := obtenerHash(clave) % uint32(h.tam)
    lista := h.tabla[posHash]

    if lista == nil {
        return nil, posHash
    }

    return buscarEnLaLista(clave, lista, h.esIgual), posHash
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
    nuevaTabla := make([]TDALista.Lista[parClaveValor[K, V]], nuevoTam)

    for _, lista := range h.tabla {
        if lista != nil {
            iter := lista.Iterador()
            for iter.HaySiguiente() {
                par := iter.VerActual()
                posHash := obtenerHash(par.clave) % uint32(nuevoTam)

                if nuevaTabla[posHash] == nil {
                    nuevaTabla[posHash] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
                }

                nuevaTabla[posHash].InsertarUltimo(par)
                iter.Siguiente()
            }
        }
    }

    h.tabla = nuevaTabla
    h.tam = nuevoTam
}
