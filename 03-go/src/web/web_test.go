package main

import "testing"

func TestWeb(t *testing.T){
  resultado := Greeting("Code.education Rocks!")
  if resultado != "<b>Code.education Rocks!</b>"{
    t.Errorf("O Texto esperado: %s, obtido: %s", "<b>Code.education Rocks!</b>", resultado)
  }
}