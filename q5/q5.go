package q5

import (
	"fmt"
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	chars := strings.Split(s, "")
	for _, ss := range s {
		s = strings.ReplaceAll(s, "A", "")
		s = strings.ReplaceAll(s, "a", "")
		s = strings.ReplaceAll(s, "E", "")
		s = strings.ReplaceAll(s, "e", "")
		s = strings.ReplaceAll(s, "I", "")
		s = strings.ReplaceAll(s, "i", "")
		s = strings.ReplaceAll(s, "o", "")
		s = strings.ReplaceAll(s, "O", "")
		s = strings.ReplaceAll(s, "U", "")
		s = strings.ReplaceAll(s, "u", "")
		fmt.Println(ss)
	}
	for _, sss := range s {
		s = strings.ReplaceAll(s, "B", "b")
		s = strings.ReplaceAll(s, "C", "c")
		s = strings.ReplaceAll(s, "D", "d")
		s = strings.ReplaceAll(s, "F", "f")
		s = strings.ReplaceAll(s, "G", "g")
		s = strings.ReplaceAll(s, "H", "h")
		s = strings.ReplaceAll(s, "J", "j")
		s = strings.ReplaceAll(s, "K", "k")
		s = strings.ReplaceAll(s, "L,", "l")
		s = strings.ReplaceAll(s, "M", "m")
		s = strings.ReplaceAll(s, "N", "n")
		s = strings.ReplaceAll(s, "P", "p")
		s = strings.ReplaceAll(s, "Q", "q")
		s = strings.ReplaceAll(s, "R", "r")
		s = strings.ReplaceAll(s, "S", "s")
		s = strings.ReplaceAll(s, "T", "t")
		s = strings.ReplaceAll(s, "V", "v")
		s = strings.ReplaceAll(s, "W", "w")
		s = strings.ReplaceAll(s, "X", "x")
		s = strings.ReplaceAll(s, "Y", "y")
		s = strings.ReplaceAll(s, "Z", "z")

		fmt.Println(sss)
	}
	for i, char := range chars {
		chars[i] = "." + char

	}
	return s
}
