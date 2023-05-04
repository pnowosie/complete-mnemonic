#!/usr/bin/env bash

FILE=single${LEN}.txt
if ! [[ -f "samples/${FILE}" ]]; then
    echo "set LEN envvar to one of 12, 18, 24 value"
    exit 1
fi

LINE=$(grep "^${WORD}" samples/${FILE})
if [[ $? != 0 ]]; then
    echo "Word ${WORD} not found in file ${FILE}"
    exit 1
fi

CHK=$(echo ${LINE} | cut -f2 -d' ')
echo "$(yes ${WORD} | head -$(( ${LEN}-1 )) | xargs) ${CHK}"