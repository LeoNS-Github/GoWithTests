package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Quint")
	esperado := "Ola, Quint"

	if resultado != esperado {
		t.Errorf("resultado %q, esperado %q", resultado, esperado)
	}
}
