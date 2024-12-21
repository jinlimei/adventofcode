package main

const dayGo = `package %%package%%

type Day struct {}

`

const partGo = `package %%package%%

import "fmt"

func (d Day) Part%%day%%Prompt() {
	fmt.Println("Part %%day%%")
	panic("not implemented")
}

func (d Day) Part%%day%%Actual() {
	fmt.Println("Part %%day%%")
	panic("not implemented")
}

`

const libGo = `package %%package%%

`
