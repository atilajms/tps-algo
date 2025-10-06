package pila

const capacidadInicial = 3

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]

}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if len(p.datos) == p.cantidad {
		redimensionar(p, 2*p.cantidad)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++

}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	p.cantidad--
	ultimo := p.datos[p.cantidad]
	if p.cantidad <= len(p.datos)/4 && p.cantidad > capacidadInicial {
		redimensionar(p, len(p.datos)/2)
	}
	return ultimo
}

func redimensionar[T any](p *pilaDinamica[T], nueva_capacidad int) {
	sliceRedimensionado := make([]T, nueva_capacidad)
	copy(sliceRedimensionado, p.datos[:p.cantidad])
	p.datos = sliceRedimensionado
}

func CrearPilaDinamica[T any]() Pila[T] {
	datos := make([]T, capacidadInicial)
	cantidad := 0
	p := pilaDinamica[T]{datos, cantidad}
	return &p
}
