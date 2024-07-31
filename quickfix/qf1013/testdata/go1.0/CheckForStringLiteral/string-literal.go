package pkg

func fn2(_ string) {
}

func fn() {
	_ = "{\"foo\":true}"                 //@ diag(`Simplify string by using raw string literal`)
	_ = "escaped backslash: \\ foo"      //@ diag(`Simplify string by using raw string literal`)
	_ = "double escaped new line: \\n"   //@ diag(`Simplify string by using raw string literal`)
	_ = "triple escaped new line: \\\\n" //@ diag(`Simplify string by using raw string literal`)

	const _ = "escaped double quotes: \"" //@ diag(`Simplify string by using raw string literal`)
	fn2("escaped double quotes: \"")      //@ diag(`Simplify string by using raw string literal`)
	a := "escaped double quotes: \""      //@ diag(`Simplify string by using raw string literal`)

	fn2(a)
	const _ = "abc"
	_ = "abc"
	_ = "escaped backslash \\ plus a back quote `"
	_ = "escaped double quotes \" plus a back quote `"
	_ = "escaped double quotes \" plus a new line \n"
	_ = "escaped double quotes \" plus a tab \t"

	_ = `abc`
}
