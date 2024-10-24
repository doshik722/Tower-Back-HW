package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	c_ = flag.Bool("c", false, "количество встречаний строки во входных данных")
	d_ = flag.Bool("d", false, "строки, которые повторились во входных данных")
	u_ = flag.Bool("u", false, "строки, которые не повторились во входных данных.")
	f_ = flag.Int("f", 0, "не учитывать первые num_fields полей в строке")
	s_ = flag.Int("s", 0, "не учитывать первые num_chars символов в строке")
	i_ = flag.Bool("i", false, "не учитывать регистр букв")
)

type flags struct {
	count_vhod   bool
	dub_string   bool
	uniq_string  bool
	f_num_fields int
	s_num_char   int
	i_register   bool
}

type OptionFunc func(*flags)

func c() OptionFunc {
	return func(f *flags) {
		f.count_vhod = true
	}
}

func d() OptionFunc {
	return func(f *flags) {
		f.dub_string = true
	}
}

func u() OptionFunc {
	return func(f *flags) {
		f.uniq_string = true
	}
}

func f(n int) OptionFunc {
	return func(f *flags) {
		f.f_num_fields = n
	}
}

func s(n int) OptionFunc {
	return func(f *flags) {
		f.s_num_char = n
	}
}

func i() OptionFunc {
	return func(f *flags) {
		f.i_register = true
	}
}

func Map(str []string, flag_param flags) ([]string, error) {
	counts := make(map[string]int)
	var final []string

	for _, line := range str {
		if flag_param.i_register {
			line = strings.ToLower(line)
		}

		if flag_param.f_num_fields > 0 {
			fields := strings.Fields(line)
			if len(fields) > flag_param.f_num_fields {
				line = strings.Join(fields[flag_param.f_num_fields:], " ")
			} else {
				line = ""
			}
		}

		if flag_param.s_num_char > 0 && len(line) > flag_param.s_num_char {
			line = line[flag_param.s_num_char:]
		}

		counts[line]++
	}

	if (flag_param.uniq_string && flag_param.dub_string) || (flag_param.dub_string && flag_param.count_vhod) || (flag_param.uniq_string && flag_param.count_vhod) {
		return nil, fmt.Errorf("нельзя использовать одновременно -c, -d и -u")
	}

	if flag_param.uniq_string {
		for key, value := range counts {
			if value == 1 {
				final = append(final, key)
			}
		}
	} else if flag_param.dub_string {
		for key, value := range counts {
			if value > 1 {
				final = append(final, key)
			}
		}
	} else if flag_param.count_vhod {
		for key, value := range counts {
			final = append(final, strconv.Itoa(value)+" "+key)
		}
	} else { // Если ни один флаг не установлен
		for key := range counts {
			final = append(final, key)
		}
	}

	return final, nil
}
func ReadFile(inputFile *string) []string {

	var lines []string

	var file *os.File
	if *inputFile == "" {
		file = os.Stdin // Используем stdin, если файл не указан
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	} else {
		var err error
		file, err = os.Open(*inputFile)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка открытия входного файла: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}
func WriteFile(outputFile *string) *os.File {
	var out *os.File
	if *outputFile == "" {
		out = os.Stdout // Используем stdout, если файл не указан
	} else {
		var err error
		out, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка создания выходного файла: %v\n", err)
			os.Exit(1)
		}

	}
	return out
}
func main() {
	inputFile := flag.String("input_file", "", "Путь к входному файлу (по умолчанию stdin)")
	outputFile := flag.String("output_file", "", "Путь к выходному файлу (по умолчанию stdout)")
	flag.Parse()

	var lines = ReadFile(inputFile)
	var out = WriteFile(outputFile)
	defer out.Close()
	flagParam := flags{
		count_vhod:   *c_,
		dub_string:   *d_,
		uniq_string:  *u_,
		f_num_fields: *f_,
		s_num_char:   *s_,
		i_register:   *i_,
	}

	linesOut, err := Map(lines, flagParam)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	for _, line := range linesOut {
		_, err := out.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			os.Exit(1)
		}
	}
}
