# gosb

Go Struct Builder

## options

|key|value|default|description|
|---|---|---|---|
|optional|(none)|no|この要素を指定しなくても良い|
|validation|(none)|no|この要素の検証を行う関数を設定できるようにする|
|nillable|(none)|no|この要素はNilでも可能|
|default|(none)|no|何も設定されていない場合に設定する値を代入する関数を指定する。Requiredを指定している場合はこの値を持って設定される|