package com.moles.xls

import org.apache.poi.ss.usermodel.WorkbookFactory
import java.io.FileInputStream

class XlsReader {
    fun readFromExcelFile(filepath: String) {
        val inputStream = FileInputStream(filepath)
        var xlWb = WorkbookFactory.create(inputStream)
        val xlWs = xlWb.getSheetAt(0)
        var rowNumber = 0
        while(true){
            var columnNumber = 0
            val row = xlWs.getRow(rowNumber)
            row?:break
            while(true){
                val cell = row.getCell(columnNumber) ?: break
                print("$cell ")
                columnNumber++
            }
            rowNumber++
            println()
        }
    }

}