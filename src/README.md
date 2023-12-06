# CleverCalc
This is golang and Fyne Desktop implementation of the calculator with equal calculating and graph building.

A programm has tab interface with several additional window such as:
 - History - display operation history
 - Graph - display graph
 - Help - display help

Tabs of main window:
 - CALC - for equating without X
 - EQUAL - for equating with X
 - GRAPH - for displaying graph for X and Y ranges.

 Programm was created with MVP architecture. Presenter communicate with model and view throw they interfaces. 

 Assigning folder:
  - assets - additional file such as help and icon
  - build - binary files, scripts, template and done packages
  - cmd - main.go
  - config - clevercalc configuration file
  - internal - go source
  - test - additional file for test
  - tmp - for tempgraph.png
  - var - for history file

Print "make help" for printing information about command for building and testing