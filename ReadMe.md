# gosb

Go Struct Builder

## options

|key|value|default|description|
|---|---|---|---|
|optional|(none)|no|この要素を指定しなくても良い|
|validation|(none)|no|この要素の検証を行う関数を設定できるようにする|
|default|(none)|no|何も設定されていない場合に設定する値を代入する関数を指定する。Requiredを指定している場合はこの値を持って設定される|


## patterns
```
// Pointer: a *int, Optional, Default
var reala *int
if a == nil {
    reala = f.aDefault()
} else {
    reala = a
}
```
```
// not Pointer: a int, Optional, Default
// a is not pointer but optional
var reala int
if a == nil {
    reala = f.aDefault()
} else {
    reala = *a
}
```
```
// Pointer: a *int, not Optional, Default
var reala *int
if a == nil {
    reala = f.aDefault()
} else {
    reala = a
}
```
```
// not Pointer: a int, not Optional, Default
var reala int
if a == nil {
    reala = f.aDefault()
} else {
    reala = a
}
```
```
// (Pointer or not Pointer): a int, (Optional or not Optional), no Default
var reala [int, *int]
reala = a
```