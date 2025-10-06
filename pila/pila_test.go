package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() {
		pila.Desapilar()
	})
	require.Panics(t, func() {
		pila.VerTope()
	})
}

func TestUlltimoApilado(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[string]()
	p.Apilar("hola")
	require.Equal(t, "hola", p.VerTope())

	q := TDAPila.CrearPilaDinamica[rune]()
	q.Apilar('a')
	q.Apilar('b')
	q.Apilar('c')
	q.Apilar('d')
	q.Apilar('e')

	require.Equal(t, 'e', q.VerTope())

}

func TestLifo(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[bool]()
	p.Apilar(true)
	p.Apilar(false)
	require.Equal(t, false, p.Desapilar())

	q := TDAPila.CrearPilaDinamica[int]()
	q.Apilar(10)
	q.Apilar(11)
	q.Apilar(12)
	q.Apilar(13)
	q.Desapilar()

	require.Equal(t, 12, q.VerTope())
}

func TestApilarDesapilarUnElemento(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[int]()
	p.Apilar(42)
	require.False(t, p.EstaVacia())
	require.Equal(t, 42, p.Desapilar())
	require.True(t, p.EstaVacia())
	require.Panics(t, func() { p.VerTope() })
}

func TestVolumen(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[int]()

	for i := range 10000 {
		p.Apilar(i)
		require.Equal(t, i, p.VerTope())
	}

	for i := 9999; i >= 0; i-- {
		require.False(t, p.EstaVacia())
		tope := p.VerTope()
		require.Equal(t, i, tope)
		valor := p.Desapilar()
		require.Equal(t, i, valor)
	}

	require.True(t, p.EstaVacia())
	require.Panics(t, func() {
		p.VerTope()
	})
	require.Panics(t, func() {
		p.Desapilar()
	})
}
