package main

import "testing"
import "reflect"

func TestMap_uniq_fields(t *testing.T) {
	input := []string{"Привет, Паша", "Пока, Паша", "Привет, Андрей", "умпалумпа"}
	flagParam := flags{
		uniq_string:  true,
		f_num_fields: 1,
	}

	expected := []string{"Андрей", " "}
	result, err := Map(input, flagParam)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMap_Registr_count(t *testing.T) {
	input := []string{"Персик", "ПеРсик", "Что?", "11", "11"}
	flagParam := flags{
		count_vhod: true,
		i_register: true,
	}

	expected := []string{"2 персик", "1 что?", "2 11"}
	result, err := Map(input, flagParam)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
