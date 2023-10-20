/*
Demos printing of data structures and the effect of trying to convert json structures to string.
*/
package main

import (
	"fmt"
	"bytes"
	"encoding/json"
)

func main() {
	str := "Hello World!"
	fmt.Println(str)

	str_map := map[string]string{"key1":"magic", "key2":"mouse"}
	fmt.Println(str_map)

	int_map := map[string]int{"key1":12, "key2":13}
	fmt.Println(int_map)

	interface_map := map[string]interface{}{
		"id": 12345,
		"text": "some text &<> text",
	}
	fmt.Println(interface_map)

	fmt.Println("\nConverting json to object, certain characters are converted to unicode codepoints automatically: See Marshal() function in https://pkg.go.dev/encoding/json");
	json_bytes,_ := json.Marshal(interface_map)
	fmt.Println(string(json_bytes))
	// Output: {"id":12345,"text":"some text \u0026\u003c\u003e text"}

	fmt.Println("\nPrevent the conversion from taking place:")
	buf := new(bytes.Buffer)
	my_encoder := json.NewEncoder(buf)
	my_encoder.SetEscapeHTML(false)

	if my_encoder.Encode(&interface_map) == nil {/* Note that Go adds an extra newline at the end of the string with this. */
		encoder_str := buf.String()
		if len(encoder_str) > 1 { 
			final_str := string(encoder_str[0:len(encoder_str)-1]) /* Extra steps to remove trailing newline, if necessary. */
			fmt.Println(final_str)
		}
	}
	// Output: {"id":12345,"text":"some text &<> text"}
}