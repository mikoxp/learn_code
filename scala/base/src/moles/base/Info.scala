package moles.base

class Info( var v:String,var i:Int) {
    var value:String=v;
    var id:Int=i;
    
    def print(): Unit ={
      printf("value=%s, id=%d\n",value,id)
    }
}
