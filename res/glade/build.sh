#!/bin/env sh

SCRIPT=$(readlink -f "$0")
INSTALLPATH=$(dirname "${SCRIPT}")

echo "create glade.go"
echo "package glade" > $INSTALLPATH/glade.go

for glade in $INSTALLPATH/*.glade; do
    window=${glade##$INSTALLPATH/}
    window=${window%.glade}
    echo append $window    
    echo "const" $window 'string =`' $(cat ${window}.glade) '`' >> $INSTALLPATH/glade.go
done
