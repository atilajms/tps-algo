package diccionario_test

import (
	"testing"

	"tdas/diccionario"

	"github.com/stretchr/testify/require"
)

func crearABBEnteros() diccionario.DiccionarioOrdenado[int, int] {
	abb := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	abb.Guardar(10, 100)
	abb.Guardar(5, 50)
	abb.Guardar(15, 150)
	abb.Guardar(3, 30)
	abb.Guardar(7, 70)
	abb.Guardar(12, 120)
	abb.Guardar(20, 200)
	return abb
}

func TestIterarSinCorte(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	abb.Iterar(func(clave, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 720, suma)
}

func TestIterarConCorte(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	abb.Iterar(func(clave, dato int) bool {
		suma += dato
		return suma < 200
	})
	require.True(t, suma >= 200)
	require.True(t, suma < 670)
}

func TestIterarRangoSinLimites(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	abb.IterarRango(nil, nil, func(clave, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 720, suma)
}

func TestIterarRangoConLimites(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	desde := 6
	hasta := 15
	abb.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 440, suma)
}

func TestIterarAbbVacio(t *testing.T) {
	abb := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	suma := 0
	abb.Iterar(func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 0, suma)

	abb.IterarRango(nil, nil, func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 0, suma)
}


func TestIterarOrdenInOrder(t *testing.T) {
	abb := crearABBEnteros()
	claves := []int{}
	abb.Iterar(func(clave, _ int) bool {
		claves = append(claves, clave)
		return true
	})
	require.Equal(t, []int{3, 5, 7, 10, 12, 15, 20}, claves)
}

func TestIteradorExternoCompleto(t *testing.T) {
	abb := crearABBEnteros()
	iter := abb.IteradorRango(nil, nil)

	var claves []int
	var datos []int
	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		claves = append(claves, clave)
		datos = append(datos, dato)
		iter.Siguiente()
	}

	require.Equal(t, []int{3, 5, 7, 10, 12, 15, 20}, claves)
	require.Equal(t, []int{30, 50, 70, 100, 120, 150, 200}, datos)
}

func TestIteradorExternoRangoLimitado(t *testing.T) {
	abb := crearABBEnteros()
	desde := 5
	hasta := 12
	iter := abb.IteradorRango(&desde, &hasta)

	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	require.Equal(t, []int{5, 7, 10, 12}, claves)
}

func TestIteradorExternoVacio(t *testing.T) {
	abb := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	iter := abb.IteradorRango(nil, nil)
	require.False(t, iter.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	})
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Siguiente()
	})
}

func TestIteradorExternoUnElemento(t *testing.T) {
	abb := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	abb.Guardar(42, 99)

	iter := abb.IteradorRango(nil, nil)
	require.True(t, iter.HaySiguiente())
	clave, dato := iter.VerActual()
	require.Equal(t, 42, clave)
	require.Equal(t, 99, dato)

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorRangoSinResultados(t *testing.T) {
	abb := crearABBEnteros()
	desde := 0
	hasta := 2
	iter := abb.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente(), "não deveria haver elementos no rango [0,2]")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoInclusivo(t *testing.T) {
	abb := crearABBEnteros()
	desde := 5
	hasta := 10
	iter := abb.IteradorRango(&desde, &hasta)

	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	require.Equal(t, []int{5, 7, 10}, claves)
}


