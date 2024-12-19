# Chapter 4 Struct

Struct is just a set of key value pair same idea like object bla bla

struct key can hold manu types

nested structs in Go

type car struct struct {
    Make string
    Model String
    Height int
    Width int
    FrontWheel Wheel
    BackWheel Wheel
}

type Wheel struct {
    Radius int
    Material string
}

the fields can be accesed using dot . operator

myCar := car{}
myCar.FrontWheel.Radius = 5


Anonymous struct 

struct instances that dont have a name 
dont use it unless have reason

only it used once

advantages used anonymous struct
-reuse one 

strcut is how we structured the json data typically


next is embedded struct 
embbeded vs nested

embedded kinda take other type 

type car struct {
    make string
    model string
}

type truck struct {
    car //inherinting everyhting in car 
    bedSize
}

notes:  but go is not and object oriented lagguage treat it like a 


the syntax to create a embbeded struct is really similar to nested

lanesTruck := truck {
    besSize:10
    car: car{
        make:"toyota"
        model:"camry"
    },
}

fmt printls(lanestrcut.BedSize)


public field is some sort of stuff we show o the leader board such as bio 

but password we dont hsow this is private field
embbeded private field within public user so easily nulify them when give aces public data but want to hide the mesage

lets talk about method in go 
struct method

go is not OOP but struct is belief to be similar defining method in struct is kinda practices.

why do we use method in struct 
mainly its jus syntehthic thing syntehtic sugar

we want to define on given typ 
called R.area kinda nic eway to do computed 