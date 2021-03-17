module github.com/Hajimara/Go_Tutorial/average

go 1.16

require ( 
    github.com/Hajimara/Go_Tutorial/datafile v0.0.0
)

replace (
    github.com/Hajimara/Go_Tutorial/datafile v0.0.0 => ../datafile
)
