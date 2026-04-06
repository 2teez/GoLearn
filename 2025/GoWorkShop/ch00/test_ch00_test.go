package main

import "testing"

func TestMyStringTypeTitleCase1(t *testing.T) {
	p := makePatientDefault()
	p.LastName = "javascript"
	if p.LastName != "Javascript" {
		t.Errorf("Excpected: %s, got: %s", "Javascript", p.LastName)
	}
}

func TestMyStringTypeTitleCase2(t *testing.T) {
	p := makePatient("java", "javascript", 23, false)
	if p.LastName != "Javascript" {
		t.Errorf("Excpected: %s, got: %s", "Javascript", p.LastName)
	}
}

func TestMyStringTypeTitleCase3(t *testing.T) {
	p := makePatient("CLOJURE", "JAVA", 23, false)
	if p.FirstName != "Clojure" {
		t.Errorf("Excpected: %s, got: %s", "Clojure", p.FirstName)
	}
}
