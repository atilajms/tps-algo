package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() {
		cola.Desencolar()
	})
	require.Panics(t, func() {
		cola.VerPrimero()
	})
}

func TestEncolarDesencolarUnElemento(t *testing.T) {
	c := TDACola.CrearColaEnlazada[string]()
	c.Encolar("a")
	require.False(t, c.EstaVacia())
	require.Equal(t, "a", c.Desencolar())
	require.True(t, c.EstaVacia())
	require.Panics(t, func() { c.VerPrimero() })
}

func TestFifo(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[rune]()

	cola.Encolar('a')
	require.False(t, cola.EstaVacia())
	require.Equal(t, 'a', cola.VerPrimero())

	cola.Encolar('b')
	require.Equal(t, 'a', cola.VerPrimero())

	cola.Encolar('c')
	require.Equal(t, 'a', cola.VerPrimero())

	val := cola.Desencolar()
	require.Equal(t, 'a', val)
	require.Equal(t, 'b', cola.VerPrimero())

	val = cola.Desencolar()
	require.Equal(t, 'b', val)
	require.Equal(t, 'c', cola.VerPrimero())

	val = cola.Desencolar()
	require.Equal(t, 'c', val)
	require.True(t, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for i := range 10000 {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}

	for i := range 10000 {
		require.False(t, cola.EstaVacia())
		primero := cola.VerPrimero()
		require.Equal(t, i, primero)
		val := cola.Desencolar()
		require.Equal(t, i, val)
	}

}
