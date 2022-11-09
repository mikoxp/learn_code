package com.moles

import com.moles.xls.XlsReader

fun main() {
    val filepath = "./Test.xlsx"
    val reader=XlsReader()
    reader.readFromExcelFile(filepath)
}

