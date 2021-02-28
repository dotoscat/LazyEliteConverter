*LazyEliteConverter*

This is a dead simple tool to convert BMP images from a folder to another format

The program reads the contents of a source folder, without doing recursion, and generates the converted files to an output folder

There is option to remove the origin files, defaults is preserve the source
There is option to select the format, defaults is png

** Example of use

EliteLazyConverter -s ./one -o ./one --remove-originals -f png
